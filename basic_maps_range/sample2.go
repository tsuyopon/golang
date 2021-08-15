/*
 * mapの要素の挿入、更新、削除、キーの存在確認、要素数の出力
 */
package main

import "fmt"

func main(){

   mapEx := make(map[string]string)   // sample1.goとは異なり要素数を指定しないことも可能

   // 初期化後を出力
   fmt.Println(mapEx)

   // 要素の挿入
   mapEx["hoge1"] = "fuga1"
   mapEx["hoge2"] = "fuga2"
   mapEx["hoge3"] = "fuga3"
   mapEx["hoge4"] = "fuga4"
   mapEx["hoge5"] = "fuga5"

   // 全要素の出力
   fmt.Println(mapEx)

   // 要素数の出力
   fmt.Println("map length: ", len(mapEx))

   // 要素の更新
   mapEx["hoge4"] = "updated4"

   // 要素を指定した出力
   fmt.Println(mapEx["hoge4"])

   // 要素を削除する
   delete(mapEx, "hoge4")
   fmt.Println(mapEx["hoge4"])     // 空が出力される

   // キーの存在確認を行う(２つ目の変数が真偽値になっているのでそれで確認できる。)
   value, isThere := mapEx["hoge4"]
   fmt.Println(value, isThere)   // valueは存在しない, isThereはfalse

   // 定義したmap自体を削除する
   fmt.Print("delete map\n")
   mapEx = nil
   fmt.Println(mapEx)
}
