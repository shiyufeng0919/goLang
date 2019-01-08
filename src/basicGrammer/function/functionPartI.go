package function

import (
	"fmt"
	"strings"
	"flag"
	"net/http"
	"os"
	"math"
	"bytes"
)

/*
第五章 函数
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

	//11。函数类型实现接口--把函数作为接口来调用
	//结构体实现接口
	execFunc()
	//函数体实现接口
	execFuncCaller()

	//12。闭包
	BiBao1()
	BiBao2()
	BiBao3()

	//13。可变参数
	testJoinStrings()
	testPrintTypeValue()
	testPrint()
}


//***********5.1声明函数*************************************************
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



//***********5.2函数变量--把函数作为值保存到变量中*************************************
//6。函数变量，把函数作为值保存到变量中
func fire(){
	fmt.Println("fire")
}



//***********5.3字符串的链式处理--操作与数据分离的设计技巧*************************************
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



//***********5.4匿名函数--没有函数名字的函数**********************************************
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



//***********5.5函数类型实现接口--把函数作为接口来调用**********************************************
//11。函数类型实现接口--把函数作为接口来调用
//(1)函数实现接口
//调用器接口
type Invoker interface { //接口
	Call(interface{}) //需要实现一个call方法,调用时会传入一个interface{}类型的变量，这种类型的变量表示任意类型的值
}
//****(2)结构体实现接口
//结构体类型
type Struct struct{
}
//实现Invoker的Call
//Call为结构体的方法
func (s *Struct) Call(p interface{}){ //实现类
	fmt.Println("from struct",p)
}

//声明接口变量
var invoker Invoker //声明Invoker类型（接口类型）的变量
func execFunc(){
	//实例化结构体
	s:=new(Struct) //new()将结构体实例化，也可写成: s:=&Struct
	//将实例化的结构体赋值到接口
	invoker=s //s类型为*Struct，在上述已经实现了Invoker接口类型，因此可以赋值给invoker
	//使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("kaixinyufeng") //调用结构体的Call方法,结果：from struct kaixinyufeng
}

//****(3)函数体实现接口
/*
函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体。
当类型方法被调用时，还需要调用函数本体
*/
//函数定义为类型
type FuncCaller func(interface{}) //将func(interface{})定义为FuncCaller类型
//实现Invoker的Call
func(f FuncCaller) Call(p interface{}){ //FuncCaller的Call方法将实现Invoker的Call()方法
	//调用f()函数本体
	f(p)
}
//声明接口变量
var invoker2 Invoker
func execFuncCaller(){
	//将匿名函数转为FuncCaller类型，再赋值给接口
	invoker2=FuncCaller(func(v interface{}){ //匿名函数func(v interface{}转换为FuncCaller类型，因上述FuncCaller类型实现了Invoker的Call方法，此处赋值成功
		fmt.Println("from function",v)
	})
	//使用接口调用FuncCaller.Call。内部会调用函数本体
	invoker2.Call("hello") //from function hello
}

//12.HTTP包中的例子
//(1)Http包中含有Handler接口定义,Handler用于定义每个Http请求和响应的处理过程
type Handler interface {
	ServeHttp(ResponseWriter,r *http.Request)
}

//(2)使用处理函数实现接口
type HandlerFunc func(w http.ResponseWriter,r *http.Request)
func(f HandlerFunc) ServeHTTP(w http.ResponseWriter,r *http.Request){
	f(w,r)
}

//(3)使用闭包实现默认的Http请求处理,可使用http.HandleFunc()函数




//***********5.6闭包（Closure）--引用了外部变量的匿名函数**********************************************
//13.闭包（Closure）--引用了外部变量的匿名函数
/*
闭包是引用了自由变量的函数，被引用的自由变量和函数一同存在，即使已经离开了自由变量的环境也不会被释放或删除。
在闭包中可继续使用这个自由变量。
即：函数+引用环境=闭包 //函数是编译期静态概念，而闭包是运行期动态概念
*/
//(1)在闭包内部修改引用的变量
func BiBao1(){
	//准备一个字符串,目的：用于修改
	str:="kaixin yufeng"
	//创建一个匿名函数
	foo:= func() {
		//匿名函数中访问str,str在匿名函数中并没有定义(在匿名函数前定义)，str被引用到了匿名函数中形成了闭包
		str="hello yufeng"
	}
	//调用匿名函数
	foo() //执行闭包，str发生修改
	fmt.Println("str:",str) //str: hello yufeng
}

