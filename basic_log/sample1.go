/*
 * ログ出力の簡単なサンプル
 */
package main

import (
    "log"
    "os"
)

func main() {

    // 標準のログ出力
    log.Println("Hello, World")

    // ログに含める情報を指定する
    /*
     * const (
     *   Ldate         = 1 << iota  // 日付
     *   Ltime                      // 時刻
     *   Lmicroseconds              // 時刻のマイクロ秒
     *   Llongfile                  // ソースファイル（ディレクトリパスを含む）
     *   Lshortfile                 // ソースファイル（ファイル名のみ）
     *   LUTC                       // タイムゾーンに依らない UTC 時刻
     *   LstdFlags     = Ldate | Ltime  // 日付 (Ldata) と時刻 (Ltime) ：デフォルト
     * )
    */
    log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

    // 出力先は標準出力
    log.SetOutput(os.Stdout)

    // 先頭に付与するprefixの文字列
    log.SetPrefix("[Hello] ")

    // 標準のログ全てに対して、上記の設定が適用される。(カスタムログについてはsample5.goを参照のこと)
    log.Println("Hello, World")
}
