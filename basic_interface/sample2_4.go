/*
 * ioパッケージのintefaceを使う例です。
 * https://golang.org/src/io/io.go?s=3303:3363#L67
 *
 * ioパッケージには以下のInterfaceが定義されています。
 * type Reader interface {
 *    Read(p []byte) (n int, err error)
 * }
 *
 * この例ではbytesパッケージの
 * bytes.Buffer は Read(p []byte) (n int, err error) を持っているので、io.Reader 型を名乗れる。
 *
 * 参考URL: https://golang.hateblo.jp/entry/golang-interface
 */
package main

import (
    "bytes"
    "io"
)

// io.Readerとして受け取っているのはbytes.Buffer型である。
// しかし、io.Readerのinterfaceには以下のReadが定義されていればよいので、io.Readerとして受け取ることができます。(もちろん自分が定義した関数で受け取ることも可能です)
//   https://golang.org/src/io/io.go?s=3303:3363#L67
//
func Print(r io.Reader) {
    buff := make([]byte, 256)
    for {
        _, err := r.Read(buff)
        if err == io.EOF {
            break
        }
        println(string(buff))
    }
}

func main() {
    var b bytes.Buffer
    b.WriteString("hello world")
    Print(&b)
}
