/*
 * mapのループ処理
 */
package main

import "fmt"

func main(){

	m := map[string]int{"foo": 100, "bar": 200, "baz": 300}

	// キーのみ取り出す
	fmt.Println("Retries only key")
	for key := range m {
		fmt.Println(key)
	}

	// キーと値
	fmt.Println("Retries key and value")
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 値のみ必要な場合
	fmt.Println("Retries only value")
	for _, value := range m {
		fmt.Println(value)
	}


}
