# 概要
テストを並行で実行するためのサンプル

# 詳細

### 実行結果

```
$ go test -v
=== RUN   Test_Func1
Test_Func1 entered
Test_Func1 returned
--- PASS: Test_Func1 (0.00s)
=== RUN   Test_Func2
Test_Func2 entered
=== PAUSE Test_Func2
=== RUN   Test_Func3
Test_Func3 entered
Test_Func3 returned
--- PASS: Test_Func3 (0.00s)
=== RUN   Test_Func4
Test_Func4 entered
=== PAUSE Test_Func4
=== RUN   Test_Func5
Test_Func5 entered
Test_Func5 returned
--- PASS: Test_Func5 (0.00s)
=== CONT  Test_Func2
Test_Func2 returned
--- PASS: Test_Func2 (0.00s)
=== CONT  Test_Func4
Test_Func4 returned
--- PASS: Test_Func4 (0.00s)
PASS
ok  	gin_test/golang/basic_test	0.308s
```

- 一時停止した場合、=== PAUSEと表示され、処理が再開した場合、=== CONTと表示されます。


# 参考資料
- https://engineering.mercari.com/blog/entry/how_to_use_t_parallel/
