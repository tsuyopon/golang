/*
 * goroutineの終了方法
 * 参考: https://jxck.hatenablog.com/entry/20130414/1365960707
 */
package main

import (
    "log"
    "time"      // for runtime.Goexit()
    "runtime"   // for time.Sleep()
)


func main() {

	// 関数が終わる
	go func() {
		log.Println("end")
	}()

	// returnを記述する
	go func() {
		log.Println("return")
		return
	}()

	// runtime.Goexit()で終わる
	go func() {
		log.Println("exit")
		runtime.Goexit()
	}()

	// goroutineの数を表示する
	log.Println("before sleep count:", runtime.NumGoroutine())

	time.Sleep(time.Second)

	// goroutineの数を表示する
	log.Println("after sleep count:", runtime.NumGoroutine())
}
