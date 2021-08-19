/*
 * http.Postを使ったサンプルです
 * タイムアウト値やクッキー値などを指定したい場合は、こちらを使用するとよいでしょう。
 */
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
)

func main() {

    // Postデータ
    values := url.Values{}
    values.Add("q1", "hoge")
    values.Add("q2", "fuga")

    // POSTリクエスト発行
    rsp, err := http.Post("http://httpbin.org/post", "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
    if err != nil {
        fmt.Println(err)
        return
    }
    // 関数を抜ける際に必ずresponseをcloseするようにdeferでcloseを呼ぶ
    defer rsp.Body.Close()

    // レスポンスを取得し出力
    body, _ := ioutil.ReadAll(rsp.Body)
    fmt.Println(string(body))
}
