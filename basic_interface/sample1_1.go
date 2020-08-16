/*
 * interace{}型を用いることで、どのような型でも入れられるサンプルです。
 * Go言語には、全ての型と互換性を持っている interface{}型 というものが存在しています。これは「空のインターフェース」と呼ばれます。
 * 
 */
package main

import "fmt"

func main() {

    // 空のインターフェース
    var i interface{}

    i = 10
    fmt.Println(i) // 10

    i = true
    fmt.Println(i) // true

    i = "hello"
    fmt.Println(i) // hello
    fmt.Println(i == interface{}("hello")) // true


    // ポイント: interface{}型を定義することで、どんな型でも入れておける入れ物(C言語のvoid型のようなもの)をていぎすることができます。
    // インターフェース型のスライス
    a := []interface{}{
        10, true, "hello",
    }
    fmt.Printf("%#v\n", a) // []interface {}{10, true, "hello"}


    // キーがstring型で値がインターフェース型のマップ
    m := make(map[string]interface{})
    m["one"] = 10
    m["two"] = true
    m["three"] = "hello"

    fmt.Printf("%#v\n", m) // map[string]interface {}{"one":10, "three":"hello", "two":true}
}
