package main

import (
	"fmt"
	"sync"
)

//读写锁  写写互斥，写读互斥； 允许同时读取
//读锁 RLock ; 写锁 Lock
//模拟 多个【读goroutine】和多个【写goroutine】并发地读取 修改 num

var num int
var wg sync.WaitGroup
var rwLock sync.RWMutex

//读取
func read(n int) {
	rwLock.RLock() //读锁

	fmt.Printf("读取goroutine序号 %d\n", n)
	value := num
	fmt.Printf("读取goroutine序号 %d,当前值%d\n", n, value)

	wg.Done()
	rwLock.RUnlock()
}

//写入
func write(n int) {
	rwLock.Lock() //写锁

	fmt.Printf("写goroutine序号 %d\n", n)
	value := num
	value += 1
	fmt.Printf("写goroutine序号 %d,当前值%d\n", n, value)

	wg.Done()
	rwLock.Unlock()
}

func main() {
	wg.Add(20)

	for i := 0; i < 10; i++ {
		go read(i)
	}

	for i := 0; i < 10; i++ {
		go write(i)
	}

	wg.Wait()
}
