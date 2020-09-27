package main

import (
	"fmt"
	"time"
)


// 创建一个job池。所有的worker都可以从工作池中拿取任务。

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second) // 睡眠1s，模拟处理job的时间
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// create channels for jobs and results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 创建3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 创建5个任务，并发送到job pool
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 从channel中拿取结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
