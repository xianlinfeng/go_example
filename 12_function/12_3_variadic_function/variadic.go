package main

import "fmt"

/* 可变参数 */
func sum(nums ...int) { //不用指定参数的个数，但是他们的类型都要是一致的
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums { // 这时候 nums就是一个数组，可以通过数组遍历来使用其内部的所有元素
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
