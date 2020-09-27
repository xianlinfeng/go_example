package main

import (
	"fmt"
	"time"
)

/* 超时控制 */
func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1") // beacuse goroutine need 2s to finish, so this select will timeout
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()


	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):  //通过select机制实现超时控制。
		fmt.Println("timeout 2")
	}
}
