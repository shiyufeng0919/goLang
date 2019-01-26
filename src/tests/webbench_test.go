package tests

import (
	"testing"
	"errors"
)

/*
编写压力测试

压力测试用来检测函数(方法)的性能，和编写单元功能测试的方法类似

压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母。
func BenchmarkXXX(b *testing.B) { ... }
● go test 不会默认执行压力测试的函数，如果要执行压力测试需要带上参数。-test.bench，
语法:-test.bench="test_name_regex"，例如go test -test.bench=".*"表示测试全部的压力测试函数。
● 在压力测试用例中，请记得在循环体内使用testing.B.N，以使测试正常运行。
● 文件名也必须以_test.go结尾。
下面我们新建一个压力测试文件webbench_test.go，代码如下所示。
*/

func Divisions(a,b float64)(float64,error){
	if b==0{
		return 0,errors.New("除数不能为0")
	}
	return a/b,nil
}

func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		Divisions(4, 5)
	}
}
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能
	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		Divisions(4, 5)
	}
}
/**
我们执行命令go test -file webbench_test.go -test.bench=".*"，
可以看到如下结果
 */