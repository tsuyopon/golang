# 概要
UDPのサーバ・クライアントサンプルを確認する

# 使い方
サーバを起動する
```
$ go run server.go 8888
```

クラインアントを起動する。STOPで終了することができます。
```
$ go run client.go localhost:8888
The UDP server is 127.0.0.1:8888
>> hoge
Reply: 306
>> fuga
Reply: 164
>> piyo
Reply: 148
>> STOP
Exiting UDP client!
```

# 参考
- https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/
