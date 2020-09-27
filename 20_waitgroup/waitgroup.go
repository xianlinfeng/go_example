package main

import (
	"fmt"
	"sync"
	"time"
)

//This is the function we’ll run in every goroutine.
//   Note that a WaitGroup must be passed to functions by pointer.
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	wg.Done() //完成一个
}

func main() {

	var wg sync.WaitGroup // 创建一个等待组

	for i := 1; i <= 5; i++ {
		wg.Add(1) // 为每个goroutine 增加计数
		go worker(i, &wg)
	}

	wg.Wait() // 阻塞，直到waitgroup的计数器归零
}
