package main

import (
	"fmt"
	"sync"
)

//互斥锁 Mutex

var (
	num   int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2)
	go addNum("A")
	go addNum("B")
	wg.Wait()
}

func addNum(name string) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		mutex.Lock()
		num = num + 1
		fmt.Println(name, num)
		mutex.Unlock()
	}
}
