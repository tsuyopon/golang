/*
 * Interface中には複数のメソッドが定義されているケースです。
 * 以下の例ではIにA, B, C, D４つのメソッドを持たなければならないことを宣言しています。
 * mainの中では、Sを定義してこれらはA, B, C, Dを持っているので、Interface Iの要件を満たしていれば正常に実行されます。(満たしていなければエラーになります)
 */
package main

// このインタフェースIで定義したメソッドを保持していない場合には、エラーとなります。
type I interface {
    A(string)
    B(name string) error
    C(age int, name string)
    D() (n int, err error)
}

type S struct {
}

func (s S) A(name string) {
    println(name)
}

func (s S) B(name string) error {
    return nil
}

func (s S) C(age int, name string) {
}

func (s S) D() (n int, err error) {
    return 0, nil
}

func main() {
    var i I
    var s S

    // 変数sは"type I interface"で定義された関数を持っていることで、Iとしての要件を満たす。これを満たさないとエラーになる。
    // 適当なメソッドをコメントした際にエラーとなることを確認するのが良い。
    i = s
    i.A("hello") // hello
}
