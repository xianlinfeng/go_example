package main

import "fmt"

/* 把对象的数据封装在struct中 */
type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// 创建obj 方法1:
	fmt.Println(person{"Bob", 20})

	// 创建obj 方法2:
	fmt.Println(person{name: "Alice", age: 30})

	// 创建obj 方法3:
	p := new(person) // 返回的是一个指针

	fmt.Println(person{name: "Fred"}) // 只初始化部分field

	fmt.Println(&person{name: "Ann", age: 40}) // return the pointer

	fmt.Println(newPerson("Jon")) //return the pointer too

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51 // change sp will effect s too
	fmt.Println(sp.age)
	fmt.Println(s.age)

	p.name = "Arthur"
	p.age = 32
	fmt.Println(p)
}
