package main

import "fmt"

//接口

//定义 接口 以及方法
type person interface {
	play()
}

//实现 结构体
type man struct {
}
type woman struct {
}

//实现方法
func (m man) play() {
	fmt.Println("man pla")
}

func (w woman) play() {
	fmt.Println("woman play")
}

func main() {
	var p person
	p = new(man)
	p.play()

	p = new(woman)
	p.play()
}
