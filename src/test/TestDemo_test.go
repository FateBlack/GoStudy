package main

import "testing"

//cd 指定文件目录下 go test -v 即可
//单元测试
func TestMethod(t *testing.T) {
	sum := Demo(2, 3)
	if sum == 5 {
		t.Log("成功")
	} else {
		t.Fatal("失败")
	}
}
