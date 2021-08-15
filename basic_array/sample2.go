/*
 * 配列の初期化方法part2
 */
package main

import "fmt"

func main(){

	var a1 [3]int
	fmt.Println(a1, len(a1), cap(a1)) // [0 0 0] 3 3

	var a2 = [2]int{}
	fmt.Println(a2, len(a2), cap(a2)) // [0 0] 2 2

	var a3 = [4]int{10}
	fmt.Println(a3, len(a3), cap(a3)) // [10 0 0 0] 4 4

	a4 := [...]int{10, 20, 30}
	fmt.Println(a4, len(a4), cap(a4)) // [10 20 30] 3 3

	a5 := [3]int{10, 20, 30}
	fmt.Println(a5, len(a5), cap(a5)) // [10 20 30] 3 3

}
