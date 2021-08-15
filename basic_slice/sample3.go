/*
 * 
 */
package main

import (
    "fmt"
)

/* ()内は注意点を説明 */
func main() {

	// スライスの宣言 (配列は固定長、スライスは可変長。 要素の各値を指定する場合は配列の生成と激似なので注意)
	slice := []int{1, 2, 3}

	// スライスへの要素の追加 (配列はappend()で拡張できない、スライスはできる)
	slice = append(slice, 4)
	fmt.Println(slice)

	// スライスへの値の変更
	slice[0] = 5
	fmt.Println(slice)

	// スライスのコピー (配列は変数への代入でコピーが作れるが、スライスはcopy()を使用しないと作れない)
	slice2_src := []int{1, 2, 3}
	slice2_dst := make([]int, len(slice2_src))
	copy(slice2_dst, slice2_src)   // copy()を使用する注意点としては、コピー先のスライスのサイズをコピー元に合わせないとエラーになってしまうようです
	fmt.Println(slice2_dst)

	// 配列からスライスを作成する
	array := [5]int{1, 2, 3, 4, 5}
	slice3 := array[:]
	fmt.Println(slice3)

}
