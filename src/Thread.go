package main

import (
	"fmt"
	//"runtime"
	"sync"
)

/**
goroutine 类似于 Java线程,以go关键字创建
Go的runtime（运行时）会在逻辑处理器上调度这些goroutine来运行，
一个逻辑处理器绑定一个操作系统线程,故goroutine非线程，它是一个协程，由Go运行时本身的算法实现

逻辑处理器	执行创建的goroutine，绑定一个线程
调度器		Go运行时中的，分配goroutine给不同的逻辑处理器
全局运行队列	所有刚创建的goroutine都会放到这里
本地运行队列	逻辑处理器的goroutine队列

创建一个goroutine的后，会先存放在全局运行队列中，
等待Go运行时的调度器进行调度，把他们分配给其中的一个逻辑处理器，
并放到这个逻辑处理器对应的本地运行队列中，最终等着被逻辑处理器执行
*/

func main() {
	//runtime.GOMAXPROCS(1) 只使用一个逻辑处理器，即单核运行

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
