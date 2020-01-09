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

[gRPC command line tool](https://github.com/grpc/grpc/blob/master/doc/command_line_tool.md) を利用するのが簡単です。

Macなら以下の手順でインストール可能です。

```
brew install gflags

brew tap grpc/grpc

brew install grpc
```

`which grpc_cli` を実行して `/usr/local/bin/grpc_cli` 等が表示されればインストール出来ています。

`grpc_cli ls localhost:9998` を実行するとgRPCサーバーの情報が取得出来ます。

次のようにメソッドを指定するとそのインターフェースを確認出来ます。

`grpc_cli ls localhost:9998 Cat.FindCuteCat -l`

以下のようにメソッド名とパラメータを渡してあげればgRPCのメソッドを実行出来ます。

```
grpc_cli call localhost:9998 Cat.FindCuteCat 'catId: "moko"'
```
