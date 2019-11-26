package main

import "log"

func main() {
	//日志
	log.Println("日志测试")
}

//【init】函数在【main】函数执行之前就可以初始化，可进行配置
func init() {
	//自定义日志
	log.SetPrefix("【管理员模块】")                 //添加 日志前缀
	log.SetFlags(log.Ldate | log.Lshortfile) //修改时间格式，显示文件名和行号
	//	Ldate         = 1 << iota     //日期示例： 2009/01/23
	//	Ltime                         //时间示例: 01:23:23
	//	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
	//	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
	//	Lshortfile                    //文件和行号: d.go:23.
	//	LUTC                          //日期时间转为0时区的
	//	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
}
