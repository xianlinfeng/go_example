package main

import (
	"strconv"
	"strings"
	"testing"
)

//string是数据类型，不是引用或者指针类型
//string是制度的byte slice，len(string)返回的是它包含的byte数
//string的byte数组可以存放任何数据

func TestString(t *testing.T) {
	var s string // 声明
	t.Log(s)     //初始化为默认零值“”

	s = "hello"
	t.Log(len(s)) // = 5 bytes
	//s[1] = '3' //string是不可变的byte slice

	s = "\xE4\xB8\xA5" //可以存储任何二进制数据
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s)) // = 3 bytes
}

func TestString1(t *testing.T) {
	// Unicode: 一种字符集（code point），e.g. "中" 在unicode中是"0x4E2D"
	// UTF8： Unicode的存储实现， e.g. "中" = "0xE4B8AD"
	// string/[]byte中存储为： 		"中" = [0xE4,0xB8,0xAD]
	// 当“中”存储在string中的时候，会转换为UTF8编码，在UTF8编码中，“中”占3bytes

	s := "中"
	t.Log(len(s)) // = 3 bytes (是byte数)

	c := []rune(s)     // rune = int32,表达为一个unicode code point
	t.Logf("%X \n", c) //	[4E2D]
	t.Log(len(c))      // = 1 rune
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %X \n", c[0]) // = 0x4E2D  可以看到unicode占2bytes
	t.Logf("中 UTF8 %X \n", s)       // = 0xE4B8AD  而UTF8占3bytes
}

func TestStringToRune(t *testing.T) {
	// string的遍历
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c) // 使用range迭代输出的，是rune（unicode），系统会自动转换
	}
	t.Log(len(s)) // = 21, 因为有7个字符，每个汉字占3个byte，一共21个（string是byte的数组）
}

func TestStringFn(t *testing.T) {
	// 字符串的拆分和连接
	s := "A,B,C"

	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	// string和int的转换
	s := strconv.Itoa(10)
	t.Log("str" + s)

	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
