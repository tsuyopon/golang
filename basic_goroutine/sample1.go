/*
 * ゴルーチン(goroutine)の単純なサンプルです
 */
package main

import (
	"fmt"
	"time"
)

func funcA() {
    for i := 0; i < 10; i++ {
        fmt.Print("funcA\n")
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {
    // ゴルーチン(goroutine)はGo言語における並行処理を実現するもので、スレッドよりも高速に並行処理を実現することができます
    go funcA()
    for i := 0; i < 10; i++ {
        fmt.Print("main\n")
        time.Sleep(20 * time.Millisecond)
    }
}
