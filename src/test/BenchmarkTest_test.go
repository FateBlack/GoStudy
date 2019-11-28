package main

import "testing"

/*
	基准测试：测试CPU和内存的效率问题，评估被测试代码的性能
*/
/*
基准测试的代码文件必须以_test.go结尾
基准测试的函数必须以Benchmark开头，必须是可导出的
基准测试函数必须接受一个指向Benchmark类型的指针作为唯一参数
基准测试函数不能有返回值
b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
最后的for循环很重要，被测试的代码要放到循环里
b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
*/
/*
go test -bench=. -run=none
-bench=标记，它接受一个表达式作为参数，匹配基准测试的函数，.表示运行所有基准测试
-run=匹配一个从来没有的单元测试方法，过滤掉单元测试的输出，使用none
-benchtime指定运行时间，例如 加入  -benchtime=5s 运行5秒
-benchmem 查看内存分配次数
*/
func BenchmarkAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Add(10, 20)
	}
}

/*
运行结果： -4代表GOMAXPROCS值(使用CPU数); 1000000000代表for循环次数; 0.392 ns/op 表示耗费纳秒
goos: windows
goarch: amd64
BenchmarkAdd-4          1000000000               0.392 ns/op
PASS
ok      _/D_/GoWorkSpace/GoStudy/src/test       2.906s
*/
