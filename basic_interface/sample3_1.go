/*
 * interfaceはポリモーフィズムを実現します。
 * 以下ではmonkeyとhumanで共通の関数を定義したい場合に、それぞれのレシーバとしてそれぞれの関数を重複して定義する必要がありません。
 * interfaceとして定義することでmonkeyもhumanもastronautというtypeになれます。
 *
 * (参考) https://thiscalifornianlife.com/2021/01/10/golang-interface/
 */
package main

import "fmt"

// monkey
type monkey struct { Name string }
func (m monkey) imitate(){
  fmt.Println("Imitating")
  }
  func (m monkey) makeADream(){
    fmt.Println("Make a Dream!")
}


// human
type human struct { Name string }
func (h human) talk(){
  fmt.Println("Talking")
  }
  func (h human) makeADream(){
    fmt.Println("Make a Dream!")
}


// 下記の定義の意味は「makeADream()メソッドを持っているものはみんなastronaut interfaceだよ。」ということを意味します。
// 例えばmonkeyとhumanに「goToTheMoon」という共通の関数を追加したい場合に、monkeyとhuman両方に定義を
type astronaut interface {
  makeADream()
}


func main() {
  jiro := monkey{Name:"JIRO"}
  hanako := human{Name:"HANAKO"}

  goToTheMoon(hanako)  //hanakoはhuman typeだけではなくastronaut typeにもなりえるのでエラーにならない。
  goToTheMoon(jiro)    //jiroはmonkey typeだけではなくastronaut typeにもなりえるのでエラーにならない。
}

// 共通で
func goToTheMoon(a astronaut){
  fmt.Println("Go to the moon")
}
