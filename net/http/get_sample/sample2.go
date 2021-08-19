/*
 * Cookieやヘッダを指定する場合にはhttp.NewRequest を使う
 * リクエストヘッダー、レスポンスヘッダーの内容を取得する
 */
package main

import (
    "net/http"
    "net/http/cookiejar"
    "io/ioutil"

    // for dump use
    "fmt"
    "net/http/httputil"
)

func main() {

    cookieJar, _ := cookiejar.New(nil)

    client := &http.Client {
        Jar: cookieJar,
    }

    req, _:= http.NewRequest("GET", "http://httpbin.org/basic-auth/myuser/mypasswd", nil)
    req.Header.Set("User-Agent", "My User Agent")
    req.Header.Set("X-Test", "ほげほげ")
    req.SetBasicAuth("myuser", "mypasswd")

    // 以下の3行はRequestダンプ出力
    dump, _ := httputil.DumpRequestOut(req, true)
    fmt.Println("### httputil.DumpRequestOut")
    fmt.Printf("%s", dump)

    res, _ := client.Do(req)

    // 以下の3行はResponseダンプ出力
    dumpResp, _ := httputil.DumpResponse(res, true)
    fmt.Println("### httputil.DumpResponse")
    fmt.Printf("%s", dumpResp)

    body, _ := ioutil.ReadAll(res.Body)
    defer res.Body.Close()

    fmt.Println("### output")
    println(string(body))
}
