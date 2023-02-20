/*
 * 関数側でチャネルを受信するまで処理を実施しない
 *
 * 参考: https://jxck.hatenablog.com/entry/20130414/1365960707
 */
package main

import (
    "log"
)

func f(yield chan string) {
	yield <- "one"
	yield <- "two"
	yield <- "three"
}

func main() {
	co := make(chan string)
	go f(co)
	log.Println(<-co) // one
	log.Println(<-co) // two
	log.Println(<-co) // three
}
