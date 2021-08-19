/*
 * Fatalによるサンプル
 */
package main

import (
    "log"
    "os"
)

func main() {

  _, err := os.Open("/root/hello-world.txt")  // hello-world.txt ファイルを開く
  if err != nil {
    log.Print("err occured")
    log.Fatal(err)  // log.Fatalの場合には、エラーメッセージを表示してプログラムを終了する
  }
  log.Print("Finished!")

}
