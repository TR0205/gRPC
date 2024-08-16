# 概要
gRPC学習

# 利用方法
.protoからgoのpbファイルを生成。goのバイナリファイルを/bin以下に生成する。
```bash
$ make greet
```

生成したバイナリファイルを実行
```bash
$ ./bin/greet/server
```
