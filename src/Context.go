package main

import (
	"context"
	"fmt"
	"time"
)

// Context 上下文 进行并发控制
//例 控制并发停止;   ContextB中进行多个goroutine的控制
func main() {
	/*
		context.Background() 返回一个空的Context，此空的Context一般用于整个Context树的根节点。
		然后使用context.WithCancel(parent)函数，创建一个可取消的子Context，
		然后当作参数传给goroutine使用，可使用此Context跟踪这个goroutine。*/
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {

		for {
			select {
			case <-ctx.Done(): //当接受到值时
				fmt.Println("结束")
				return
			default:
				time.Sleep(2 * time.Second)
				fmt.Println("继续任务")
			}
		}

	}(ctx)

	time.Sleep(20 * time.Second)
	cancel() //发送 取消任务信号
	time.Sleep(10 * time.Second)
}
