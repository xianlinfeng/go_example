package main

import fm "fmt" // 可以用fm来做fmt的别名

/* 包管理
1. go get -u 强制从⽹网络更更新远程依赖
*/

/* init 函数*/
// 在main被执行前，所有依赖包的init方法都会被执行
// 不同包的init函数按照包倒入的依赖关系决定执行顺序
// 每个包可以有多个init函数
// 每个源文件也可以有多个init函数

func init() {
	fm.Println("init1")
}

func init() {
	fm.Println("init2")
}

func main() {
	fm.Println("Hello, world!")
}
