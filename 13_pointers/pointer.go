package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	iPtr := &i
	fmt.Println("initial:", i)
	fmt.Printf("%T  %T \n", i, iPtr) //打印各自的类型

	zeroval(i)
	fmt.Println("zeroval:", i) // every function in go is copied by value

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
