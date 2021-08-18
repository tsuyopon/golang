/*
 * 単純な関数のサンプルと複数の戻り値を返す関数の例
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

// 複数の戻り値を返す関数
func swap(x, y int) (int, int) {
	return y, x
}

func main() {
    p("Hello World")
    p(100)
    p(true)
    p(add(1000, 9000))
    a, b := swap(1000, 9000)
    fmt.Println(a, b)
}
