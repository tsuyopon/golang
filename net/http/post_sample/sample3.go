/*
 * client.Postのサンプルです
 * デバッグ出力もしています。
 */
package main

import (
    "net/http"
    "net/url"
    "strings"
    "io/ioutil"

    // for dump use
    "fmt"
    "net/http/httputil"
)

func main() {

    client := &http.Client{}
    data := url.Values{"foo": {"bar"}}

    resp, _ := client.Post(
        "http://httpbin.org/post",
        "application/x-www-form-urlencoded",
        strings.NewReader(data.Encode()),
    )

    // 以下の3行はResponseダンプ出力
    dumpResp, _ := httputil.DumpResponse(resp, true)
    fmt.Println("### httputil.DumpResponse")
    fmt.Printf("%s", dumpResp)

    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    fmt.Println("### output")
    println(string(body))
}
