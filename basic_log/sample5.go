/*
 * sample1.goで設定を行うと、全体のログに適用されてしまいます。
 * 一部のカスタムログだけに対して、設定を適用させたい場合にはlog.New関数によってLoggerオブジェクトを生成して設定を行います。
 */
package main

import (
    "log"
    "os"
)

var wlog *log.Logger

func init(){
  wlog = log.New(os.Stdout, "[PREFIX]", log.LstdFlags|log.LUTC)
}

func main(){
  log.Print("Hello, world!")         // 普通のログ
  wlog.Print("Hello, Waman world!")  // カスタムログ
}
