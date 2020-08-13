/*
 * 構造体の初期化の様々な方法について
 */
package main

import "fmt"

type User struct {
  name string
  age  int
}

// Goだと、構造体は下記のような初期化関数を作るのが一般的らしいです。
func newUser(name string, age int) *User {
    u := new(User)
    u.name = name
    u.age = age
    return u
}

func main() {

	// 宣言後にフィールド設定して初期化する方法(sample1.goと比較してみてください)
	var u User
	u.name = "taro"
	u.age = 30
	fmt.Printf("name: %s\n", u.name)

	// 順番で初期化
	u2 := User{"taro", 30}
	fmt.Printf("name: %s\n", u2.name)

	// Key:Valueで初期化。
	u3 := User{name: "taro", age: 30}
	fmt.Printf("name: %s\n", u3.name)

	// Key:Valueで初期化。u4はポインタ型
	u4 := &User{name: "taro", age: 30} // var u4 *User
	fmt.Printf("name: %s\n", u4.name)

	// newで初期化.u5はポインタ型
	u5 := new(User) // var u5 *User
	u5.name = "taro"
	fmt.Printf("name: %s\n", u5.name)

	// 初期化関数newUserを使って初期化する
	var u6 *User = newUser("taro",30)
	fmt.Printf("name: %s\n", u6.name)

}
