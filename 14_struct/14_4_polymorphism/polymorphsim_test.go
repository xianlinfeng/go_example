package polymorphism

/* 多态：一个父对象和一堆子对象，对父对象和子对象创建相同的method。
当给父对象赋值（为不同的子对象）时，并且call 这个method，父对象会根据子对象不同的类型，来call 相应子对象的method */

import (
	"fmt"
	"testing"
)

// 基于interface来实现多态
type Code string            // 自定义一个类型
type Programmer interface { // 定义接口
	WriteHelloWorld() Code // 接口内的方法
}

// struct(1) 和它接口的实现
type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World!\")"
}

// struct(2) 和它接口的实现
type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World!\")"
}

func writeFirstProgram(p Programmer) { // func的参数为接口
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) // %T 输出传入参数的类型
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{} // Programmer接口类型对应的必须为指针类型的实例
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg) // 传入不同的struct，会引用不同的接口实现。
	writeFirstProgram(javaProg)
}
