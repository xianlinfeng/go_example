package main

import (
	"fmt"
	"time"
)

//This is the function we’ll run in a goroutine. The done channel will
//	be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true //Send a value to notify that we’re done.
}

func main() {
	done := make(chan bool, 1) //Start a worker goroutine, giving it the channel to notify on.
	go worker(done)

	<-done //根据channel的特点，这里会被一直阻塞，直到woker执行完成。
}
