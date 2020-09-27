/*Rate limiting is an important mechanism for controlling resource utilization and maintaining
quality of service. Go elegantly supports rate limiting with goroutines, channels, and tickers. */
package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		均
		匀
		速
		率
	*/
	// 创建一个channel，并往里面发送5个任务。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 创建一个ticker （200ms）
	limiter := time.Tick(2 * time.Second)

	// 为了保证服务质量，每200ms处理一个请求。
	for req := range requests {
		<-limiter // 每隔200ms，limiter会收到一个信号，并且unblock。
		fmt.Println("request", req, time.Now())
	}

	/*
		突
		发
		请
		求
	*/
	// 我们可能会想要处理突发事件，但又保留原有但速率。下面这个burstyLimiter channel允许3个突发事件
	burstyLimiter := make(chan time.Time, 3)

	// 先放入3个time
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每隔200ms，再放入一个event
	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstyLimiter <- t
		}
	}()

	// 模拟5个request，前面3个，将会被当作突发事件处理。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
