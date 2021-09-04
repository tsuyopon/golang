/*
 * 独自で定義した構造体に対してjson設定ファイルを適用させる。
 */
package main

import (
    "log"
    "os"
    "strconv"
    "sample/config"
)

// jsonの設定ファイルを読み込むサンプル
func main() {
    // 処理開始
    log.Println("start")

    // 構造体へconfigの結果を設定
    c, err := config.Read("config.json")
    if err != nil {
        log.Println(err)
        os.Exit(1)
    }

    // 結果出力
    log.Println("url   = " + c.Url)
    log.Println("count = " + strconv.Itoa(c.Count))
    log.Println("user  = " + c.User)
    log.Println("pass  = " + c.Pass)

    // 処理終了
    log.Println("end")
}
