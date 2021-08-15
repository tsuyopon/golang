/*
 * クラスのメンバー変数を扱うような操作を行います
 * (goにはクラスを扱う概念が存在しませんのでクラスもどきです)
 */
package main

import "fmt"

// 変数を構造体でまとめる
type myStruct struct {
    x int
    y int
}

// 以下のレシーバー(後述)を持つ定義のことをメソッドと呼びます。
// 以下のメソッド宣言で (p *myStruct) の部分をレシーバーと呼びます。ここで登録した構造体の変数から呼び出すことができます。
//
// レシーバは次の２つがある。
//    1) 値レシーバ        (例) (p myStruct)
//    2) ポインタレシーバ  (例) (p *myStruct)
//
// 構造体を操作する関数 構造体の「.」演算子で呼び出す

// ポインタレシーバ
func (p *myStruct) test() int {
	var ret int
	ret = p.x + p.y

	// ポインタレシーバなので、構造体中の値は変更可能
	p.x = 100
	p.y = 200

    return ret
}

// 値レシーバ
func (p myStruct) test2() int {
	var ret int
	ret = p.x + p.y

	// ポインタレシーバなので、構造体中の値は変更できない。mainから確認すると値が変更されていないことを確認できる。
	p.x = 1000
	p.y = 2000

    return ret
}

func main() {
    var a myStruct
    a.x = 1
    a.y = 2
    fmt.Println(a)
    a.test()             // ポインタレシーバのメソッドを呼び出す
    fmt.Println(a)
    a.test2()            // 値レシーバのメソッドを呼び出す
    fmt.Println(a)
}
