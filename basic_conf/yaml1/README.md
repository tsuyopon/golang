# 概要
yamlから設定ファイルを読み込むサンプルです
"gopkg.in/yaml.v3"を使っています。


# 使い方
```
$ go run main.go 
main.Config{DB:main.DBConfig{Host:"127.0.0.1", Port:3306, Database:"sample", Username:"user", Password:""}, APIKey:""}
```

環境変数を指定することでyamlに埋め込むことができます。
```
$ DB_PASSWORD=mypass API_KEY=testkey go run main.go
main.Config{DB:main.DBConfig{Host:"127.0.0.1", Port:3306, Database:"sample", Username:"user", Password:"mypass"}, APIKey:"testkey"}
```
