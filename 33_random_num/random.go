package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(rand.Intn(100), ",") // rand.Intn returns a random int n, 0 <= n < 100.
	fmt.Print(rand.Intn(100))
	fmt.Println()

	fmt.Println(rand.Float64()) //rand.Float64 returns a float64 f, 0.0 <= f < 1.0.

	// This can be used to generate random floats in other ranges, for example 5.0 <= f' < 10.0.
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	/* 默认的数字生成器是确定性的，因此默认情况下每次都会产生相同的数字序列。 要产生变化的序列，请给它一个变化的种子。
	请注意，对于打算保密的随机数使用此方法并不安全，请对它们使用crypto/rand。 */
	s1 := rand.NewSource(time.Now().UnixNano()) // 使用时间做seed
	r1 := rand.New(s1)                          // 返回*rand.Rand对象

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	/* 如果你的种子是同一个，那么每次将会产生相同的random num */
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
