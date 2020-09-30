package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 声明一个map
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}

	m2 := map[string]int{}
	m2["one"] = 1

	m3 := make(map[string]int, 10) // 因为make出来的map和slice一样，是自增长的，如果我们刚开始就知道map的长度，则可以make的时候就指定map长度，以节省内存空间
	t.Log("the length of m3 is ", len(m3))

	// 删除元素
	delete(m1, "one")

	// 判断改值是否在map中存在
	// 在访问的 Key 不不存在时，仍会返回零值，不能通过返回 nil 来判断元素是否存在
	if v, ok := m3["two"]; ok {
		t.Logf("m1[\"two\"] = %d \n", v)
	} else {
		t.Log("m1[\"two\"] is not exist!")
	}

	// 对map进行遍历
	for i, v := range m1 {
		t.Log(i, v)
	}
}

func TestFactory(t *testing.T) {
	/* 工厂模式 */
	// map 的值可以是一个func
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

func TestMapasSet(t *testing.T) {
	/* 用map来实现set */
	mySet := map[int]bool{} // 声明
	mySet[1] = true         //	添加元素
	if i := 3; mySet[i] {   // 检查元素是否存在
		t.Logf("%d is existing\n", i)
	} else {
		t.Logf("%d is not existing\n", i)
	}
	t.Log(len(mySet)) // 查看set的长度
	delete(mySet, 1)  // 删除元素
}
