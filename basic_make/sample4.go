/*
 * Go言語での1次元配列のmap生成方法
 *    https://qiita.com/suin/items/7225ab9f2aeb6f55c606
 */
package main

import (
  "fmt"
)

func main() {

	// 長い書き方
	var languages1 map[string]string = map[string]string{"go":"golang", "rb":"ruby", "js":"javascript"}
	fmt.Println(languages1)  // map[go:golang js:javascript rb:ruby]

	// 短い書き方
	var languages2 = map[string]string{"go":"golang", "rb":"ruby", "js":"javascript"}
	fmt.Println(languages2)  // map[go:golang js:javascript rb:ruby]

	// 改行する場合の書き方
	var languages3 = map[string]string{
		"go":"golang",
		"rb":"ruby",
		"js":"javascript",
	}
	fmt.Println(languages3)  // map[go:golang js:javascript rb:ruby]

	// 1つずつ指定する場合の書き方
	var languages4 = make(map[string]string)
	languages4["go"] = "golang"
	languages4["rb"] = "ruby"
	languages4["js"] = "javascript"
	fmt.Println(languages4)  // map[go:golang js:javascript rb:ruby]

}

