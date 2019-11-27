package main

import (
	"context"
	"fmt"
	"time"
)

//Context 控制多个goroutine
func main() {
	ctx, cancle := context.WithCancel(context.Background())

	go play(ctx, "任务A")
	go play(ctx, "任务B")
	go play(ctx, "任务C")

	time.Sleep(10 * time.Second)
	cancle() //所有基于这个Context ctx或者衍生的子Context都会收到通知，此时可结束，释放goroutine
	time.Sleep(5 * time.Second)
}

//任务
func play(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "结束")
			return
		default:
			fmt.Println(name, "正在执行")
			time.Sleep(2 * time.Second)
		}

	}
}
