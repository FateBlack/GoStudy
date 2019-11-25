package main

import "fmt"

func main() {
	//创建 Map
	//例如 key类型int , value类型String
	m := make(map[int]string)
	//m:=map[string]int{}
	//m:=map[string]int{2333:"宝可梦"}

	//添加或修改元素
	m[3] = "奇拉美"

	//获取值
	name := m[3]
	fmt.Println(name)

	//判断 key 是否存在. exists为boolean类型，可使nameA为_
	nameA, exists := m[777]
	fmt.Println(nameA)
	fmt.Println(exists)

	//删除元素
	delete(m, 3)
	fmt.Println(m)

	//遍历
	m[3] = "杰尼龟"
	m[4] = "卡比"
	for key, value := range m {
		fmt.Println(key, value)
	}

}
