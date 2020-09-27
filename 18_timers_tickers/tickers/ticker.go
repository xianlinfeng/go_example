/*
Timers are for when you want to do something once in the future - tickers are for when you want to do
something repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it. */
package main

import (
	"fmt"
	"time"
)

func main() {
	// 节拍器： 当你想要重复一个事件的时候。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: // 发送时间标签。
			// 当ticker停止的时候，它不会再从该通道收到任何信息。
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(3600 * time.Millisecond)
	ticker.Stop() 	// 停止节拍器
	done <- true
	fmt.Println("Ticker stopped")
}
