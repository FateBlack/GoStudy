package main

import (
	"fmt"
)

//数组
func main() {
	//创建空数组 var arr [3]int

	//创建数组
	arr := [3]int{3, 4, 5} //或 arr:=[...]int{3, 4, 5} 自动判断长度

	//迭代方式
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}

	//迭代方式 2 ， i为索引，v为值;
	//可将 i或v 替换为 _ 表示忽略不使用
	for i, v := range arr {
		fmt.Printf("索引 %d ,值 %d\n", i, v)
	}
}
