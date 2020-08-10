package main

import (
    "crypto/tls"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"

    "golang.org/x/net/http2"
)

func main() {
    hc := &http.Client{
        Transport: &http2.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: true,
            },
        },
    }

    resp, err := hc.Get("https://www.yahoo.co.jp/")
    if err != nil {
        log.Fatal(err)
    }

    body, _ := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    fmt.Printf("%s", body)

}
