# 概要
単純なソケット通信をするサーバ・クライアントプログラムです。

# 使い方
以下でport番号を8888と指定してサーバを起動して
```
$ go run server.go 8888
```

クライアントで接続します。以下の例ではlocalhost:8888に接続します。終了時はSTOPと入力します。
```
$ go run client.go localhost:8888
>> hoge
->: 2020-08-11T05:59:52+09:00
>> fuga
->: 2020-08-11T05:59:53+09:00
>> piyo
->: 2020-08-11T05:59:55+09:00
>> STOP
->: TCP client exiting...
```

# 参考
- https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
