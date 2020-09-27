package main

import (
	"fmt"
	"testing"
	"time"
)

// 用函数作为变量的值

type IntConv func(op int) int

func TestFuncVar(t *testing.T){
	tsSF := Spent(slowFun)
	fmt.Println(tsSF(10))
}
/* 给一个函数记时 */
func Spent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n) // 通过inner func 来处理该参数
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}


