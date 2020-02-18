# golang-grpc-server
GoでgRPCを実装する時のサンプルコード

## 起動

Dockerで動作させます。

Docker for Mac等をインストールして下さい。

### 起動スクリプト1（ホットリロードが有効）

```
docker-compose-up.sh
```

### 起動スクリプト2（ホットリロード、リモートデバッグが有効）

```
docker-compose-up-debug.sh
```

デバッグの方法等は筆者が以前書いた下記の記事を参考にして下さい。

https://qiita.com/keitakn/items/f46347f871083356149b

## ソースコードのフォーマット

プロジェクトルートで以下を実行して下さい。

`gofmt -l -s -w .`

## 動作確認

[grpcurl](https://github.com/fullstorydev/grpcurl) を利用するのが簡単です。

Macなら以下の手順でインストール可能です。

```
brew install grpcurl
```

`which grpcurl` を実行して `/usr/local/bin/grpcurl` 等が表示されればインストール出来ています。

`grpcurl -plaintext localhost:9998 list` を実行するとgRPCサーバーの情報が取得出来ます。

以下のようにメソッド名とパラメータを渡してあげればgRPCのメソッドを実行出来ます。

```
grpcurl -plaintext \
-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE' \
-d '{"catId":"moko"}' \
localhost:9998 Cat.FindCuteCat
```

認証パラメータに渡す `authorization: Bearer` は有効なJWTである必要があります。

https://jwt.io/ でJWTを作成します。

payloadの中身ですが `{ sub: "1" }` のようにユーザーIDをセットします。

参考までに以下のJWTを記載しておきます。

```
# { "sub": "1" }
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE

# { "sub": "2" }
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIn0.a7ktMGTybA32ykWHRvhp8FTEsBb-g3FN8aBB6FbgBo0

# { "sub": "3" }
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzIn0.FYQ95U6iKXxxRIAydKNuDmxsQybrIQh-lXs6Cqs_97M
```

正常に認可されるのは `{ "sub": "1" }` のみで、それ以外はgRPCサーバーがエラーを返します。

GUI製だと [BloomRPC](https://github.com/uw-labs/bloomrpc) が便利です。

こちらも `brew cask install bloomrpc` で手軽にインストールが可能です。

左上のプラスボタンから `pb/cat.proto` をロードする事が可能です。

![BloomRPC](https://user-images.githubusercontent.com/11032365/74523153-ea52be00-4f5f-11ea-94b7-944c6241dd7d.png)

METADATAの部分にはauthorizationトークンをJSONで指定します。

```
# ユーザーID 1 これしか正常に認可が通らない
{
  "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE"
}

# ユーザーID 2
{
  "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIyIn0.a7ktMGTybA32ykWHRvhp8FTEsBb-g3FN8aBB6FbgBo0"
}

# ユーザーID 3
{
  "authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIzIn0.FYQ95U6iKXxxRIAydKNuDmxsQybrIQh-lXs6Cqs_97M"
}
```

### ヘルスチェックメソッドの動作確認

`authorization: Bearer` の設定は必要ありません。

以下で呼び出しが可能です。

```bazaar
grpcurl -plaintext \
localhost:9998 grpc.health.v1.Health/Check
```

## `.proto` からGoのインターフェースを作成する

例えば `pb/dog.proto` を以下の内容で作成します。

```
syntax = "proto3";

service Dog {
    rpc FindCuteDog (FindCuteDogMessage) returns (CuteDogResponse) {}
}

message FindCuteDogMessage {
    string DogId = 1;
}

message CuteDogResponse {
    string name = 1;
    string kind = 2;
}
```

コンテナ内のアプリケーションプロジェクトルートで以下のコマンドを実行します。

```
cd /go/app/
protoc --go_out=plugins=grpc:. pb/dog.proto
```

そうすると `pb/dog.pb.go` が出力されます。

出力された `pb/dog.pb.go` のインターフェースを満たすようにgRPCサーバーを実装します。

gRPCサービスやメソッドを増やす際は必ず上記の手順を行います。

### healthCheck用のgRPCサービス

`pb/health.proto` のような `.proto` 内で `import` を利用している場合は以下のように `.proto` ファイルのBuild時に `import` されたライブラリのパスを指定する必要があります。

```
cd /go/app/
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  pb/health.proto
```

## gRPCドキュメントの自動生成

プロジェクトルートで以下を実行するとHTML形式のドキュメントを生成します。

```
cd /go/app/
protoc --doc_out=html,index.html:./docs pb/*.proto
```

`docs/index.html` が自動生成されたファイルになります。
