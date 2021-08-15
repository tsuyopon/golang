/*
 * 関数でスライスを扱う場合
 */
package main

import (
    "fmt"
)

func change1(a []int) {
    a[2] += 100
}

func change2(b []int) {
    b[2] += 200
}

func main() {
    a := []int{1, 2, 3, 4, 5}
    change1(a)
    fmt.Println(a) // [1 2 103 4 5]
    change2(a)
    fmt.Println(a) // [1 2 303 4 5]
}
