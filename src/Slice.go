package main

import (
	"fmt"
	"sort"
)

//切片 ，类似 list
func main() {

	/**
	创建切片	方式1
	参数1:类型 ， 参数2:长度 ，参数3:容量
	此例目前仅可访问前3个，扩容后方可访问剩下2个
	*/
	//sliceA:=make([]int,3,5)

	//创建切片	方式2
	//sliceA:=[]int{}  //空切片
	//var sliceA []int //nil切片

	sliceA := []int{3, 4, 5, 6}

	//切片截取 , 截取范围 在 0<= x < 2
	sliceB := sliceA[0:2]
	fmt.Println(sliceB)
	//截取全部
	sliceC := sliceA[:]
	fmt.Println(sliceC)
	/*
		注:切片截取后，实质上仍共用一个底层数组
		若想复制出全新切片，应使新切片 长度=容量
		这样append元素后，可扩容生成新底层数组
	*/

	//切片 添加元素 , 可一次添加多个
	sliceA = append(sliceA, 8, 9, 10)
	fmt.Println(sliceA)

	//切片 添加另一切片,实质上,是打散切片进行添加
	sliceD := []int{22, 33, 444}
	sliceD = append(sliceD, sliceA...)
	fmt.Println(sliceD)

	//迭代 方式1
	for i, v := range sliceA {
		fmt.Printf("索引 %d ,值 %d\n", i, v)
	}

	//迭代 方式2
	for i := 0; i < len(sliceD); i++ {
		fmt.Println(sliceD[i])
	}
	//排序
	sort.Ints(sliceD)
	fmt.Println(sliceD)
}
