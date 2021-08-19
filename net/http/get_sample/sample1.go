/*
 * ’Pƒ‚ÈGET‚ÌƒTƒ“ƒvƒ‹‚Å‚·
 */
package main

import (
    "net/http"
    "io/ioutil"
)

func main() {

    response, _ := http.Get("http://httpbin.org/get")
    body, _ := ioutil.ReadAll(response.Body)
    defer response.Body.Close()

    println(string(body))

}
