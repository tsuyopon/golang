# 概要
go言語でテスト実行、カバレッジレポートを作成する際のサンプル

以下のドキュメントも併せて参考にすると良い
- The Go Blog: The cover story
  - https://blog.golang.org/cover
- How to Write Go Code
  - https://golang.org/doc/code#Testing
- 参考ソースコード
  - https://github.com/codecov/example-go
- Go testing: Testingパッケージの説明
  - https://pkg.go.dev/testing

# 詳細

### テストを実行する
テスト実行で"go test"だけだと味気ない出力しかでないので
```
$ go test
PASS
ok  	gin_test/golang/basic_test/sample1	0.221s
```

vオプションを付与することでどのテストが成功して、どのテストが失敗したのかがわかります。
```
$ go test -v
=== RUN   TestSqure
--- PASS: TestSqure (0.00s)
=== RUN   Test_Division_1
    division_test.go:11: はじめのテストがパスしました
--- PASS: Test_Division_1 (0.00s)
PASS
ok  	gin_test/golang/basic_test/sample1	0.109s
```


### coverageを表示する
```
$ go test -v -cover
=== RUN   TestSqure
--- PASS: TestSqure (0.00s)
=== RUN   Test_Division_1
    division_test.go:11: はじめのテストがパスしました
--- PASS: Test_Division_1 (0.00s)
=== RUN   TestSize
--- PASS: TestSize (0.00s)
PASS
coverage: 58.3% of statements
ok  	gin_test/golang/basic_test/sample1	0.279s
```

### coverageレポートを取得してみる
テストが作成されていれば、coverageレポートも取得できます。
```
$ go test -v -coverprofile=cover.out
=== RUN   TestSqure
--- PASS: TestSqure (0.00s)
=== RUN   Test_Division_1
    division_test.go:11: はじめのテストがパスしました
--- PASS: Test_Division_1 (0.00s)
=== RUN   TestSize
--- PASS: TestSize (0.00s)
PASS
coverage: 58.3% of statements
ok  	gin_test/golang/basic_test/sample1	0.103s
```

ではHTMLで視覚的にカバレッジレポートを確認します。
```
$ go tool cover -html=cover.out -o cover.html
$ open cover.html 
```

### 特定のコードをカバレッジ対象から除外する
モック関連のソースコード(mock.go)を除外する場合には、次のようにします。
```
$ go test -v -coverprofile=cover.out.tmp
$ cat cover.out.tmp | grep -v "**_mock.go" > cover.out
$ go tool cover -html=cover.out -o cover.html
$ open cover.html 
```


# 参考資料
- https://www.yoheim.net/blog.php?q=20170903
