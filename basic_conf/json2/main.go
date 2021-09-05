/*
 * JSONを取得するサンプル
 * 定義された型を意識しないようにinterface{}が利用されている
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONの読み込み
func loadJson(inputPath string) interface{} {
	byteArray, _ := ioutil.ReadFile(inputPath)
	var jsonObj interface{}
	_ = json.Unmarshal(byteArray, &jsonObj)
	return jsonObj
}

// JSONの書き出し
func saveJson(jsonObj interface{}, outputPath string) {
	file, _ := os.Create(outputPath)
	defer file.Close()
	_ = json.NewEncoder(file).Encode(jsonObj)
}

func main() {

	// JSON読み込み
	jsonObj := loadJson("sample.json")

	// KeySchema -> 0 -> KeyTypeを表示
	fmt.Println(jsonObj.(map[string]interface{})["KeySchema"].([]interface{})[0].(map[string]interface{})["KeyType"].(string))

	// KeySchema -> 0 -> KeyTypeに"HOGE"を代入
	jsonObj.(map[string]interface{})["KeySchema"].([]interface{})[0].(map[string]interface{})["KeyType"] = "HOGE"

	// KeySchema -> 0 -> KeyTypeを表示
	fmt.Println(jsonObj.(map[string]interface{})["KeySchema"].([]interface{})[0].(map[string]interface{})["KeyType"].(string))

	// JSON書き出し
	saveJson(jsonObj, "sample.json.bak")
}
