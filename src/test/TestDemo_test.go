package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

//注：测试文件后缀必须 _test
//cd 指定文件目录下 go test -v 即可
//单元测试
func TestAdd(t *testing.T) {
	sum := Add(2, 3)
	if sum == 5 {
		t.Log("成功")
	} else {
		t.Fatal("失败")
	}
}

/*
	测试 API
*/
func init() {
	Routes() //初始化调用 待测试函数的 路由配置函数
}
func TestSendJson(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/sendjson", nil)
	if err != nil {
		t.Fatal("创建request失败")
	}

	rw := httptest.NewRecorder()            //模拟服务端响应
	http.DefaultServeMux.ServeHTTP(rw, req) //触发响应

	t.Log("code:", rw.Code)
	t.Log("body:", rw.Body.String())
}