//(2)闭包的记忆效应
/*
被捕获到闭包中的变量让闭包本身拥有了记忆效应。
闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在
闭包本身就如同变量一样拥有了记忆效应
*/
func BiBao2(){
	//创建一个累加器，初始值为1
	accumulator:=accumulate(1)
	//累加1并打印
	fmt.Println(accumulator)
	fmt.Println(accumulator)
	//打印累加器的函数地址
	fmt.Printf("%p\n",accumulator)
	//创建一个累加器，初始值为1
	accumulator2:=accumulate(10)
	fmt.Println(accumulator2)
	fmt.Printf("%p\n",accumulator2)
}
//提供一个值，每次调用函数会指定对值进行累加
func accumulate(value int) func() {
	//返回一个闭包
	//return func() int{
	//	//累加
	//	value++
	//	return value
	//}
	return func() {
		value++
	}
}

//(3)示例：闭包实现生成器
/*
闭包的记忆效应进程被用于实现类似于设计模式中工厂模式的生成器
*/
//创建玩家生成器
func playerGen(name string) func()(string,int){
	hp:=50 //血量一直为50
	//返回创建的闭包
	return func() (string, int) {
		//将变量引用到闭包中
		return name,hp //此匿名函数中未声明name和hp，直接引用外部.将hp和name变量引用到匿名函数中形成闭包
	}
}
func BiBao3(){
	//创建一个玩家生成器
	generator:=playerGen("highnoon")
	//返回玩家的血量和名字
	name,hp:=generator()
	fmt.Println(name,hp)//highnoon 50

}



//***********5.7可变参数--参数数量不固定的函数形式**********************************************
//14.可变参数--参数数量不固定的函数形式
/*
可变参数格式：
func 函数名(固定参数列表,v ... T)(返回参数列表){
   函数体
 }
说明：
（1）可变参数一般被放置在函数列表的末尾。前面是固定参数列表，当没有固定参数时，所有变量都是可变参数
（2）v为可变参数变量，类型为[]T，即拥有多个T元素的T类型切片。v和T间由...组成
（3）T为可变参数的类型，当T为interface{}时，传入的可以是任意类型
*/
//fmt包中的例子
//(1)所有参数都是可变参数:fmt.Println函数声明如下
func Println(a ...interface{}) (n int,err error){
	return fmt.Fprintln(os.Stdout,a...)
}
func KeBianParam(){
	//fmt.Println在使用时，传入的值类型不受限制
	fmt.Println(5,"hello",&struct {
		a int
	}{1},true)

	//fmt.Printf
	fmt.Printf("value:%v %f\n",true,math.Pi)
}
//(2)部分参数是可变参数：fmt.Printf
func Printf(format string,a ...interface{})(n int,err error){
	return fmt.Fprintf(os.Stdout,format,a...)
}
//(3)遍历可变参数列表--获取每一个参数的值
//定义一个函数，参数数量为0-n，类型约束为字符串
func joinStrings(slist ...string)string{//slist类型为[]string
    fmt.Println(len(slist)) //获取可变参数长度 3 / 2
	//定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer //bytes.Buffer<=>StringBuilder,可高效进行字符串连接操作
	//遍历可变参数列表slist ，类型为[]string
	for _,s:=range slist{
		//将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	//将连接好的字节数组转换为字符串并输出
	return b.String()
}
func testJoinStrings(){
	fmt.Println(joinStrings("kaixin","yufeng","fighting"))//kaixinyufengfighting
	fmt.Println(joinStrings("hello","go"))//hellogo
}
//(4)获得可变参数类型--获得每一个参数的类型
/*
当可变参数为interface{}类型时，可传入任何类型的值
*/
func printTypeValue(slist ...interface{}) string{
	//字节缓冲作为快速字符串连接
	var b bytes.Buffer
	//遍历参数
	for _,s:=range slist {
		//将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)//%v动词，可将interface{}格式的任意值转为字符串
		//类型的字符串描述
		var typeString string
		//对s进行类型断言
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}
		//写入字符串前缀
		b.WriteString("value:")
		//写入值
		b.WriteString(str)
		//写类型前缀
		b.WriteString("type:")
		//写类型字符串
		b.WriteString(typeString)
		//换行
		b.WriteString("\n")
	}
	return b.String()
}
func testPrintTypeValue(){
	fmt.Println(printTypeValue(100,"str",true))
}

//(5)在多个可变参数函数中传递参数
//示例：可变参数传递
//实际打印的函数
func rawPrint(rawList ...interface{}){
	//遍历可变参数切片
	for _,a:=range rawList{
		//打印参数
		fmt.Println(a)
	}
}
//打印函数封装
func print(slist ...interface{}){
	//将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)//在多个可变参数中传递需要加... 输出：1 2 3
	rawPrint("fmt",slist) //[1 2 3],此时slist将作为一个整体传入rawPrint
	/*
	注：可变参数使用...进行传递 与  切片间使用append连接是同一个特性
	*/
}
func testPrint()  {
	print(1,2,3)
}
