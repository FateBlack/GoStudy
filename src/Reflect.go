package main

import (
	"fmt"
	"reflect"
)

//反射
func main() {
	c := Car{name: "什么车"}
	//获取对象的 type 和 value
	value := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	fmt.Println(value, typ) //main.Car {什么车}

	/*
		逆向： value转为原始类型
		反射把对象分为【reflect.Value】和【reflect.Type】,
		但【reflect.Value】中有属性 typ *rtype; 即 value也持有type*/
	carNow := value.Interface().(Car)
	fmt.Println(carNow)
	//fmt.Println(value.Type())

	//获取type的底层类型
	fmt.Println(typ.Kind()) //例 struct

}

type Car struct {
	name string
}
