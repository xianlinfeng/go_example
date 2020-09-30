package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var a [5]int // 数组的声明
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5} // 声明的时候同时初始化
	fmt.Println("dcl:", b)

	c := [...]int{1, 2, 3, 4, 5, 6} //不预先指定array的长度，而是根据实际使用情况来定
	fmt.Println("c: ", c)

	//数组的遍历1：
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]= %d \n", i, c[i])
	}

	//数组的遍历2：
	for index, elem := range c {
		fmt.Printf("c[%d]= %d \n", index, elem)
	}
}

// 数组的比较
func TestArrayCompare(t *testing.T) {
	//	(在其他语言中不可以比较，因为是地址引用）
	// 		相同维数且含有相同个数元素的数组才可以⽐比较
	// 		每个元素都相同的才相等
	e := [...]int{1, 2, 3, 4}
	f := [...]int{1, 3, 4, 5}
	g := [...]int{1, 2, 3, 4, 5}
	h := [...]int{1, 2, 3, 4, 5}
	fmt.Println(e == f)
	//fmt.Println(e == g)  // 不可以比较，因为长度不一样
	fmt.Println(g == h)
}

// 多维数组
func TestMultiDimensionArray(t *testing.T) {
	var d [2][3]int
	d = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(d)

	//二维array和二维slices不一样
	var twoD [2][3]int //二维矩阵
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
