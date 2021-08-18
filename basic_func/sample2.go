/*
 * Named return Value
 *  Goでの戻り値となる変数に名前をつける( named return value )ことができます。戻り値に名前をつけると、関数の最初で定義した変数名として扱われます。
 */
package main

import "fmt"

// 以下の関数はreturnしていないが、定義でx, yをreturnするので２つの値を返す関数です。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
