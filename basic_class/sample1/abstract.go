package main

import "fmt"

// 親構造体
type parentStruct struct {
    x int
    y int
}

// 親構造体の関数
func (p *parentStruct) testParent() {
    fmt.Println("Parent", p.x, p.y)
}

// 子構造体
type childStruct struct {
    z int
    parentStruct  // 親構造体
}

// 子構造体の関数
func (p *childStruct) testChild() {
    // 親構造体の変数にアクセス
    fmt.Println("Child", p.parentStruct.x, p.parentStruct.y, p.z)  
}

func main() {
    var a childStruct
    a.parentStruct.x = 1  // 親の変数にアクセス
    a.parentStruct.y = 2  // 親の変数にアクセス
    a.z = 3  // 子の変数にアクセス
    a.parentStruct.testParent() // Parent 1 2  親の関数にアクセス
    a.testChild()               // Child 1 2 3  子の関数にアクセス
}
