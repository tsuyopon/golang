/*
 * 接頭辞 F
 *     書き込み先を指定
 * 接頭辞 S
 *     結果を文字列で返す
 * 接尾辞 f
 *     フォーマットを指定
 * 接尾辞 ln
 *     オペランド間にスペース、最後に改行を追加
*/
package main

import (
    "fmt"
    "os"
)

func main() {

	// 標準形
	fmt.Print("Helloworld!\n")
	fmt.Print("Hello", "world!\n")

	// はじめに「F」が付いているものは、書き込み先を明示的に指定できます。
	fmt.Fprint(os.Stdout, "Hello prefix_F world!\n")

	// はじめに「S」が付いたものは、出力ではなくフォーマットした結果を文字列で返します。
	shello := fmt.Sprint("Hello prefix_S world!\n")
	fmt.Print(shello)

	// 後ろに「f」が付いたものは、フォーマットを指定することができます。
	hellof := "Hello suffix_f world!"
	fmt.Printf("%s\n", hellof)
	fmt.Printf("%#v\n", hellof)

	// 後ろに「ln」が付いたものは、オペランドの間に半角スペースが入り、文字列の最後に改行が追加されます。
	fmt.Println("Hello", "world!")

}
