package main

import (
	"fmt"
	"sort"
)

/* In order to sort by a custom function in Go, we need a corresponding type. Here we’ve created a
byLength type that is just an alias for the builtin []string type. */
type byLength []string

/* 用我们自定义的type来实现sort.interface的接口：Len, Swap, Less三个functions。从而可以用内置的sort函数来按照我们的要求排序 */
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j]) // 由小到大
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
