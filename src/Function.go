package main

import "fmt"

//函数惊喜用法 不同于Java部分

//多值返回; 括号二：返回值类型
func more(x int) (string, int) {
	if x == 0 {
		x -= 1
		return "err", x
	}
	return "success", x
}

//可变参数 任意类型和数量，
// 空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。
func allChange(people ...interface{}) {
	fmt.Println("参数 类型 和 数量 均可变化")
	for _, v := range people {
		fmt.Println(v)
	}

}

//可变参数 ,任意数量参数；例 传入任意 int 数字
func change(people ...int) {
	sum := 0
	for _, v := range people {
		sum += v
	}
	fmt.Println("和：", sum)
}

func main() {
	msg, code := more(0)
	fmt.Println(msg, code)

	change(1, 2, 3)

	allChange("a", "b", "c")
	allChange(7, 8, 9)
}
