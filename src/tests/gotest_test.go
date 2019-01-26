package tests

import "testing"

/*
单元测试文件。原则：
- 文件名必须以`_test.go`结尾，这样在执行`go test`的时候才会执行到相应的代码
- 你必须import `testing`这个包
- 所有的测试用例函数必须是`Test`开头
- 测试用例会按照源代码中写的顺序依次执行
- 测试函数`TestXxx()`的参数是`testing.T`，我们可以使用该类型来记录错误或者是测试状态
- 测试格式：`func TestXxx(t*testing.T)`,`Xxx`部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如`Testintdiv`是错误的函数名
- 函数中通过调用`testing.T`的`Error`, `Errorf`, `FailNow`, `Fatal`,`FatalIf`方法，说明测试不通过，调用`Log`方法用来记录测试的信息
以下是我们的测试用例的代码。进入到该测试文件目录，执行命令$go test测试
*/
func Test_Division_1(t *testing.T) {
	//尝试单元测试在function
	if i,e:=Division(6,2);i!=3 || e!=nil{
		t.Error("除数测试未通过") //若不是预期，则报错
	}else{
		t.Log("第一个测试通过了") //记录一些你期望的信息
	}
}

func Test_Division_2(t *testing.T){
	if _,e:=Division(6,0);e!=nil{
		t.Error("Division not work as expected")
	}else{
		t.Log("one test passed",e)
	}
}

