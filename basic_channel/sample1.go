/*
 * チャネル(chan)の利用と受信演算子(<-)によりチャネルからのメッセージを取得するサンプルです
 */
package main

import (
	"fmt"
	"time"
)

/*
 *  funcAでは "chan<-" (受信チャネル)、 mainでは "<-chA" (送信チャネル) となっている点に注意してください。
 */
func funcA(chA chan <- string) {  // chanは受信チャネルを表します。 この関数の引数として、string型の受信チャネルchAを定義しています。
    time.Sleep(3 * time.Second)
    chA <- "Finished"             // チャネルにメッセージを送信する ( <- は受信演算子)
}

func main() {
    chA := make(chan string)      // チャネルを作成する
    defer close(chA)              // 使い終わったらクローズする
    go funcA(chA)                 // チャネルをゴルーチンに渡す

    // Goでは受信側では常に、受信可能なデータが来るまでブロックされます。
    msg := <- chA                 // チャネルからメッセージを受信する (チャネルのメッセージを受信する場合には <-chAとして"<-"を付与する。これは受信演算子と呼ぶ。chan<-や<-chanは受信演算子ではないのでその点、混乱しがちなので注意すること
    fmt.Println(msg)
}
