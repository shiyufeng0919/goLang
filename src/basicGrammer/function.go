package basicGrammer

import (
	"fmt"
	"strings"
	"flag"
)

/*
函数:
函数是组织好的，可重复使用的，用来实现单一或相关联的代码段，可以提高应用模块性和代码重复利用率
GO语言支持普通函数，匿名函数，闭包
注：
函数本身可作为值进行传递
支持匿名函数和闭包(closure)
函数可以满足接口

GO语言支持安全指针，也支持多返回值
*/
//1.普通函数声明形式
//func 函数名(参数列表)(返回参数列表){函数体}
/*
  func add(a,b int) int{  #a和b同类型,可一同定义int类型
     return a+b
  }
*/
func Test(){
	a,b:=test1()
	fmt.Println(a,b) //1 2
	a1,b1:=test2()
	fmt.Println(a1,b1) //1 2

	//3.调用函数
	fmt.Println(add(1,2)) //3

	//4.将秒解析为时间单位
	fmt.Println(resolveTime(1000))//将返回值打印 0 0 16
	_,hour,minute:=resolveTime(18000) //只获取小时和分钟
	fmt.Println(hour,minute) //5 300
	day,_,_:=resolveTime(90000) //只获取天
	fmt.Println(day) //62

	//5.GO语言中的参数值传递
	in:=Data{
		complax:[]int{1,2,3},
		instance:InnerData{
			5,
		},
		ptr:&InnerData{1},
	}
	//输出结构体的成员情况
	fmt.Printf("in Value:%v+\n",in) //{[1 2 3] {5} 0xc000096050}+
	//输出结构的指针地址
	fmt.Printf("in ptr:%p\n",&in) //0xc000084030
	//传入结构体，返回同类型的结构体
	out:=passByValue(in)
	//输出结构体的成员情况
	fmt.Printf("out value:%+v\n",out)//{complax:[1 2 3] instance:{a:5} ptr:0xc000096050}
	//输出结构体的指针地址
	fmt.Printf("out ptr:%p\n",&out)//0xc000084090

	//6.函数变量，把函数作为值保存到变量中
	var f func() //将变量f声明为func()类型，此时f即为"回调函数".此时f的值为nil
	f=fire //将fire()函数名作为值，赋值给f变量，此时f的值为fire函数
	f() //使用f变量进行函数调用，实际调用的是fire()函数,输出值：fire

	//7。字符串的链式处理--操作与数据分离的设计技巧
	removeGoPrefix:=removePrefix("golang")
	fmt.Println("removeGoprefix:",removeGoPrefix) //lang

	stringDealFlow() //字符串处理流程

	//8.匿名函数--没有函数名字的函数
	noNameFunc()

	//9.匿名函数用作回调函数
	visit([]int{1,2,3}, func(v int) {
		fmt.Println(v) //1 2 3
	})

	//10.使用匿名函数实现操作封装
	/*
	命令行测试：go run main.go --skill=fly
	*/
	NoNameFuncSeal()
}
//1。同一返回值类型
func test1()(int,int){
	return 1,2
}

//2。带有变量名的返回值
func test2()(a,b int){ //将两个整型返回值进行命名为a,b
	a=1
	b=2
	return  //<=>return a,b
}

//3.调用函数
func add(a,b int) int{
  return a+b
}

//4.将秒解析为时间单位
const(
	SecondsPerMinute=60 //定义每分钟秒数
	SecondsPerHour=SecondsPerMinute*60//定义每小时秒数
	secondsPerDay=SecondsPerMinute*24 //定义每天秒数
)
//将传入的秒解析为3种时间单位
func resolveTime(seconds int)(day int,hour int,minute int){
	day=seconds/secondsPerDay
	hour=seconds/SecondsPerHour
	minute=seconds/SecondsPerMinute
	return
}

//5.GO语言中的参数值传递
type Data struct{ //用于测试值传递效果的结构体
	complax []int //切片：整型切片类型（切片是一种动态类型，内部以指针存在）
	//测试切片在参数传递中的效果
	instance InnerData //结构体：声明为结构体类型（拥有多个字段的复杂类型）
	ptr *InnerData //指针：ptr声明为Inner Data的指针类型
}
type InnerData struct{
	a int
}
//值传递的测试函数
/*
注：Data的内存会被复制后传入函数
当函数返回时，又会将返回值复制一次，赋给函数返回值的接收变量
注：指针，切片和map等引用类型对象指向的内容在参数传递中不会发生复制，而是将指针进行复制，类似创建一次引用
*/
func passByValue(inFunc Data) Data{
	//输出参数的成员情况
	fmt.Printf("inFunc Value:%+v\n",inFunc)//%+v动词输出变量的详细结构 {complax:[1 2 3] instance:{a:5} ptr:0xc000096050}
	//打印inFunc的指针
	fmt.Printf("inFun ptr:%p\n",&inFunc) //0xc0000840c0
	return inFunc//将传入的变量作为返回值返回
}

