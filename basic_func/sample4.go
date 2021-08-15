/*
 * 構造体と関数を紐づける(レシーバ)
 */
package main

import "fmt"

type Human struct {
    Name   string
    Age    int
    Weight float32
}

// Goにはクラス構文が存在しません。代わりに、typeで定義した型に処理(method)を紐づけることができます。
// funcと関数名の間に「(変数名 構造体)」を書き加えることで構造体と関数が紐付けられます。
// (h Human)の部分をレシーバと呼びます
func (h Human) addWeight(f float32) float32 {
    return h.Weight + f
}

func main() {
    // インスタンス生成
    personal := Human{
        Name:   "taro",
        Age:    22,
        Weight: 80.6,
    }

    fmt.Printf("%#v \n", personal)
    personal.Weight = personal.addWeight(10)    // addWeight関数はHuman構造体の変数であるpersonalから呼び出される
    fmt.Printf("%#v \n", personal)
}
