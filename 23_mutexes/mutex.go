/*
In the previous example we saw how to manage simple counter state using atomic operations.
For more complex state we can use a mutex to safely access data across multiple goroutines. */
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

/*锁和原子操作*/

func main() {

	var state = make(map[int]int)

	//var mutex sync.Mutex //此互斥锁将同步对状态的访问。
	var lock sync.RWMutex	// 也可以用读写锁

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ { // 100个goroutine，持续读取，间隔1ms
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				//mutex.Lock()
				lock.RLock()
				total += state[key]
				//mutex.Unlock()
				lock.RUnlock()
				atomic.AddUint64(&readOps, 1) // 原子写入

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ { // 10个goroutine，持续写入，间隔1ms
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				//mutex.Lock()
				lock.Lock()
				state[key] = val
				//mutex.Unlock()
				lock.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second * 2) // 整个程序运行2秒钟

	readOpsFinal := atomic.LoadUint64(&readOps) // 原子读
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	//mutex.Lock()
	//fmt.Println("state:", state)
	//mutex.Unlock()

	lock.RLock()
	fmt.Println("state:", state)
	lock.RUnlock()
}
