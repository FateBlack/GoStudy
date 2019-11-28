package main

import (
	"encoding/json"
	"net/http"
)

//待测试函数
func Add(a, b int) int {
	return a + b
}

/*
	处理Http请求函数
*/
func Routes() {
	http.HandleFunc("/sendjson", SendJson) //设置路由
}

func SendJson(rw http.ResponseWriter, req *http.Request) {
	u := struct {
		Name string
	}{
		Name: "老王",
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK) //code码
	json.NewEncoder(rw).Encode(u)
}
