package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//写
	var buffer bytes.Buffer
	buffer.Write([]byte("测试")) //字符串转为一个字节切片,并写入数据流
	buffer.WriteTo(os.Stdout)  //buffer内容打印到控制台

	//读
	var p [300]byte
	n, err := buffer.Read(p[:])
	fmt.Println(n, err, string(p[:n])) //n 为读取字节数, err返回信息

	//读 方式2
	data, errB := ioutil.ReadAll(&buffer)
	fmt.Println(string(data), errB)
}
