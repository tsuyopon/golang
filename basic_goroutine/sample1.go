/*
 * ゴルーチン(goroutine)の単純なサンプルです
 * 
 * goroutine (ゴルーチン)は、Goのランタイムに管理される軽量なスレッドです。
 * goroutineは、同じアドレス空間で実行されるため、共有メモリへのアクセスは必ず同期する必要があります。
 */
package main

import (
	"fmt"
	"time"
	"log"
	"runtime"
)

func funcA(threadid int) {
    for i := 0; i < 10; i++ {
        fmt.Printf("funcA: threadid=%d, num=%d\n", threadid, i)
        time.Sleep(10 * time.Millisecond)
    }
}

func main() {

    // ゴルーチン(goroutine)はGo言語における並行処理を実現するもので、スレッドよりも高速に並行処理を実現することができます
    go funcA(1)
    go funcA(2)

    // 現在起動しているgoroutine(ゴルーチン)の数を知ることが可能
    log.Println(runtime.NumGoroutine()) 

    for i := 0; i < 10; i++ {
        fmt.Print("main\n")
        time.Sleep(20 * time.Millisecond)
    }

    // 1秒待って、他のスレッドが
    time.Sleep(1000 * time.Millisecond)

    // 現在起動しているgoroutine数が少なくなったことを表示
    log.Println(runtime.NumGoroutine()) 

}
