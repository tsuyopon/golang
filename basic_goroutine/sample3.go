/*
 * 関数の引数でチャネルを受け取る(Directions)
 *
 * 参考資料: https://qiita.com/ta1m1kam/items/fc798cdd6a4eaf9a7d5e
 */

// 送信として定義しているので「chan<-」として定義されている
func ping(pings chan<- string, msg string) {
    // msgをpingsに送信しているので送信チャネル
    pings <- msg
}

// 下記関数内では2つのチャネルが使われていて
// 1つめのチャネルがpingsを受信したらmsgに格納しています。
// 2つめのチャネルがmsgを受信したらpongsチャネルに書き込んでいます。
// pong関数の引数定義はpingsは「<-chan」、pongsは「chan<-」と異なっていることに注意してください。
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "Hello")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
