/*
 * 単純な関数のサンプル
 */
package main

import "fmt"

// 1つの引数
func p(i interface{}) {
    fmt.Println(i)
}

// 複数引数を指定する
func add(i int, j int) int {  // func add(i, j int) int {   と記述することも可能
    return i + j
}

func main() {
    p("Hello World")
    p(100)
    p(true)
    p(add(1000, 9000))
}
