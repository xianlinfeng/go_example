# Go 语言格式化输出

## 一， 通用

```go
%v  // 值的默认格式表示 Value
%+v // 类似%v，但是输出结构体的时候，会添加字段名
%#v // 值的Go语法表示
%T  // 值的类型的Go语法表示 Type
fmt.Printf("the key of map[%[1]d] is %[1]d ", k)    // 其中[1]表示都使用第一个参数
```

## 二， 布尔值

```go
%t      // 单词true或false true
```

## 三， 整数

```go
%b  // 表示为二进制 binary
%c  // 该值对应的unicode码值 char
%d  // 表示为十进制 digital
%8d // 表示该整型长度是8，不足8位则在数值前补空格，超过8则以实际为准
%08d// 表示该整型长度是8，不足8位则在数值前补0，超过8则以实际为准
%o  // 表示为八进制， octal
%q  // 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x  // 表示为十六进制，使用a-f  hex
%X  // 表示为十六进制，使用A-F  hex
%U  // 表示为Unicode格式：U+1234，等价于“U+%04x” unicode

```

## 四， 浮点数与复数

```go
%b  // 无消暑部分，二进制数的科学计数法。如-12343p-78；
%e  // （=%.6e)有6位消暑部分的科学计数法，如 -1234.456e+78
%E  // 科学计数法，如-1234.456E+78
%f  // （=%.6f) 有6位小数部分， 如1.234567 float
%6.3f // 宽度为6，精度为3
%F  // 等价于 %f
%g  // 根据实际情况采用%e或者%f格式 （以获得更简洁，准确的输出）
%G  // 根据实际情况采用%e或者%f格式 （以获得更简洁，准确的输出）
```

## 五， 字符串和[]byte

```go
%s  // 直接输出字符串或者[]byte     string
%q  // 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x  // 每个字节使用两个字符的十六进制数表示（使用a-f）
%X  // 每个字节使用两个字符的十六进制数表示（使用A-F）
```

## 六， 指针

```go
%p  // 表示为十六进制数，并加上前导的0x     pointer
无%u // 整数如果数无符号类型，输出自然也是无符号的。类似的，也没有必要指定操作数的尺寸（int8，int64）
```

## 七，其他flag

- `+` 总是输出数值的正负号；%q（%+q）会生成全部都是ASCII字符的输出（通过转义)  
- ``  对数值，正数前加空格，而对数前加负号
- `-` 在输出的右边填充，而不是左边。
- `#` 切换格式:  
- 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前门的0x（%#p）  
- 对%q（%#q），如果strconv.CanBackquote返回真会输出反括号括起来的未转义的字符串  
- 对%U（%#U），输出Unicode格式后可，如字符可打印，还会输出空格和单括号括起来的go字面值。  
- 对字符串采用%x或%X时，会给各打印的字节之间加空格
- `0` 使用0而不是空格填充，对于数值类型，会把填充的0放在正负号后面。