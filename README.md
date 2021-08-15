# golang
golang


# モジュール

- https://github.com/golang/go/wiki/Modules


### GOモジュールの使い方の流れ
- 1. $ go env -w GO111MODULE=onを設定する
- 2. $ go mod init
- 3. $ go build or $ go get でモジュールをインストールする
- 4. $ go mod tidy で使われてないモジュールを削除する
- 5. その他

### モジュール関連コマンド
```
# go.modファイルの初期化
$ go mod init <module-name>

# モジュールのインストール
$ go get <package>

# 不要なモジュールのお掃除用
$ go mod tidy

# 使っているモジュールのリストアップ
$ go list -m all

# moduleのバージョンのリストアップ
$ go list -m -versions <module>
```


# 参考資料
- とほほのGo言語入門
  - 最速言語マスター的な資料としてざっと見るのにおすすめ
  - http://www.tohoho-web.com/ex/golang.html
- Go by Example
  - 非常に多くのサンプルがあって全て理解しておくべき資料
  - https://gobyexample.com/
  - https://github.com/mmcgrana/gobyexample
