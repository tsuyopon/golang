/*
 * スライスを関数に渡した場合は、配列と異なり参照渡しであることを理解する
 */
package main;

import (
    "fmt"
)

// スライスの場合には参照渡しであることを理解する
func ChangeFirst(people []string) {
  people[0] = "Sirou"
}


func main(){
	peopleSlice := []string{
	  "Taro",
	  "Jiro",
	  "Saburou",
	}

	// 関数の中で、スライスの1つ目の要素"Taro"を"Sirou"に変更しようとしている
	ChangeFirst(peopleSlice)

	// 参照渡しのため、配列とは異なりスライスの場合は関数の外には変更が反映される
	fmt.Println(peopleSlice[0])         // Sirou

}
