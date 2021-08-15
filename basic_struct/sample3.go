/*
 * Structの定義と同時に初期化する方法について
 */
package main

import "fmt"

func main() {

    // 宣言と同時に初期化
    normalStruct := struct {
        name string
        age  int
    }{
        name: "anonymous",
        age:  30,
    }
    fmt.Println(normalStruct)

    // 関数や計算を使用して初期化する
    useFuncStruct := struct {
        name string
        age  int
    }{
        name: fmt.Sprintf("hoge"),
        age:  1+29,
    }
    fmt.Println(useFuncStruct)

    // 宣言と同時に初期化(JSON)
    jsonStruct := struct {
        Message string `json:"message"`
        Number int `json:"number"`
    }{
        "hello",
        1,
    }
    fmt.Println(jsonStruct)

}
