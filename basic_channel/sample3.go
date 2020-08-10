/*
 * mainからチャネルの送受信の両方を行ないます。
 * 
 * 参考: https://qiita.com/TsuyoshiUshio@github/items/d94a3d0f934bde6d6aed
 */
package main

import (
    "fmt"
)

// "chan<-" なので引数は受信チャネルが定義
func receive(c chan<- string, message string) { // channel for send only
    c <- message
}

// "<-chan" なので引数は送信チャネルが定義
func send(c <-chan string) { // channel for recieve only
    fmt.Println(<-c)              // チャンネルからの値を取得するので <-c とする。この演算子は受信演算子と呼ばれる
}

func main() {

    c := make(chan string, 1)
    // c <- "hello then"
    receive(c, "hello")
    fmt.Println(<-c)              // チャンネルから値を取得する場合には <-c とする。この演算子は受信演算子と呼ばれる(重要なので2回目)

    c = make(chan string, 1)
    c <- "world"                 // 送信する値を設定する際には「c <- "message」とする
    send(c)
}
