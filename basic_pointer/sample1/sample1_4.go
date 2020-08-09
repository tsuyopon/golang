/*
 * アドレスとアドレスから変数の値を参照する方法
 */
package main

import "fmt"

func main() {
    name := "太郎"
    fmt.Printf("name :%v\n", name)

    namePoint := &name

    // namePointは、&nameが格納されているだけなので、stringへのポインタである *string型の値が格納されている。
    fmt.Printf("namePoint :%v\n", namePoint)

    // namePointが指している変数は、"*namePoint"という感じで、"*"をつけて表す。
    fmt.Printf("namePoint :%v\n", *namePoint)
}

