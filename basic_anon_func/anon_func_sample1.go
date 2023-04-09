// see: https://code-graffiti.com/anonymous-function-in-golang/
package main

import "fmt"

func foo() {
	fmt.Println("foo関数の処理です。")
}

func main() {
	foo()

	func() {
		fmt.Println("無名関数の処理です。")
	}()

	func(x int) {
		fmt.Println("無名関数の処理に用いた引数の値:", x)
	}(2020)
}
