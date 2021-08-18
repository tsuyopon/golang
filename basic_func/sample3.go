/*
 * 可変長引数を扱う関数サンプル
 */

package main

import "fmt"

// 可変長関数のラッピング関数
func sampleWrap(s string, strs ...string) {
    sample(s, strs...)
}

// 可変長は...で定義される
func sample(value string, s ...string) {
    fmt.Printf("#####"+value+"start#####\n")
    fmt.Println(s)
    for _, str := range s {
        fmt.Printf(str+"\n")
    }
    fmt.Printf("\n")
}

func main() {
    sampleWrap("test1")// test1 []
    sampleWrap("test2", "a")// test2 [a] a
    sampleWrap("test3", "a", "b", "c")// test3 [a, b, c] abc
}
