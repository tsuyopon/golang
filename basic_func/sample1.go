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
// 関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。 ( see: https://go-tour-jp.appspot.com/basics/5 )
// 以下の例では1つ目も2つ目もint型なので「func add(i, j int) int { 」と記述可能です。
func add(i int, j int) int {
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
