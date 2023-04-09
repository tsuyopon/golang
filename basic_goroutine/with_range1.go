/*
 * チャネルをfor rangeする際の注意点。
 * ちゃんとチャネルをcloseしないとデッドロックを引き起こします。
 */
package main

import "fmt"

// チャネルに値を渡す
func sendValue(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

	// 以下の行をコメントしてからgo runを実行してみてください。deadlockとエラー画面が表示されます。
	close(c)
}

// チャネルから値を受け取る
func receiveValue(c <-chan int) {
	for v := range c {
		fmt.Println("チャネルから受け取った値：", v)
	}
	fmt.Println("receiveValueはチャネルがcloseされたので終了します")
}

func main() {
	c := make(chan int)
	go sendValue(c)
	receiveValue(c)

	fmt.Println("処理を終了します。")
}
