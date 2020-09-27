package main

import (
	"fmt"
	"math"
)

func main() {

	// 基础类型： bool, int, float, byte, rune, complex
	// 基础类型的零值： string的零值为nil（即""）
	// go 不支持隐式类型转换（必须使用显示的类型转换）：
	var a int = 1
	var b int64
	b = int64(a)
	fmt.Println(b)

	type float float32 //包括自定义的类型
	var f float = 1.1
	var g float64
	g = float64(f)
	fmt.Println(g)

	//类型的预定义值
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MinInt64)

}
