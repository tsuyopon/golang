/*
 * 様々なmapの型の例
 */
package main

import "fmt"

func main(){

	// 例1: rune型
	fmt.Println("\n### Example1");
	m1 := map[rune]int{}      // rune(int32)
	m1['あ'] = 100
	fmt.Println(m1['あ'])     // 100


	// 例2:  真偽値型
	fmt.Println("\n### Example2");
	m2 := map[bool]string{
		true:  "value-1",
		false: "value-2",
	}
	println(m2[true])         // value-1


	// 例3:  ポインター
	fmt.Println("\n### Example3");
	m3 := map[*string]int{}
	key1 := "key"
	key2 := "key"
	key3 := "key"
	m3[&key1] = 100
	m3[&key2] = 200
	m3[&key3] = 300

	for k, v := range m3 {
		fmt.Println(*k, v)
	}


	// 例4:  インターフェース
	fmt.Println("\n### Example4");
	m4 := map[interface{}]int{}
	m4["key"] = 100
	m4[true] = 200
	m4[300] = 300
	println(m4["key"]) // 100
	println(m4[true])  // 200
	println(m4[300])   // 300


	// 例5:  構造体
	fmt.Println("\n### Example5");
	type Key struct {
		name string
	}

	m5 := map[Key]int{}
	m5[Key{"foo"}] = 100
	m5[Key{"bar"}] = 200
	m5[Key{"baz"}] = 300

	for k, v := range m5 {
		println(k.name, v)
	}

}
