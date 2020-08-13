/*
 * jsonをパースするだけの簡単なサンプル
 * 
 * 参考: https://www.bravesoft.co.jp/blog/archives/10023
 */
package main

import (
   "encoding/json"
   "fmt"
   "log"
)

// 構造体の型を事前に宣言しておく必要がある
type (
   Company struct {
       Name    string   `json:"name"`
       Students []StudentInfo`json:"student"`
   }
   StudentInfo struct {
       ID   int    `json:"id"`
       Name string `json:"name"`
       Age  int    `json:"age"`
   }
)

func main() {
   jsonStr := `{"name": "国語","student": [{"id": 1,"name": "太郎","age": 80},{"id": 2,"name": "二郎","age": 90}]}`

   jsonBytes := ([]byte)(jsonStr)
   companyData := new(Company)
   if err := json.Unmarshal(jsonBytes, &companyData); err != nil {
       log.Fatal(err)
   }

   for _, info := range companyData.Students {
       fmt.Println(info.Name)
   }
}
