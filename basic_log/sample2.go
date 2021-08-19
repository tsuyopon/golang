/*
 * ログ出力の簡単なサンプル
 */
package main

import (
    "log"
)

func main() {

  log.Print("Hello, world!")
  log.Print("Hello, world2!\n")  // 最後に改行をつける
  log.Print("Hello", ", " , "world", 3, "!")  // 複数の引数を与える

  log.Println("Hello, world!")
  log.Println("Hello, world2!\n")
  log.Println("Hello", ", " , "world", 3, "!")

  log.Printf("Hello, %s%d!", "world", 1)
  log.Printf("Hello, world%d!\n", 2)
  log.Printf("Hello, world3!")

}
