package main

import (
	"context"
	"fmt"
	"time"
)

// Context 上下文 进行并发控制
//例 控制并发停止;   ContextB中进行多个goroutine的控制
/* 原则
不要把Context放在结构体中，要以参数的方式传递
以Context作为参数的函数方法，应该把Context作为第一个参数，放在第一位。
给一个函数方法传递Context的时候，不要传递nil，如果不知道传递什么，就使用context.TODO
Context的Value相关方法应该传递必须的数据，不要什么数据都使用这个传递
Context是线程安全的，可以放心的在多个goroutine中传递
*/
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
	cancel() //发送 取消任务信号,取消当前Context 以及该节点下所有Context
	time.Sleep(10 * time.Second)
}
