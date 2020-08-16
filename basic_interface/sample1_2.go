/*
 * sample1_1.goではinterface{}はどのような型でも利用できることを説明しました。
 * sample1_2.goでは型アサーションとして利用できるinterface{}について説明しています。
 *
 * 型アサーションは以下の構文で表されます。2つの記法があります。
 *      1) value := <変数>.(<型>)
 *      2) value, ok := <変数>.(<型>)
 *
 * 1)だと型アサーションでエラーとなった場合にはpanic終了し、2)では型アサーションでエラーとなってもpanic終了しません。
 *
 * このプログラムではそれを確認します。
 *
 */
package main

import "fmt"

func main() {

    // 型アサーションでinterface{}宣言をしている
    var hellovar interface{} = "hello"

    variable := hellovar.(string)
    fmt.Println(variable)             //=> hello

    // 1番目の変数variableには型アサーション成功時に実際の値が格納され、2番目の変数okには型アサーションの成功の有無（true/false）が格納されます。
    variable, ok := hellovar.(string)
    fmt.Println(variable, ok)         //=> hello true


    fmt.Println("\n\nこの後は成功の有無を確認しているので、interfaceの型アサーションがエラーでもpanic終了しない\n");
    float, ok := hellovar.(float64) 
    fmt.Println(float, ok)            //=> 0 false
    //格納失敗が、成功したかの有無を確かめるokが存在するのでエラーにはならない。


    fmt.Println("\n\nこの後は成功の有無を確認していないので、interfaceの型アサーションがエラーでpanic終了になる\n");
    float = hellovar.(float64) 
    fmt.Println(float) //=> panic: interface conversion: interface {} is string, not float64
    //成功したかの有無を確かめるokが存在しないのでエラーが発生する。
}
