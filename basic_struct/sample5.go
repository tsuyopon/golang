/*
 * structとinterfaceを合わせたサンプルです。
 * main側でuser変数としてinterfaceとしてJSONableを指定して、Userの型を格納しています。
 * これによって、User型は必ずJSON()を実装しなければなりません。
 *
 * この例では親のUserのJSON()を呼び出す方法、子となるAdminUserのJSON()がオーバーライドされて呼び出す方法のサンプルとなります。
 *
 */

package main

import "fmt"

// 「JSON化できる」インターフェース
type JSONable interface {
    JSON() string
}

type User struct {
    Id   int
    Name string
}

// JSON()メソッドを実装
func (s *User) JSON() string {
    return fmt.Sprintf(`{ "Id": %d, "Name": "%s" }`, s.Id, s.Name)
}

type AdminUser struct {
    User
    Admin bool
}

// User.JSON()をオーバーライド
func (s *AdminUser) JSON() string {
    return fmt.Sprintf(`{ "Id": %d, "Name": "%s", "Admin": %v }`, s.Id, s.Name, s.Admin)
}


func main() {
    // JSONable を実装しているのでJSONable型に代入できる
    var user JSONable = &User{1, "oinume"}
    fmt.Println(user.JSON())

    // AdminUserもJSONableを実装している
    var adminUser JSONable = &AdminUser{User{0, "admin"}, true}
    fmt.Println(adminUser.JSON())

    // ★型アサーションでadminUserがJSONableを実装しているかチェックする
    jsonable, ok := adminUser.(JSONable)
    if ok {
        // 実装してる場合
        fmt.Printf("JSON(): %s\n", jsonable.JSON())
    } else {
        // 実装してない場合
        fmt.Printf("Not JSONable\n")
    }
}
