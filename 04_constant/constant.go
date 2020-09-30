package main

import (
	"fmt"
	"math"
)

const s string = "constant"

// 星期
// The predeclared iota identifier resets to 0 whenever the word const appears in the source code
//   and increments after each const specification:
const (
	Monday = 1 + iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

// 文件读写
const (
	Readable = 1 << iota
	Writable
	Executable
)

func main() {
	fmt.Println(s)
	fmt.Println(Thursday)   // 	4
	fmt.Println(Executable) // 100(二进制）= 4
	fmt.Println(6&Readable == Readable)

	const n = 500000000 // 可以声明一个常量，但是不指定其类型。再使用的时候，再转化为相应的类型
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
