/*
 *
 * appendを利用すると容量がいっぱいになった時、 現在の容量を倍にして確保してくれる。
 */
package main

import (
  "fmt"
)

func main() {
  a := []string{"Haruka", "Chihaya", "Miki", "Yayoi", "Iori"}
  fmt.Println(a)      // [Haruka Chihaya Miki Yayoi Iori]
  fmt.Println(len(a)) // 5
  fmt.Println(cap(a)) // 5

  a = append(a, "Hibiki", "Takane")
  fmt.Println(a)      // [Haruka Chihaya Miki Yayoi Iori Hibiki Takane]
  fmt.Println(len(a)) // 7
  fmt.Println(cap(a)) // 10

  a = append(a, "Makoto", "Yukiho", "Ritsuko", "Azusa")
  fmt.Println(a)      // [Haruka Chihaya Miki Yayoi Iori Hibiki Takane Makoto Yukiho Ritsuko Azusa]
  fmt.Println(len(a)) // 11
  fmt.Println(cap(a)) // 20
}
