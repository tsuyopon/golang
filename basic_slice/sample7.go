/*
 * スライスの再構成、指定要素の削除
 */
package main

import (
    "fmt"
)

func main() {

    // case1
    fmt.Println("\ncase1: Sliceの抽出")
    s := []string{"A", "B", "C", "D", "E"}

    // 範囲を指定して取得する。[start:end]
    fmt.Println(s[1:4]) // 1 番目から 3 番目まで [B C D]
    fmt.Println(s[2:])  // 2 番目から 最後(4番目まで) [C D E]
    fmt.Println(s[:3])  // 0 番目から 2 番目まで [A B C]
    fmt.Println(s[:])   // 最初から最後まで [A B C D E]


    // case2
    fmt.Println("\ncase2: 先頭の値を取り出して残りを元のスライスに入れる")
    s2 := []string{"A", "B", "C", "D", "E"}
    v2_res, s2_res := s2[0], s2[1:]
    fmt.Println(v2_res, s2_res) // A [B C D E]


    // case3
    fmt.Println("\ncase3: 末尾の値を取り出して残りを元のスライスに入れる")
    s3 := []string{"A", "B", "C", "D", "E"}
    v3_res, s3_res := s3[len(s3)-1], s3[:len(s3)-1]
    fmt.Println(v3_res, s3_res)


    // case4
    fmt.Println("\ncase4: 先頭、末尾以外の値を取り出して残りを元のスライスに入れる")
    s4 := []string{"A", "B", "C", "D", "E"}
    v4_res := s4[2]
    s4 = append(s4[:2], s4[3:]...)
    fmt.Println(v4_res, s4) // C [A B D E]


    // case5
    fmt.Println("\ncase5: 値の削除")
    s5 := []string{"A", "B", "C", "D", "E"}
    // 先頭を削除
    s5 = s5[1:]
    fmt.Println(s5) // [B C D E]
    // 末尾を削除
    s5 = s5[:len(s5)-1]
    fmt.Println(s5) // [B C D]
    // 途中の値を削除
    s5 = append(s5[:1], s5[2:]...)
    fmt.Println(s5) // [B D]


    // case6
    fmt.Println("\ncase6: 容量を指定して削除")
    s6 := []string{"A", "B", "C", "D", "E"}
    s6 = s6[:1]
    fmt.Println(s6, len(s6), cap(s6)) // [A] 1 5


    // case7
    fmt.Println("\ncase7: スライス自体の削除(キャパシティも0にする)")
    s7 := []int{10, 20, 30}
    s7 = nil
    fmt.Println(s7, len(s7), cap(s7)) // [] 0 0


    // case8
    fmt.Println("\ncase8: スライス自体の削除(キャパシティはそのまま確保しておく)")
    s8 := []int{10, 20, 30}
    s8 = s8[:0]
    fmt.Println(s8, len(s8), cap(s8)) // [] 0 3

}
