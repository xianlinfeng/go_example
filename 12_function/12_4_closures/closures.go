package main

import "fmt"

// 闭包：一个func1的返回值是另一个func2，每次call这个func1的时候，都会访问（或更改）其内部的参数，该参数也在func2中使用。

// 这里是将function作为返回值，而28_collection_func 是将function作为参数传入一个function

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func adder() func(int) int { // 连续相加
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int { //返回斐波那契数列
	a := 0
	b := 1
	return func() int { // 该函数可以访问外部的a，b值，随着每次call该函数，更新外部的a，b值。
		c := a
		a = b
		b = b + c
		return c
	}
}

func main() {

	nextInt := intSeq()   // 每次call intSeq的时候，就会重置intSeq内部的i的值
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println()
	pos := adder() // 连续相加
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}

	fmt.Println()
	f := fibonacci() // 用闭包打印斐波那契数列
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
