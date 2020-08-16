/*
 * sample1_3.goではintとstringの型アサーション判定を行なっていた。
 * ここでは、boolやintはint32, int16, int64と細分化して判定できることを示すサンプルです。
 *
 */
package main

func main() {
    var v int64 = 100

    // interface型を定義する
    i := interface{}(v)

    // interface型をassertするには i.(type) を使います。 go言語のswitch分にはbreakは不要です。
    switch i.(type) {
    case string:
        s := i.(string)
        println("i is string:", s)
    case bool:
        b := i.(bool)
        println("i is boolean:", b)
    case int32:
        println("int32")
    case int16:
        println("int16")
    case int64:
        println("int64") // int64
    }
}
