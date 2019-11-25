package main

import "fmt"

//方法
//与Java有差别，方法!=函数
//方法属于结构体或类型   函数不属于

type tea struct {
	name string
	age  int
}

//可称为 tea 有方法 addAge
func (t *tea) addAge() {
	t.age = t.age + 1
}

//方法 返回值: 例 string 为返回值类型
func (t *tea) getNewName() string {
	return t.name + "NEW!"
}

//此方法 为修改副本，不影响原实例
//func (t tea) addAgeB(){
//	t.age = t.age+1
//}

func main() {
	t := tea{"薄荷茶", 5}

	t.addAge() //等同  (&t).addAge() ; 原方式 go自动取指针
	fmt.Println(t)

	newName := t.getNewName()
	fmt.Println(newName)
}
