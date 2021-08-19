/*
 * 犬と猫は「こんにちは」だけ共通で話せるとします。
 * その時に、DogStructとCatStructの構造体で共通の関数を定義するにはどうすれば良いでしょうか?
 *
 * 次の手順で実装します。(コードはsample2_6.goからの改良です)
 *    1. ポイント(1)でCommonStructをDogStructとCatStructの構造体に含めます。
 *    2. ポイント(2)でCommonStructの構造体とそのメソッドチェーンcommon_bark()を定義します。
 *    3. ポイント(3)でAnimalのinterface中にcommon_bark()を定義します。
 * 以上でDogStructとCatStructで共通で使える共通関数を定義できるようになりました。
 * 
 * interfaceでは共通関数は定義できないので、DogStructやCatStructで共通のstructを作るようにする
 * cf. https://stackoverflow.com/questions/31505587/how-can-two-different-types-implement-the-same-method-in-golang-using-interfaces/31505807
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
 * 共通の構造体  (** ポイント(2) **)
 */
type CommonStruct struct {
}
func (c *CommonStruct)common_bark(){
    fmt.Println("こんにちは")
}

/*
 * 犬の構造体
 */
type DogStruct struct{
    CommonStruct             /** ポイント(1) **/
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
    CommonStruct             /** ポイント(1) **/
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
    bark_angry()
    common_bark()            /** ポイント(3) **/
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

    // 共通の鳴き声「こんにちは」
    BaseStruct.Animals[0].common_bark()
    BaseStruct.Animals[1].common_bark()
}
