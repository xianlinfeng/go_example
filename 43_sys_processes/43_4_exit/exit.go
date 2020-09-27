package main

import (
	"fmt"
	"os"
)

func main() {

	//如果使用os.exit，defer函数不会执行。
	defer fmt.Println("!")

	fmt.Println("OS Exit...")
	os.Exit(1) // 可以使用其他值来标识异常退出
}
