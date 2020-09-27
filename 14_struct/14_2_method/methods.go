package main

import "fmt"

type rect struct {
	width, height int
}

// 会修改成员的field
func (r *rect) area() int {
	return r.width * r.height
}

// 不会修改成员的field
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r rect)String()string{
	return "a string to describe the object"
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println(rp.String())
}
