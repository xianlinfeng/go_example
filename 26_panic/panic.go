/*
A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors
that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully. */
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	/*
		panic 程序异常中断，⽤用于不不可以恢复的错误
		panic 退出前会执⾏行行 defer 指定的内容 */

	/*
	   os.Exit 退出时不不会调⽤用 defer 指定的函数
	   os.Exit 退出时不不输出当前调⽤用栈信息	*/

	defer func() {
		if err := recover(); err != nil { // recover() 接收panic传入的错误消息。
			fmt.Println("recovered from ", err)
		}
	}()

	fmt.Println("start .... ")

	//panic(errors.New("Something wrong"))
	/* A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to)
	handle. Here’s an example of panicking if we get an unexpected error when creating a new file. */
	_, err := os.Create("/tmp/file")

	err = errors.New("something went wrong")

	if err != nil {
		panic(err)		//一般panic里面都是error的信息
	}
	/* Running this program will cause it to panic, print an error message and goroutine traces, and exit with a
	non-zero status.
	*/
}
