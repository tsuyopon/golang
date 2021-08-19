/*
 * 継承とabstractライクなことを実現するためのサンプル
 * 
 * Animal型はBaseをinterfaceしているので、Baseがinterfaceするbark()を実装する必要があります。
 */
package main

import(
    "fmt"
)

// Animal構造体
// この中にインターフェースのスライスを定義して、その中に他の動物の構造体を入れる
type AnimalStruct struct{
    Animals []Animal
}

/*
 * 犬の構造体
 */
type DogStruct struct{
    no int
}
func (d *DogStruct)bark(){
    fmt.Println("わんわん")
}
func (d *DogStruct)bark_angry(){
    fmt.Println("ぐるるる")
}

/*
 * 猫の構造体
 */
type CatStruct struct{
    no int
}
func (c *CatStruct)bark(){
    fmt.Println("にゃーにゃー")
}
func (c *CatStruct)bark_angry(){
    fmt.Println("ぎゃー")
}

/*
 * Base interface
 */
type Base interface{
    bark()
}

/*
 * Animal interface
 */
type Animal interface{
    Base
    bark_angry()      // Animalのみ専用の関数
}

func main(){
    // ベースとなる構造体を定義
    BaseStruct := AnimalStruct{}

    // 犬の構造体と猫の構造体を生成
    dog := DogStruct{no: 100}
    cat := CatStruct{no: 101}

    // それぞれ追加
    BaseStruct.Animals = append(BaseStruct.Animals, &dog)
    BaseStruct.Animals = append(BaseStruct.Animals, &cat)

    //使ってみる
    BaseStruct.Animals[0].bark()
    BaseStruct.Animals[1].bark()

    BaseStruct.Animals[0].bark_angry()
    BaseStruct.Animals[1].bark_angry()

    // これはできない　Interfaceはメソッドだけ持つ
    // BaseStruct.Animals[0].no
    // BaseStruct.Animals[1].no

    // こうやったら使えるようになる
    fmt.Println(BaseStruct.Animals[0].(*DogStruct).no)
    fmt.Println(BaseStruct.Animals[1].(*CatStruct).no)
}
