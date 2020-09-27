package main

import (
	"errors"
	"fmt"
	"math"
)

// 定义一个接口
type geometry interface {
	area() float64
	perim() float64
}

// 两个不同的struct，并各自实现接口定义的function
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量有实现了该接口，那么我们可以call该接口内的方法。
// 这里以一个接口作为参数
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	var a geometry
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	a = circle{radius: 7} // 因为circle实现了geometry接口，那么就可以将一个circle赋值给geometry。

	// 因为r和c都实现了该接口，那么r和c就可以传入该（以接口作为参数的）function。
	measure(r)
	measure(c)
	measure(a)

	/* 断言 type assertion */
	var i interface{} = "hello"

	s := i.(string) // 返回i保存的string值，并赋值给s
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // 因为i未保存float64的值，ok = false， f=0
	fmt.Println(f, ok)

	// v := i.(type) 可以在switch语句中使用.(type)语句，然后判断返回值的类型，比如 int， string 或者其他

}

/*
	注意事项：
  1. 如果一个struct A，其指针实现了接口B，但是，其传值method未实现接口B，那么他的obj 无法使用接口B内定义的方法。
  2. 接口是非入侵性的。实现不依赖于接口的定义。
  3. 接口可以和使用者在一个包内 （和java不一样）
*/

/*
	Go接口的最佳实践
 */
//倾向于使用小的接口定义，很多接口只包含一个方法，比如：
type Reader interface {
	Read(p []byte)(n int, err error)
}

type Writer interface {
	Write(p []byte)(n int, err error)
}

//较大的接口定义，可以由多个小接口定义组合而成
type  ReaderWriter interface {
	Reader
	Writer
}

//只依赖于必要功能的最小接口
func StoreData(reader Reader)error{ //这里只依赖Reader
	// ....
	return errors.New("err")
}
