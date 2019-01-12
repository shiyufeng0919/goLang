package 编译与工具

import (
	"testing"
	"fmt"
)

/*
1.使用基准测试框架测试加法性能

注：测试代码不使用全局变量等带有记忆性质的数据结构
*/
func Benchmark_Add(b *testing.B){
	var n int
	for i:=0;i<b.N;i++{ //b.N是基准测试框架提供
		n++
	}
}
/*
测试：go test -v -bench=. jizhun_test.go  ##-bench=.表运行jizhun_test.go文件里的所有基准测试

ZBMAC-C02VX5Y7H:编译与工具 shiyufeng$ go test -v -bench=. jizhun_test.go

Benchmark_Add-8         2000000000               0.31 ns/op
PASS
ok      command-line-arguments  0.667s

Benchmark_Add-8 #基准测试名称
2000000000 #测试的次数，即testing.B结构中提供给程序使用的N
0.31 ns/op #表每一个操作耗费多少时间(纳秒)

注：windows下使用go test命令行时，-bench=. 需要写成 -bench="."

*/


/*
2.基准测试原理

基准测试框架对一个测试用例的默认测试时间是1s
*/

/*
3.自定义测试时间 -benchtime
go test -v -bench=. -benchtime=5s jizhun_test.go
*/

/*
4.测试内存

基准测试可以对一段代码可能存在的内存分配进行统计。
*/
func Benchmark_Alloc(b *testing.B){
	for i:=0;i<b.N;i++{
		fmt.Println("%d",i)
	}
}
/*
测试:

go test -v -bench=Alloc -benchmem jizhun_test.go  ##-bench=Alloc表示只测试Benchmark_Alloc函数

结果:
24 B/op          2 allocs/op
#24 B/op表每一次调用需要分配16个字节
# 2 allocs/op表每一次调用有两次分配

开发者根据这些信息可快速找到可能的分配点，进行优化和调整
*/

/*
5.控制计时器
*/
func Benchmark_Add_TimerControl(b *testing.B){
	//重置计时器
	b.ResetTimer()
	//停止计时器
	b.StopTimer()
	//开始计时器
	b.StartTimer()
	var n int
	for i:=0;i<b.N;i++{
		n++
	}
	/*
	说明：从Benchmark()函数开始Timer就开始计数
	计数器内部不仅包含耗时数据，还包括内存分配的数据
	*/
}