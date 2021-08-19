/*
 * http.NewRequestのサンプルです
 */
package main

import (
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"
)

func main() {

    client := &http.Client{}
    data := url.Values{"foo": {"bar"}}

    req, _ := http.NewRequest(
        "POST",
        "http://httpbin.org/post",
        strings.NewReader(data.Encode()),
     )

    // リクエストにヘッダ情報を追加
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    // POSTリクエスト発行
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    println(string(body))
}
