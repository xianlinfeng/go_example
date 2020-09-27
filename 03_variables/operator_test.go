package main

import (
	"testing"
)

const (
	Readable   = 1 << iota // 0001  = 1
	Writable               // 0010  = 2
	Executable             // 0100  = 4
)

/*
iota是golang语言的常量计数器,只能在常量的表达式中使用。
iota在const关键字出现时将被重置为0(const内部的第一行之前)，const中每新增一行常量声明将使iota计数一次
(iota可理解为const语句块中的行索引)。
*/

func TestBitClear(t *testing.T) {

	/*  &^ 按位置清零操作符
	1 &^ 0 = 1		右边为0的时候，保持左边的不变
	0 &^ 0 = 0
	1 &^ 1 = 0		右边为1的时候，结果为0
	0 &^ 1 = 0
	*/
	a := 7              //0111
	a = a &^ Readable   // 0111 &^ 0001 = 0110
	a = a &^ Executable // 0110 &^ 0100 = 0010
	t.Log(Readable)
	t.Log(Writable)
	t.Log(Executable)
	t.Log(a & Readable) // 0010 & 0001 = 0000
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

func TestOper(t *testing.T) {
	t.Log(0 ^ 1)
}
