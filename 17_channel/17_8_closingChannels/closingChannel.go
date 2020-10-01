package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // 一直从channel中获取数据
			if more {
				fmt.Println("goroutine 1: received job", j)
			} else {
				fmt.Println("received all jobs for goroutine 1")
				done <- true // 如果jobs里没有更多的任务，则往done里传入true
				return
			}
		}
	}()

	go func() {
		for i := range jobs {
			fmt.Println("goroutine 2: received job", i)
		}
		done <- true
		fmt.Println("received all jobs for goroutine 2")
	}()

	for j := 1; j <= 6; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // 如果不close channel， receiver会一直等待。
	// close channel会发送一个0的信号。
	fmt.Println("sent all jobs")

	<-done // 等待 true的信号（所有任务完成）
}
