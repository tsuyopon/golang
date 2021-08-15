/*
 * 構造体の宣言
 */

package main

import "fmt"

// typeとstructを使用して定義します。
// go言語の場合、大文字から始まる名前 (関数名、型名、フィールド名) は、他のパッケージからアクセス可能となります(public扱い)
// User構造体定義。大文字から始まっているので他パッケージからアクセス可能
type User struct {
  name string
  age  int
}

func main() {

	//User構造体にアクセス
	var u User
	u.name = "taro"
	u.age = 30
	fmt.Printf("name:%s", u.name)

}
