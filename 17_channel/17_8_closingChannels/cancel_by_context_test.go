package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelledbyContext(ctx context.Context) bool {
	select {
	case <-ctx.Done(): //接收消息通知
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	// context.Background() 创建根节点
	// context.WithCancel(） 创建子节点

	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelledbyContext(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel() // 取消context任务
	time.Sleep(time.Second * 1)
}
