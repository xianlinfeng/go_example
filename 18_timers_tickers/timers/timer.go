package main

import (
	"fmt"
	"time"
)

func main() {
	/* 计时器： 在未来一段时间后，给通道发送一个通知 */
	timer1 := time.NewTimer(1 * time.Second)

	<-timer1.C // 当计时结束后，给timer.c通道发信息
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second * 2)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop() //手动停止计时器
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

}
