package basicGrammer

import (
	"fmt"
	"io"
)

/*
第7章 接口 (interface)
*/
/*
接口本身是调用方和实现方均需遵守的一种协议
*/
/*
7.1 声明接口
接口是双方约定的一种合作协议。接口是一种类型也是一种抽象结构，不会暴露所含数据的格式，类型及结构
*/
/*
7.1.1接口声明的格式

type 接口类型名 interface{ //GO语言声明接口类型名一般在后面加er如写接口writer
   方法名1(参数列表1) 返回值列表1
   方法名2(参数列表2) 返回值列表2
   ...
}
*/
type writer interface {
	Write([] byte) error
}
/*
7.2.2开发中常见的接口及写法
*/

//该接口可调用Write()方法写入一个字节数组，返回值为写入字节数和可能发生的错误
type Writer interface {
	Write(p []byte)(n int,err error)
}
//将一个对象以字符串形式展开的接口
type Stringer interface { //经常使用 <=>java/c#中的toString()
	String() string
}

/*
7.2实现接口的条件

接口定义后，需要实现接口，调用方才能正确编译通过并使用接口。
*/
/*
7.2.1 接口被实现的条件1：接口的方法与实现接口的类型方法格式一致

在类型中添加与接口签名一致的方法就可实现该方法
签名包括：方法中的名称，参数列表，返回参数列表
*/
//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error //参数，返回值
}
//定义文件结构，用于实现DataWriter
type file struct {}
//定义文件结构，用于实现DataWriter
func (d *file)WriteData(data interface{}) error{
	//模拟写入数据
	fmt.Println("Write Data:",data) //Write Data: data...
	return nil
}
func InterfaceDemo1(){
	//实例化file
	f:=new(file)
	//声明一个DataWriter的接口
	var writer DataWriter
	//将接口赋值f,即*file类型
	writer=f
	//使用DataWriter接口进行数据写
	writer.WriteData("data...")
}
/*
常见接口无法实现的错误
1。函数名不一致导致的报错
2。实现接口的方法签名不一致导致的报错
*/

/*
7.2.2接口中所有方法均被实现
当一个接口中有多个方法时，只有这些方法都被实现，接口才能被正确编译并使用
*/
//定义一个数据写入器
type DataWriter2 interface {
	WriteData(data interface{}) error
	//能否写入
	CanWrite()bool
}

/*
7.3 理解类型与接口的关系

类型和接口间有一对多 和 多对一关系
*/
/*
7.3.1 一个类型可以实现多个接口
*/
type Socket struct {
}
func (s *Socket) Write(p []byte)(n int,err error){ //实现了io.Writer接口
	return 0,nil
}
func (s *Socket) Close() error{ //实现了io.Closer接口
	return nil
}
//使用io.Writer的代码，并不知道socket和io.Closer的存存
func usingWriter(writer io.Writer){
	writer.Write(nil)
}
//使用io.Closer并不知道Socket和io.Writer的存存
func usingCloser(closer io.Closer){
	closer.Close()
}
func InterfaceDemo2(){
	//实例化socket
	s:=new(Socket)
	usingWriter(s)
	usingCloser(s)
}

/*
7.3.2多个类型可以实现相同的接口
*/
//一个服务需要满足能够开启和写日志的功能
type Service interface {//定义服务接口，一个服务需要实现Start()和Log()方法
	Start() //开启服务
	Log(string)//日志输出
}
//日志器
type Logger struct {}  //定义能输出日志的日志器结构
//实现service的log方法
func (g *Logger) Log(l string){ //为Logger添加Log方法,同时实现Service的Log方法

}
//游戏服务
type GameService struct {
	Logger //嵌入日志器(由于Log方法己被Logger结构体实现，所以此处只需要嵌入不需要再实现一遍)
}
//实现Service的start()方法
func (g *GameService) Start(){ //GameService实现Start()方法

}
func InterfaceDemo3(){
	var s Service=new(GameService) //实例化GameService
	s.Start() //由GameService实现
	s.Log("hello") //由Logger实现
}


/*
7.4示例：便于扩展输出方式的日志系统

本例搭建一个支持多种写入器的日志系统，可自由扩展多种日志写入设备
*/
/*
1.日志对外接口
*/
//声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}
//日志器
type logger struct {
	//这个日志器用到的日志写入器
	writerList []LogWriter
}
//注册一个日志写入器(注册即将日志写入器的接口添加到writerList)
func (l *logger) RegisterWriter(writer LogWriter)  {
	l.writerList=append(l.writerList,writer)
}
//将一个data类型的数据写入日志
func (l *logger) Log(data interface{}){
	//遍历所有注册的写入器
	for _,writer:=range l.writerList{ //遍历所有日志写入器
		//将日志输出到每一个写入器中
		writer.Write(data) //将本次内容写入日志写入器
	}
}
//创建日志器的实例
func NewLogger() *logger{
	return &logger{}
}
