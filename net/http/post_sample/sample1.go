/*
 * http.PostForm(url, data) により、POSTリクエストを投げてみます
 * 非常に単純なPOST処理のサンプルです
 */
package main

import (
    "net/http"
    "net/url"
    "io/ioutil"
)

func main() {

    resp, _ := http.PostForm(
        "http://httpbin.org/post",
        url.Values{"foo": {"bar"}, },
    )

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    println(string(body))
}
