package main

import (
	"fmt"
	"time"
)

/*  GO的function：
	1. 可以有多个返回值
	2. 所有的参数都是值传递，map，slice，channel可能会有传引用的错觉。
		比如slice： 一个slice其实是一个数据结构，其中存储着指向一个array的指针。当传入一个函数的时候
   		，其内部的指针被复制。当操作这些复制后的值时，其实其指向的是和原slice指向的同一片内存空间。所以
		才会有传引用的错觉
	3. 函数可以作为变量的值
	4. 函数可以作为参数和返回值
*/

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	defer func() { //延时执行函数，即使panic后，defer仍然会执行
		// 在os.exit()后，defer不会执行
		fmt.Println("clear the resources...")
		fmt.Println("this is the defer function")
	}()

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	tsSF := timeSpent(slowFun)
	fmt.Println(tsSF(10))
}

/* 给一个函数记时 */
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n) // 通过inner func 来处理该参数
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
