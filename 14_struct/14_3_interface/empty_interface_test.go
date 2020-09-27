package main

import (
	"fmt"
	"testing"
)

func DoSomething(i interface{}){ 	// 空接口可以接受任何类型，但是必须用断言来判断数据的类型
	switch t := i.(type) { // 断言
	case bool:
		fmt.Println("I'm a bool")
	case int:
		fmt.Println("I'm an int")
	case string:
		fmt.Println("I'm a string")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}

}

func TestEmptyInterface(t *testing.T){
	DoSomething("string")
	DoSomething(10)
	DoSomething(3.14)


}