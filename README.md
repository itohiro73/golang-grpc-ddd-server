# golang-grpc-server
GoでgRPCを実装する時のサンプルコード

## 起動

Dockerで動作させます。

Docker for Mac等をインストールして下さい。

### 起動スクリプト1（ホットリロードが有効）

```
./scripts/docker-compose-up.sh
```

### 起動スクリプト2（ホットリロード、リモートデバッグが有効）

```
./scripts/docker-compose-up-debug.sh
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

レイヤードアーキテクチャーでCRUDを実装しているDogのgRPCメソッドは下記のように実行できます。

```
grpcurl -plaintext \
-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE' \
-d '{"id":1, "name":"Chiro", "kind":"Shiba-ken"}' \
localhost:9998 Dog.AddCuteDog
```

```
grpcurl -plaintext \
-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE' \
-d '{"id":1}' \
localhost:9998 Dog.FindCuteDog

```

```
grpcurl -plaintext \
-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE' \
-d '{"id":1, "name":"チロ", "kind":"柴犬"}' \
localhost:9998 Dog.UpdateCuteDog
```

```
grpcurl -plaintext \
-H 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxIn0.rTCH8cLoGxAm_xw68z-zXVKi9ie6xJn9tnVWjd_9ftE' \
-d '{"id":1}' \
localhost:9998 Dog.DeleteCuteDog
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

例えば `api/proto/dog.proto` を以下の内容で作成します。

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
protoc --go_out=plugins=grpc:. api/proto/dog.proto
```

そうすると `api/proto/dog.pb.go` が出力されます。

出力された `api/proto/dog.pb.go` のインターフェースを満たすようにgRPCサーバーを実装します。

なお出力された `pb.go` は他プロジェクトでも利用する可能性があるので `pkg/pb/` に移動して下さい。

gRPCサービスやメソッドを増やす際は必ず上記の手順を行います。

### healthCheck用のgRPCサービス

`api/proto/health.proto` のような `.proto` 内で `import` を利用している場合は以下のように `.proto` ファイルのBuild時に `import` されたライブラリのパスを指定する必要があります。

```
cd /go/app/
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  api/proto/health.proto
```

## grpc-gatewayの生成

[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) を使ってREST APIでgRPCメソッドを呼び出せるサーバーを立ち上げます。

以下を実行すると `google.golang.org/grpc/health/grpc_health_v1/health.pb.go` が更新されます。

`api/proto/health.proto` を変更したら必ずこの手順を行って下さい。

```
cd /go/app/
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  api/proto/health.proto
```

普通にHTTPサーバーなので以下のコマンドで動作確認が可能です。

```
curl -v http://localhost:8081/grpc/health

# 結果
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8081 (#0)
> GET /grpc/health HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/7.64.1
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Grpc-Metadata-Content-Type: application/grpc
< Date: Tue, 18 Feb 2020 07:12:21 GMT
< Content-Length: 20
<
* Connection #0 to host localhost left intact
{"status":"SERVING"}* Closing connection 0
```

## gRPCドキュメントの自動生成

プロジェクトルートで以下を実行するとHTML形式のドキュメントを生成します。

```
cd /go/app/
protoc --doc_out=html,index.html:./docs api/proto/cat.proto
```

`docs/index.html` が自動生成されたファイルになります。

## ディレクトリ構成

下記のようなディレクトリ構成になっています。

[Standard Go Project Layout](https://github.com/golang-standards/project-layout) を参考にしてあります。

```
golang-grpc-server/
  ├ api/
  │  └ proto/           # gRPCのインターフェース定義となるProtocol Bufferのファイル置き場
  ├ build/
  │  └ package/
  │    └ docker/        # Dockerfileの置き場
  ├ docs/
  ├ google.golang.org/  # Googleが公開しているpackage
  ├ internal/           # 本アプリケーションでのみ利用するpackage
  ├ pkg/                # 他プロジェクトに公開するpackage、ここではgRPCのGoのインターフェース
  ├ scripts             # デプロイ等に利用するシェルスクリプト
```
