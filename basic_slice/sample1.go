/*
 * 単純なsliceサンプルです
 */
package main;

import (
	"fmt"
)

func main(){

    /*
     * [ポイント] 配列とスライスは定義方法が非常に似ているので違いを押さえておきます。
     *
     * 配列の場合
     *    array0 := [3]int {1, 2, 3}
     *    array1 := [...]int {1, 2, 3, 4}
     *
     * スライスの場合
     *    slice0 := []int {1, 2, 3}
     *    slice1 := make([]int, 3, 5)
     */

	// スライスの宣言は配列の場合と異なり、[]の中に要素数の記述は不要です
	var hoge1 []int
	var hoge2 = []int{1, 2, 3}
	hoge3 := []int{1, 2, 3}

	fmt.Println(hoge1)
	fmt.Println(hoge2)
	fmt.Println(hoge3)

	// 要素数がわかっている場合にはその値を[]の中に入れる。
	str1 := [5]string{"Taro", "Jiro", "Saburo", "Shiro", "Goro"}
	str2 := str1[:]
	str3 := str1[2:4] // (2+1)番目から(4-1)番目まで
	str4 := str1[2:]  // (2+1)番目から最後まで

	fmt.Println(str1) // [Taro, Jiro, Saburo, Shiro, Goro]
	fmt.Println(str2) // [Taro, Jiro, Saburo, Shiro, Goro]
	fmt.Println(str3) // [Saburo, Shiro]
	fmt.Println(str4) // [Saburo, Shiro, Goro]

	str1[1] = "Second"
	fmt.Println(str1) // [Taro, Second, Saburo, Shiro, Goro]

}
