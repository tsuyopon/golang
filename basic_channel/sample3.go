/*
 * チャネルを引数として受け取る関数の使い方 (Channel Direction)
 * 
 * 参考: https://qiita.com/TsuyoshiUshio@github/items/d94a3d0f934bde6d6aed
 */
package main

import (
    "fmt"
)

// 「chan<-」となっているので送信専用チャネルが引数として定義されている
// なお「c chan string」のように関数の引数定義も可能だがその場合は送受信可能となり、関数内から不要な受信を防止できない。
func receive(c chan<- string, message string) {
    // チャネルcにmessageを送信している
    c <- message
}

// 「<-chan」となっているので受信専用チャネルが引数として定義されている
// なお「c chan string」のように関数の引数定義も可能だがその場合は送受信可能となり、関数内から不要な送信を防止できない。
func send(c <-chan string) {
    // 「<-c」はcを受信している
    // 受信演算子ともよばれる
    fmt.Println(<-c)
}

func main() {

    c := make(chan string, 1)
    receive(c, "hello")
    fmt.Println(<-c)              // チャネルから値を取得する場合には <-c とする。この演算子は受信演算子と呼ばれる(重要なので2回目)

    c = make(chan string, 1)
    c <- "world"                 // 送信する値を設定する際には「c <- "message"」とすると、"message"をチャネルcに伝えることができる。
    send(c)
}
