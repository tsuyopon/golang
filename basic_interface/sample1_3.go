/*
 * interface{}型を使うことでswitch文の中でも型判定できるようになります。
 * i.(type) はswitch文の中でのみ利用することができます。
 */
package main

import "fmt"

func do(i interface{}) {

    // switch文の中で変数の型を判定できるようにします。
    // i.(type)  で判定でき、これはswitch文の中でしか利用できないようです。
    switch variable := i.(type) {
    case int:
        fmt.Printf("int: %d\n", variable)
    case string:
        fmt.Printf("string: %s\n", variable)
    default:
        fmt.Println("Default")
    }
}

func main() {
    do(23) //=> 23
    do("hello") //=> hello
    do(true) //=> Default
}
