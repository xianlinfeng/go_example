package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // 返回它的index和value
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { // 返回 key 和 values
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // return keys
		fmt.Println("key:", k)
	}

	for i, c := range "golang" { // 返回 index 和 rune（unicode code point）
		fmt.Println(i, c)
	}
}
