/*
 * 様々なスライスの初期化
 */
package main

import (
    "fmt"
)

func main() {

	var s1 []int
	fmt.Println("s1: ", s1)

	var s2 []int = []int{10, 20, 30}
	fmt.Println("s2: ", s2)

	var s3 = []int{10, 20, 30}
	fmt.Println("s3: ", s3)

	s4 := []int{}
	fmt.Println("s4: ", s4)

	s5 := []int{10, 20, 30}
	fmt.Println("s5: ", s5)

	// 多重スライス
	var s6 [][]int
	fmt.Println("s6: ", s6)

	s7 := [][]int{
		{10, 20, 30},
		{10, 20, 30},
	}
	fmt.Println("s7: ", s7)

	s8:= make([]int, 0, 10)
	fmt.Printf("s8: %#v, len: %d, cap: %d\n", s8, len(s8), cap(s8))

	// 長さに 1 以上の数値を指定すると、その分がゼロ値(int型の場合は0)で初期化される
	s9 := make([]int, 3, 10)
	fmt.Printf("s9: %#v, len: %d, cap: %d\n", s9, len(s9), cap(s9))

}
