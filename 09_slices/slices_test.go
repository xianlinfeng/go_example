package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	// 切片不可以相互比较（只可以和nil比较）

	var a []int // 切片的声明
	fmt.Println(a, len(a), cap(a))

	b := []int{1, 2, 3, 4, 5} // 创建一个slice
	fmt.Println(b, len(b), cap(b))

	s := make([]string, 3) // 创建一个slice（长度为3）
	fmt.Println("emp:", s)
	s = append(s, "new") //添加元素

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f") // 一直添加的话，会改变slice的容量
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) // 切片的复制
	fmt.Println("cpy:", c)
}

func TestMultiDimensionSlice(t *testing.T) {
	twoD := make([][]int, 3) // 二维slices
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
	fmt.Println("2d: ")

	//打印一个二维三角形
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		fmt.Print("    ")
		for j := 0; j < innerLen; j++ {
			fmt.Printf("%2d", twoD[i][j])
		}
		fmt.Println()
	}
}
