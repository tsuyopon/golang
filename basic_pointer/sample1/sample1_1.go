/*
 * 単純なポインタ型変数利用サンプル
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
    // pがポインタ変数
    // *Personポインタ型
    var p *Person

    p = &Person{
        Name: "太郎",
        Age:  20,
    }
    fmt.Printf("変数pに格納されているアドレス :%p", p)
}
