/*
 * 値渡しとポインタ渡しをする関数サンプル
 */
package main

import "fmt"

type Product struct {
    name string
    price string
}

func main() {
    p := Product{name:"hoge", price:"priceless"}

    // 値渡し
    printProduct(p)

    // ポインタ渡し
    printProductPointer(&p)
}

// 値渡し
func printProduct(p Product) {
	fmt.Println(p)
}

// ポインタ渡し
func printProductPointer(p *Product) {
	fmt.Println(*p)
}
