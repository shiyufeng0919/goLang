package 避坑与技巧

import (
	"github.com/go-martini/martini"
	"net/http"
	"testing"
	"reflect"
)

/*
第12章 "避坑"与技巧

12.2反射：性能和灵活性的双刃剑

现在一些流行设计思想需要建立在反射基础上。如IOC/DI

GO语言中web框架martini(https://github.com/go-martini/martini)即是通过DI依赖注入技术进行中间件的实现
*/
/*
martini框架搭建的http服务器

go get github.com/go-martini/martini

martini运行效率较低，主要因为大量应用了反射

一般I/O延迟远大于反射代码所造成的延迟。但更低的响应速度和更低的CPU占用依然是web服务器追求的目标
*/
func Demo5(){
	m:=martini.Classic()
	m.Get("/", func() string{ //响应路径/的代码使用一个闭包实现
		return "hello martini"
	})
	//提供请求和响应接口
	m.Get("/", func(res http.ResponseWriter,req *http.Request) {
	})
	m.Run()
}
/*
1.结构体成员赋值对比

反射经常被使用在结构体上，因此结构体的成员访问性能就成为了关注的重点
*/
/*
eg1:使用一个被实例化的结构体，访问它的成员，然后使用GO语言的基准化测试可以迅速测试出结果
*/
type data struct {
	Hp int
}
//原生结构体赋值过程
func BenchmarkNativeAssign(b *testing.B){
	//实例化结构体
	v:=data{Hp:2}
	//停止基准测试的计时器
	b.StopTimer()
	//重置基准测试计时器数据
	b.ResetTimer()
	//重新启动基准测试计时器
	b.StartTimer()
	//根据基准测试数据进行循环测试
	for i:=0;i<b.N;i++{
		//结构体成员赋值测试
		v.Hp=3
	}
}
//使用反射访问结构体成员并赋值过程
func BenchmarkReflectAssign(b *testing.B){
	v:=data{Hp:2}
	//取出结构体指针的反射值对象并取其元素
	vv:=reflect.ValueOf(&v).Elem()
	//根据名字取结构体成员
	f:=vv.FieldByName("Hp")
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i:=0;i<b.N;i++{
		//反射测试设置成员值性能
		f.SetInt(3)
	}
}

/*
2.结构体成员搜索并赋值对比
*/
func BenchmarkReflectFindFieldAndAssign(b *testing.B){
	v:=data{Hp:2}
	vv:=reflect.ValueOf(&v).Elem()
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i:=0;i<b.N;i++{
		//测试结构体成员的查找和设置成员的性能
		vv.FieldByName("Hp").SetInt(3)
	}
}

/*
3.调用函数对比
*/
//一个普通函数
func foo(v int){}

//对原生函数调用的性能测试
func BenchmarkNativeCall(b *testing.B){
	for i:=0;i<b.N;i++{
		//原生函数调用
		foo(0)
	}
}
func BenchmarkReflectCall(b *testing.B){
	//取函数的反射值对象
	v:=reflect.ValueOf(foo)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i:=0;i<b.N;i++{
		//反射调用函数
		v.Call([]reflect.Value{reflect.ValueOf(2)})
	}
	/*
	反射函数调用的参数构造过程非常复杂，构建很多对象会造成很大的内存回收负担。
	*/
}
/*
以上总结：
1。能使用原生代码时，尽量避免反射操作
2。提前缓冲反射值对象，对性能有很大帮助
3。避免反射函数调用，实在需要调用时，先提前缓冲函数参数列表，并且尽量减少地使用返回值
*/
