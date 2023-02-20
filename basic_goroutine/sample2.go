/*
 * goroutineがデッドロックされるサンプルです
 * このようなプログラムを組むと良くない例です。
 */
package main

import (
	"fmt"
)

func main() {

    // 受信バッファとして2つのチャネルしか定義していない
    ch := make(chan int, 2)

    // 受信バッファ2つに大して、それ以上の3つを送っているのでdeadlockが発生する。
    // 下記エラーが出力されます
    //    fatal error: all goroutines are asleep - deadlock!
    ch <- 1
    ch <- 2
    ch <- 3

    fmt.Println(<-ch)
    fmt.Println(<-ch)

}
