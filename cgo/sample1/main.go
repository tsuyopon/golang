/*
 * import "C"の後のコメント部分がC言語を利用するためのコードです。
 * LDFLAGSでコンパイル時に使用するC言語ライブラリの情報を与える。-Lがライブラリの場所を指定するオプション。
 * これらは、gcc -o main main.c -L. -lhelloで指定した2つの引数と同じ意味を持ちます。
 */
package main

/*
#cgo LDFLAGS: -L. -lhello
#include <hello.h>
*/
import "C"

func main() {
    C.hello()
}
