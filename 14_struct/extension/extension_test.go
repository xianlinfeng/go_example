package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	Pet // go提供匿名嵌套类型
}

/* 内嵌的结构类型是不能当成继承来用的，它不支持override和LSP */

func (d *Dog) Speak() {
	fmt.Print("Wang!")
}

// func (d *Dog) SpeakTo(host string) {
// 	d.Speak()
// 	fmt.Println(" ", host)
// }

func TestDog(t *testing.T) {
	// var dog Pet = new(Dog) 编译错误，不支持继承和隐式类型转换。不支持LSP
	dog := new(Dog)

	dog.SpeakTo("Chao")
}
