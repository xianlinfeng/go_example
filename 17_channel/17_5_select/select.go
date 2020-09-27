package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	/* Each channel will receive a value after some amount of time, to simulate
	e.g. blocking RPC operations executing in concurrent goroutines. */
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	/* We’ll use select to await both of these values simultaneously,
	printing each one as it arrives. */
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: // 如果没有匹配的，就会被block，等待。
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*
nuc:18_select xianlinfeng$ time go run select.go
received one
received two

real    0m2.178s	//Note that the total execution time is only ~2 seconds
					//	since both the 1 and 2 second Sleeps execute concurrently
user    0m0.194s
sys     0m0.105s
*/
