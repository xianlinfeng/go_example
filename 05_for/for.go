package main

import "fmt"

func main() {
	for j := 7; j <= 9; j++ { // 正常for循环
		fmt.Println(j)
	}

	i := 1
	for i <= 3 { // while循环
		fmt.Println(i)
		i += 1
	}

	for { //无限循环
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // continue， break 关键字
		}
		fmt.Println(n)
	}

	//for i := range set
}
