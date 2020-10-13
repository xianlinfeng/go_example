package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 新建一个big Int
	var a = big.NewInt(1) //只接受int64
	var b = big.NewInt(3)
	var c = big.NewInt(4)
	var d = new(big.Int) // 创建一个*bing.Int

	a.Add(b, c)                // 加法 a = b + c
	a.Sub(b, c)                // 减法	a = b -c
	a.Mul(b, c)                // 乘法 a = b * c
	a.Div(b, c)                // 除法 a = b / c
	a.Mod(b, c)                // 求余 a = b % c
	a.Cmp(b)                   // 比较a和b，返回 -1 if a<b; 0 if a == b ; +1 if a > b
	a.Exp(b, c, nil)           // 指数 a = b^c
	a.Exp(b, c, big.NewInt(7)) // 指数 a = b^c mod 10

	fmt.Println(a.String())

	// 设置一个值
	d.Set(a)                        // 从a复制一个值
	d.SetUint64(2345)               // 根据uint64来设置
	a.SetInt64(int64(332))          // 根据int64来设置
	a.SetString("220302030034", 10) // 根据string设置为10进制

	d.Uint64() // 更改类型为uint64
	// big Int to string
	fmt.Printf("a = %v \n", a)
	fmt.Println(a.String())
}
