/*
 * 構造体がネストするサンプルプログラム
 */
package main;

import "fmt"

// 構造体はネストさせることができます。下記の例ではFooがBarの構造体を含んでいます。
type Foo struct {
	Bar
	name    string
	fooProp string
}

type Bar struct {
	barProp string
}

func main() {
	foo := Foo{Bar{"barProp"}, "foo","fooProp"}
	fmt.Println(foo.name)        // foo
	fmt.Println(foo.barProp)     //barProp
	fmt.Println(foo.Bar.barProp) //barProp
}
