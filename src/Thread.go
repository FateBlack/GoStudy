package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //计数信号量 控制等待
	wg.Add(2)             //信号量+2

	go func() {
		defer wg.Done() //执行完 信号量-1
		for i := 0; i < 20; i++ {
			fmt.Println("A", i)
		}
	}()

	go func() {
		defer wg.Done() //执行完 信号量-1
		for i := 0; i < 20; i++ {
			fmt.Println("B", i)
		}
	}()

	wg.Wait() //等待信号量 回归初始值
}
