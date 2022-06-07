# 初期構築

## .protoをビルドする

※本来ならば下記のリンクの内容で対応ができるが、installしたモジュールなどがうまく使えない場合は、手動でインストールしてパスを通す必要がある
https://grpc.io/docs/languages/go/quickstart/

1. [protoc](https://github.com/protocolbuffers/protobuf)をインストールする。Protocol buffersのコンパイラとなる
2. [Go用のProtocolBuffers](https://github.com/protocolbuffers/protobuf-go)をインストールする
3. [Go用のgRPCフレームワーク](https://github.com/grpc/grpc-go)をインストールする
4. 1-3の各ツールをビルド等して、パスを通す

## .protoファイルを作成する


## .protoをビルドする

1. protoc --go_out=. --go-grpc_out=. {protoファイルパス}
2. ボイラープレートコードが自動生成されるので、確認する

## serverコードを実装する

1. 
