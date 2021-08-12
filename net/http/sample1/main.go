package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

func main() {
    os.Exit(run(os.Args))
}

func run(args []string) int {

    // httpのクライアントを作成する
    client := &http.Client{}

    // タイムアウトの設定
    client.Timeout = time.Second * 15

    // リクエストを作成
    req, err := http.NewRequest("GET", "http://www.yahoo.co.jp",nil)
    if err != nil {
        return 1
    }

    // リクエストを実行
    resp, err := client.Do(req)
    if err != nil {
        return 2
    }
    defer resp.Body.Close()

    // レスポンスの読み込み
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return 3
    }

    // レスポンスの表示
    fmt.Printf("%s", body)

    return 0
}
