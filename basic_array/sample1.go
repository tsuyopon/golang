/*
 * 配列の初期化方法part1
 */
package main

import "fmt"

func main(){

    //宣言方法(1) 宣言後に要素毎に初期化する
    var arr1 [2]string 
    arr1[0] = "Golang"
    arr1[1] = "Ruby"

    //宣言方法(2) 宣言と同時に要素数とその要素を指定する方法
    var arr2 [2]string = [2]string{"Golang", "Ruby"}

    //宣言方法(3) 要素数を指定せずにその要素を指定する方法
    var arr3 = [...]string{"Golang", "Ruby"}

    fmt.Println(arr1, arr2, arr3) //=> [Golang Ruby] [Golang Ruby] [Golang Ruby]


}
