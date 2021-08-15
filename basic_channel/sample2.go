/*
 * チャネル(chan)と受信演算子(<-)を利用したサンプルです
 * sample1.goから少し改良して、2つのチャネルを受信したら終了となります。
 */
package main

import (
	"fmt"
	"time"
)

func funcA(chA chan <- string) {
    time.Sleep(1 * time.Second)
    chA <- "funcA Finished"
}

func funcB(chB chan <- string) {
    time.Sleep(2 * time.Second)
    chB <- "funcB Finished"
}

func main() {

    // 組み込み関数make()を使用する事で生成可能
    chA := make(chan string)
    chB := make(chan string)
    defer close(chA)
    defer close(chB)
    finflagA := false
    finflagB := false
    go funcA(chA)
    go funcB(chB)
    for {
        fmt.Print("for loop\n")
        // チャネルオペレータの<-を用いる事で値の送受信が可能
        select {
        case msg := <- chA:
            fmt.Print("chA recieved\n")
            finflagA = true
            fmt.Println(msg)
        case msg := <- chB:
            fmt.Print("chB recieved\n")
            finflagB = true
            fmt.Println(msg)
        }
        if finflagA && finflagB {
            fmt.Print("chA and chB recieved\n")
            break
        }
    }
    fmt.Print("Finished\n")
}
