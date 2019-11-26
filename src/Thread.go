package main

//goroutine 类似于 Java线程,以go关键字创建
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup //计数信号量 控制等待
	wg.Add(2)             //信号量+2

	go func() {
		defer wg.Done() //defer:函数结束时触发。信号量-1
		for i := 0; i < 20; i++ {
			fmt.Println("A", i)
		}
	}()

	go func() {
		defer wg.Done() //执行完后 信号量-1
		for i := 0; i < 20; i++ {
			fmt.Println("B", i)
		}
	}()

	wg.Wait() //等待信号量归零，否则阻塞
}
