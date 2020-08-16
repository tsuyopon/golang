/*
 * makeを利用した単純なサンプルプログラムです
 */
package main

import (
    "fmt"
)

func main() {
    var a []int
    a = make([]int, 5, 10)
    fmt.Println(a)      // [0 0 0 0 0]
    fmt.Println(len(a)) // 5    lenは現在利用している長さ
    fmt.Println(cap(a)) // 10   capは最大格納容量として宣言している長さを表す
    for i, _ := range a {
        a[i] = i
    }
    fmt.Println(a) // [0 1 2 3 4]

    b := make([]int, 10)
    fmt.Println(b)      // [0 0 0 0 0 0 0 0 0 0]
    fmt.Println(len(b)) // 10
    fmt.Println(cap(b)) // 10
}
