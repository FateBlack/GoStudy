package main

import "fmt"

//通道 Chan
//在两个routine之间架设的管道，
// 一个goroutine可以往这个管道里塞数据，
// 另外一个可以从这个管道里取数据，类似队列

//注意 【无缓冲管道】 线程写入数据后，必须要【其他线程】读取，否则阻塞！
//     【缓冲通达】塞满缓冲通道，无【其他线程】读取才阻塞
func main() {
	//创建 【无缓冲通道】，指定通道传输的数据类型，例如int
	ch := make(chan int)

	go func() {
		ch <- 5 //将一个数据写入channel，会导致阻塞！，直到有其他goroutine从这个channel中读取数据
	}()

	//num:=<-ch//从channel中读取数据，如果channel之前没有写入数据，也会导致阻塞！，直到channel中被写入数据为止
	fmt.Println(<-ch) //从通道里读取值，然后忽略
	//close(ch)	// 关闭channel

	//【可缓冲通道】 例：可写入三个
	chB := make(chan int, 3)
	fmt.Println(len(chB), cap(chB)) //当前长度，容量

}

//单向管道
//var send chan<- int //只允许 发送
//var receive <-chan int //只允许 接收
// 例 : 限制了传入管道的功能，只允许其发送数据
func test(send chan<- int) {
}
