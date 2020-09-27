/*我们经常需要我们的程序对数据集合执行操作，例如选择满足给定谓词的所有项或使用自定义函数将所有项映射到新集合。

在某些语言中，惯用的是通用的数据结构和算法。 Go不支持泛型； 在Go中，如果您的程序和数据类型特别需要收集功能，通常会提供这些功能。

这是一些用于字符串切片的示例收集函数。 您可以使用这些示例来构建自己的功能。 请注意，在某些情况下，最简单的方法是直接直接内联集合
操作代码，而不是创建和调用辅助函数。 */
package main

import (
	"fmt"
	"strings"
)

//Index returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//Include returns true if the target string t is in the slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//Any returns true if one of the strings in the slice satisfies the predicate f.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the strings in the slice satisfy the predicate f.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	// Output: 2 // as "pear" appears in the strs of index of 2

	fmt.Println(Include(strs, "grape"))
	// Output: false // as not contain "grape" in strs

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	// func(v string) bool { return strings.HasPrefix(v, "p") } 作为一个参数
	// Output: true

	fmt.Println(All(strs,
		func(v string) bool { // 将一个function作为参数。
			return strings.HasPrefix(v, "p") // 如果这个function返回true，则f(v)返回true
		}))
	// Output: false

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	// Output: [peach apple pear]

	fmt.Println(Map(strs, strings.ToUpper))
	// Output: [PEACH APPLE PEAR PLUM]
}
