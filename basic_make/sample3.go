/*
 * copy関数を使うサンプル
 */
package main

import (
  "fmt"
)

func main() {
  a := []string{"Haruka", "Chihaya", "Miki", "Yayoi", "Iori"}
  b := make([]string, 5, 10)
  cLen := copy(b, a)
  fmt.Println(a)    // [Haruka Chihaya Miki Yayoi Iori]
  fmt.Println(b)    // [Haruka Chihaya Miki Yayoi Iori]
  fmt.Println(cLen) // 5

  b[0] = "Ritsuko"
  fmt.Println(a)    // [Haruka Chihaya Miki Yayoi Iori]
  fmt.Println(b)    // [Ritsuko Chihaya Miki Yayoi Iori]
}

