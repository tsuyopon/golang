# 概要
go言語でテストを作成する際のサンプル

# 詳細

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


### 簡単なcoverageを取得してみる
テストが作成されていれば、coverageレポートも取得できます。
```
$ go test -v -coverprofile=cover.out 
=== RUN   TestSqure
--- PASS: TestSqure (0.00s)
=== RUN   Test_Division_1
    division_test.go:11: はじめのテストがパスしました
--- PASS: Test_Division_1 (0.00s)
PASS
coverage: 80.0% of statements
ok  	gin_test/golang/basic_test/sample1	0.338s
```

ではHTMLで視覚的にカバレッジレポートを確認します。
```
$ go tool cover -html=cover.out -o cover.html
$ open cover.html 
```


# 参考資料
- https://www.yoheim.net/blog.php?q=20170903
