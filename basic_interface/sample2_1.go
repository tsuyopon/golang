/*
 * "type XXX interface" で定義した要件を満たしているかどうかの判別にも i.(T) が使えます。
 *
 * 構造体が定義したメソッドを持っているかどうかをチェックするためにインターフェース宣言ができます。
 * このプログラムでは「type MyInterface interface」としてHelloメソッドを必ず実装するインターフェースとしてMyInterfaceを定義しています。
 *
 * Helloはfunc定義にレシーバとしてBを持っていますので、Bは持っていますが、Aは持っていません。
 * このプログラムではAとBがHelloメソッドを持っているかどうかを判定します。
 */

package main

type A struct {
}

type B struct {
}

func (b B) Hello(name string) {
    println("hello", name)
}

// ここでinterfaceを定義している。
type MyInterface interface {
    Hello(name string)
}

func main() {
    var ok bool

    // A の構造体は Hello メソッドを持っていないので false
    a := interface{}(A{})
    _, ok = a.(MyInterface)
    println("A:", ok) // A: false

    // B の構造体は Hello メソッドを持っているので true
    b := interface{}(B{})
    _, ok = b.(MyInterface)
    println("B:", ok) // B: true
}
