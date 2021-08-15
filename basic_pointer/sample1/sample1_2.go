/*
 * 変数の値とアドレスを出力させるだけの単純なプログラム
 */
package main

import "fmt"

// Person は人間を表す構造体。
type Person struct {
    Name string
    Age  int
}

func main() {
    // ポインタ型の変数を宣言する
    //   pがポインタ型変数
    //   Personポインタ型
    var p *Person

    // ポインタ型変数pにPersonのポインタを格納するので&を付与する
    p = &Person{
        Name: "太郎",
        Age:  20,
    }
    fmt.Printf("p :%+v\n", p)
    fmt.Printf("変数pに格納されているアドレス :%p\n", p)
}

