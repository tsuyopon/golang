/*
 * 配列を関数に渡した場合は、スライスと異なり値渡しであることを理解する
 */
package main;

import (
    "fmt"
)

// 配列の場合には値渡しであることを理解する
func ChangeFirst(people [3]string) {
  people[0] = "Sirou"
}


func main(){
	peopleArray := [3]string{
	  "Taro",
	  "Jiro",
	  "Saburou",
	}

	// 関数の中で、配列の1つ目の要素"Taro"を"Sirou"に変更しようとしている
	ChangeFirst(peopleArray)

	//  関数の外には変更が反映されない
	fmt.Println(peopleArray[0])         // Taro

}
