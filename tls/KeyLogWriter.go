/*
 * wiresharkに解釈させるためのkeylogfileを出力させるサンプルプログラム
 * tls.ConfigにKeyLogWriter: wを指定するだで出力させることができます。
 *
 * inspired by
 *   https://golang.org/pkg/crypto/tls/#example_Config_keyLogWriter
 */
package main

import (
//	"fmt"
//	"io/ioutil"
	"crypto/tls"
	"log"
	"net/http"
	"os"
)

func main() {
	w := os.Stdout

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				// KeyLogWriterがwiresharkで解釈させるSSLKEYLOGFILEフォーマットの出力となる。
				KeyLogWriter: w,
				InsecureSkipVerify: true,         // test server certificate is not trusted.
			},
		},
	}
	resp, err := client.Get("https://www.yahoo.co.jp")
	if err != nil {
		log.Fatalf("Failed to get URL: %v", err)
	}

//	以下で取得した内容を出力する
//	byteArray, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println(string(byteArray))
	resp.Body.Close()

}
