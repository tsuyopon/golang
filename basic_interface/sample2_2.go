/*
 * sample2_1.goと同じようにtype xxx interfaceしたサンプルです
 */
package main

import "fmt"

type Stringer interface {
	String() string
}

type Hex int

// この関数をコメントするとエラーになる。
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// Hex does not implement Stringer (missing String method)
	var h Stringer = Hex(100)
	fmt.Println(h)
}
