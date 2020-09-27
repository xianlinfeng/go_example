package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs) // sort是就地进行的，不会返回新的slice
	fmt.Println("Strings:", strs)

	sort.Sort(sort.Reverse(sort.StringSlice(strs))) // reverse string逆排序
	fmt.Println("Reverse:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints))) // reverse int逆排序
	fmt.Println("Reverse:", ints)

	s := sort.IntsAreSorted(ints) //是否已排序
	fmt.Println("Sorted: ", s)
}
