/*
 * StudentとTeacherが共通するInterfaceであるPersonを定義する。
 * StudentとTeacherのemail取得関数getEmailではどちらの構造体で定義されているかによって出しわけられる。
 */
package main

import "fmt"

type Student struct {
	Name   string
	Number int
	Grade  int
}

type Teacher struct {
	Name string
}

type Person interface {
	getEmail() string
}

func (s Student) getEmail() string {
	return s.Name + "@student.jp"
}
func (t Teacher) getEmail() string {
	return t.Name + "@teacher.jp"
}

func sendEmail(p Person) (context string) {
	from := p.getEmail()
	context = `
  送信元 : ` + from + `です`
	return context
}

func main() {
	var s, t Person       // interface Personとして定義

	// interface PersonにはそれぞれStudent, Teacherの構造体として定義
	s = Student{
		Name:   "Tarou",
		Number: 100,
		Grade:  1,
	}
	t = Teacher{
		Name: "Satou",
	}

	// sendEmailは定義したクラスによってgetEmail()の戻り値が異なります。
	cxtStu := sendEmail(s)
	fmt.Println(cxtStu)
	cxtTea := sendEmail(t)
	fmt.Println(cxtTea)
}

