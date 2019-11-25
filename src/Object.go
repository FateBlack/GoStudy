package main

import "fmt"

//结构体  类似于Java对象
type user struct {
	name     string
	password string
	age      int
	//初始化零值 var age
}

func main() {
	//实例化
	u := user{"校长", "pig", 80}
	//仅填充部分属性
	//u2:=user{age:80,name:"主任"}

	/*
		函数传参
		向函数 传入结构体实例(实际为传副本) 修改后，并不影响原实例
		故，可以传指针 以此影响原实例
	*/
	//例1 传结构体
	fmt.Println(u)
	updateAge(u)
	fmt.Println(u)
	//例2 传指针
	updateAgeB(&u)
	fmt.Println(u)
}

func updateAge(us user) {
	us.age = us.age + 1
}
func updateAgeB(us *user) {
	us.age = us.age + 1
}
