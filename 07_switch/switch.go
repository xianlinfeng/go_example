package main

import (
	"fmt"
	"time"
)

func main() {
	// go 语言中不需要break来明确退出一个case

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i { // 正常switch语句
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // 一个case包含多种可能，用逗号分隔
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { // 没有表达式的switch语句，可以看作一组if...else...语句
	case t.Hour() < 12: // case语句为一个表达式, case为true的时候，执行。

		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { //空接口可以被任何类型实现
		switch t := i.(type) { // 断言
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
