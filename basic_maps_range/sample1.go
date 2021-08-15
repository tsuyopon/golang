/*
 * mapの初期化方法について
 */
package main

import "fmt"

func main(){

    // 初期化例1
    ex1 := make(map[string]string, 2) //マップの宣言
    fmt.Println(ex1) //=> map[] //宣言された空のマップ

    ex1["firstName"] = "Mike"
    ex1["lastName"] = "Smith"
    fmt.Println(ex1)

    // 初期化例2
    var ex2 = map[string]string{"firstName":"John", "lastName": "Smith"}
    fmt.Println(ex2)

}
