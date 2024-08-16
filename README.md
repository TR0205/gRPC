# 概要
gRPC学習

# 利用方法
.protoからgoのpbファイルを生成。goのバイナリファイルを/bin以下に生成する。
```bash
$ make greet
```

生成したバイナリファイルを実行。まずはサーバー
```bash
$ ./bin/greet/server
2024/08/16 12:31:41 Listening on: 0.0.0.0:50051
```
クライアント
```bash
$ ./bin/greet/client
2024/08/16 12:31:46 doGreat was invoked
2024/08/16 12:31:46 Greeting: Hello Takashi
```