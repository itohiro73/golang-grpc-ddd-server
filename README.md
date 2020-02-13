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
-H 'authorization: Bearer CatSecret9999' \
-d '{"catId":"moko"}' \
localhost:9998 Cat.FindCuteCat
```

認証パラメータに渡す `CatSecret9999` は文字列であれば何でも良いですが、何か文字列を渡さないと認証エラーになります。

GUI製だと [BloomRPC](https://github.com/uw-labs/bloomrpc) が便利です。

こちらも `brew cask install bloomrpc` で手軽にインストールが可能です。

左上のプラスボタンから `pb/cat.proto` をロードする事が可能です。

<img width="1318" alt="BloomRPC" src="https://user-images.githubusercontent.com/11032365/74407093-9028fe80-4e74-11ea-9112-2371364140d7.png">

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

## gRPCドキュメントの自動生成

プロジェクトルートで以下を実行するとHTML形式のドキュメントを生成します。

```
cd /go/app/
protoc --doc_out=html,index.html:./docs pb/*.proto
```

`docs/index.html` が自動生成されたファイルになります。
