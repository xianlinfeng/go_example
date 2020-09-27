package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
		sync.Pool 对象获取:
	1. 尝试从私有对象获取(一个processor只有一个对象，私有对象是协程安全的-写入的时候，是不需要锁的)
	2. 私有对象不不存在，尝试从当前 Processor 的共享池获取（共享池是协程不安全的-需要锁）
	3. 如果当前 Processor 共享池也是空的，那么就尝试去其他 Processor 的共享池获取
	4. 如果所有⼦子池都是空的，最后就⽤用户指定的 New 函数 产⽣生⼀一个新的对象返回
*/

/*
		sync.Pool 对象的放回
	1. 如果私有对象不不存在则保存为私有对象
	2. 如果私有对象存在，放⼊入当前 Processor ⼦子池的共享池中
*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} { // New 函数：在所有方法都没找到对象的时候，通过这个方法创建一个对象
			fmt.Println("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int) //对获取的对象做一个断言，判断是否为int
	// 第一次get，因为所有池都是空的，所以会自动创建一个对象。
	fmt.Println(v)

	pool.Put(3)
	v = pool.Get().(int) // 因为已经有对象了，所以不用创建心的对象
	fmt.Println(v)

	pool.Put(22)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	runtime.GC() // (好像在新版本中，两次gc才会晴空对象池)
	v1 := pool.Get().(int)
	fmt.Println("After gc", v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 7
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		/* 会先取出3个100，然后按照默认的New方法，创建出7个7，然后取出 */

		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
