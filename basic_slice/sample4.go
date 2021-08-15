/*
 * スライスの容量とキャパシティを確認するプログラム
 * 参考: https://simple-minds-think-alike.hatenablog.com/entry/golang-array-and-slice
 */
package main

import (
    "fmt"
)

func main() {

	people := make([]string, 3, 4)
	people[0] = "Taro"
	people[1] = "Jiro"
	people[2] = "Saburou"

	fmt.Println("=================== 3 person =====================");
	fmt.Println(people)
	fmt.Println("len",len(people))
	fmt.Println("cap",cap(people))

	fmt.Println("=================== 4 person =====================");
	people = append(people, "Sirou")
	fmt.Println(people)
	fmt.Println("len",len(people))
	fmt.Println("cap",cap(people))

	fmt.Println("=================== 5 person =====================");
	people = append(people, "Gorou")
	fmt.Println(people)
	fmt.Println("len",len(people))
	fmt.Println("cap",cap(people))

	fmt.Println("=================== 8 person =====================");
	people = append(people, "Rokurou")
	people = append(people, "Nanarou")
	people = append(people, "Hachirou")
	fmt.Println(people)
	fmt.Println("len",len(people))
	fmt.Println("cap",cap(people))

	fmt.Println("=================== 9 person =====================");
	people = append(people, "Kurou")
	fmt.Println(people)
	fmt.Println("len",len(people))
	fmt.Println("cap",cap(people))
}
