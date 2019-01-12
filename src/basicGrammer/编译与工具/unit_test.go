package 编译与工具

import (
	"testing"
	"fmt"
)

/*
注意测试文件名必须为x_test.go
*/
func TestHelloWorld(t *testing.T){ //函数名必须以Test为前缀,参数为*testing.T的函数，一个单元测试文件可以有多个测试入口
	t.Log("hello unit test")
}

/*
执行:

(1) go test unit_test.go ##测试该文件里的所有测试用例

ZBMAC-C02VX5Y7H:编译与工具 shiyufeng$ go test unit_test.go
ok      command-line-arguments  0.005s

ok:表示测试通过
command-line-arguments：表测试用例需要用到的一个包名
0.005s:表测试花费的时间

(2) go test unit_test.go -v ##可以让测试时显示详细的流程

ZBMAC-C02VX5Y7H:编译与工具 shiyufeng$ go test unit_test.go -v
=== RUN   TestHelloWorld
--- PASS: TestHelloWorld (0.00s)
        unit_test.go:9: hello unit test
PASS
ok      command-line-arguments  0.005s

*/


//2.运行指定单元测试用例
func TestA(t *testing.T){
	t.Log("A")
}
func TestAK(t *testing.T){
	t.Log("B")
}
func TestB(t *testing.T)  {
	t.Log("C")
}

/*
go test -v -run TestA unit_test.go #TestA和TestAK均被执行，-run后跟随的测试用例名称支持正则表过式

go test -v -run TestA$ unit_test.go #TestA后加$，仅执行TestA用例
*/


//3.标记单元测试结果
//终止当前测试用例
//测试：go test -v -run TestFailNow$ unit_test.go
func TestFailNow(t *testing.T){
	t.FailNow()
}

//只标记错误不终止测试的方法
//测试：go test -v -run TestFail$ unit_test.go
func TestFail(t *testing.T){
	fmt.Println("before fail")
	t.Fail()
	fmt.Println("after fail")
}


//4。单元测试日志
/*
每个测试用例可能并发执行，使用testing.T提供的日志输出可以保证日志跟随这个测试上下文一起打印输出
*/