//6。函数变量，把函数作为值保存到变量中
func fire(){
	fmt.Println("fire")
}

//7。字符串的链式处理--操作与数据分离的设计技巧
/*
链式处理：对数据的操作进行多步骤处理称为链式处理
数据与操作分离：SQL将数据的操作与遍历过程作为两个部分进行隔离，这样操作与遍历过程即可各自独立地进行设计。
链式处理器是一种常见的编程设计。
链式处理开发思想：将数据和操作拆分，解耦
另：Netty是使用java编写的一款异步事件驱动的网络应用程序框架，支持快速开发可维护的高性能的面向协议的服务器和客户端。
netty中就有类似的链式处理器的设计。
*/
//(1)字符串处理函数
//系统自定义处理函数，如：字符串大小写等
func stringProcess(list []string,chain []func(string)string){
	//遍历每个字符串
	for index,str:=range list{
		//第一个需要处理的字符串
		result:=str
		//遍历每一个处理链
		for _,proc:=range chain{
			//输入一个字符串进行处理，返回数据作为下一个处理链的输入
			result=proc(result)
		}
		//将结果放回切片
		list[index]=result
	}
}
//(2)自定义处理函数
func removePrefix(str string) string{
	//移除指定的go前缀
	return strings.TrimPrefix(str,"go")
}
//(3)字符串处理流程
func stringDealFlow(){
	//待处理的字符串列表
	list:=[]string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	//处理函数链
	chain:=[]func(string)string{
		removePrefix, //上述自定义函数，去掉go前缀
		strings.TrimSpace,//移除左右空格
		strings.ToUpper, //字符串转大写
	}

	//处理字符串
	stringProcess(list,chain) //传入字符串切片和处理链

	//输出处理好的字符串
	for _,str:=range list{
		fmt.Println(str) //SCANNER PARSER COMPILER PRINTER FORMATER
	}
}

//8.匿名函数--没有函数名字的函数
/*
匿名函数没有函数名，只有函数体。函数可以被作为一种类型被赋值给函数类型的变量。
匿名函数往往以变量方式被传递。
匿名函数经常被用于实现回调函数，闭包等。
匿名函数用途广泛，匿名函数本身是一种值，可以方便地保存在各种容器中实现回调函数和操作封装。
*/
func noNameFunc(){
	//(1)定义一个匿名函数
	/*
	   func(参数列表)(返回参数列表){
		 函数体
	   }
	*/
	func(data int) { //没有名字的函数为"匿名函数"
		fmt.Println("匿名函数：",data) //100
	}(100) //表示对匿名函数进行调用，传递参数为100

	//(2)将匿名函数赋值给变量
	f:=func(data int){
		fmt.Println("匿名函数赋值给变量：",data) //100
	}
	//使用f()调用匿名函数
	f(100)
}

//9.匿名函数用作回调函数
func visit(list []int,f func(int)){
	//遍历切片的每个元素，通过给定函数进行元素访问
	for _,v:=range list{
		f(v)
	}
}

//10.使用匿名函数实现操作封装
//定义命令行参数skill,从命令行输入skill即可将空格后的字符串传入skillParam指针变量
var skillParam=flag.String("skill","","skill to perform")
func NoNameFuncSeal(){

	//解析命令行参数，解析完成后，skillParam指针变量将指向命令行传入的值
	flag.Parse()

	//定义一个从字符串映射到func()的map，然后填充这个map
	var skill=map[string]func(){
		//初始化map的键值对，值为匿名函数
		"fire": func() {
			fmt.Println("fire")
		},
		"fly": func() {
			fmt.Println("fly")
		},
		"run": func() {
			fmt.Println("run")
		},
	}

	//skillParam是一个*skillParam类型的指针变量，用*skillParam获取命令行传过来的值,并在map中查找对应命令行指定的字符串函数
	if f,ok:=skill[*skillParam];ok{
		f() //若在map中存在，则调用
	}else{
		fmt.Println("skill not found")
	}
}