package main

import "fmt"

<<<<<<< HEAD
var s string = "Hello world!" // function外不能使用短格式
=======
var test string = "hello " //function 外，不支持短格式赋值
>>>>>>> b98a5f0... modify on mac

func main() {

	// 基础类型： bool, int, float, byte, rune, complex
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var e int
	fmt.Println(e) // 默认零值为0

	var ( // 批量声明变量
		g int32   = 3
		h float32 = 0.343
		i int64   = 32323
	)
	fmt.Println(g, h, i)

	f := "apple"
	fmt.Println(f)
}
