/*
 * iniをパースするサンプル。"gopkg.in/ini.v1"を使っています。
 */
package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type ConfigList struct {
    Apikey    string
    Apisecret string

}
var Config ConfigList

func main() {

  initialize()

  //設定値を取り込む
  var apikey string
  var apisecret string

  apikey = Config.Apikey
  apisecret = Config.Apisecret

  fmt.Printf("%#v\n", apikey)
  fmt.Printf("%#v\n", apisecret)
}

func initialize() {
    cfg ,_ := ini.Load("config.ini")

    Config = ConfigList{
        Apikey:    cfg.Section("api").Key("api_key").String(),
        Apisecret: cfg.Section("api").Key("api_secret").String(),
    }
}

