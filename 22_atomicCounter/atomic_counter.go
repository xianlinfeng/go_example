/*
The primary mechanism for managing state in Go is communication over channels. We saw this for example
with worker pools. There are a few other options for managing state though. Here we’ll look at using
the sync/atomic package for atomic counters accessed by multiple goroutines. */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 10000; c++ {
				atomic.AddUint64(&ops, 1) // 给counter增加1
			}
			wg.Done()
		}()
	}

	fmt.Println("ops:", atomic.LoadUint64(&ops)) // 在协程进行过程中，可以这样读取该时刻ops的值

	wg.Wait()

	/* 所有协程都完成后，现在读取是安全的 */
	fmt.Println("ops:", ops)
}
