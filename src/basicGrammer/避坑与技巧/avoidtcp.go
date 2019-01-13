package 避坑与技巧

import (
	"io"
	"bytes"
	"encoding/binary"
	"net"
	"fmt"
	"strconv"
	"sync"
	"bufio"
)

/*
第12章 "避坑"与技巧

12.5 优雅地处理TCP粘包

TCP协议是一种面向连接的，可靠的，基于字节流的传输通信协议

TCP特性：不丢包，并且保证按顺序到达

TCP数据在发送和接收时会形成粘包，也就是没有按照预期大小得到数据，数据包不完整。

*/
func Demo9(){
	io.ReadAtLeast(nil,nil,1) //该函数用于封包
	io.ReadFull(nil,nil) //该函数用于优雅地完成对TCP粘包的处理
}

//1.封包发送
//二进制封包格式
type Packet struct {
	Size uint16 //包体大小
	Body []byte //包体数据
}
//将数据写入dataWriter
func writePacket(dataWriter io.Writer,data []byte) error{
	//准备一个字节数组缓冲
	var buffer bytes.Buffer
	//将size写入缓冲
	if err:=binary.Write(&buffer,binary.LittleEndian,uint16(len(data)));err!=nil{
		return err
	}
	//写入包体数据
	if _,err:=buffer.Write(data);err!=nil{
		return err
	}
	//获得写入的完整数据
	out:=buffer.Bytes()
	//写入dataWriter
	if _,err:=dataWriter.Write(out);err!=nil{
		return err
	}
	return nil
}

/*
2.连接器

连接器可以通过给定的地址和发送次数，不断地通过Socket给地址对应的连接发送封包
*/
func connector(address string,sendTimes int){
	//尝试用socket连接地址
	conn,err:=net.Dial("tcp",address)
	//发生错误时退出
	if err!=nil{
		fmt.Println(err)
		return
	}
	//循环指定次数
	for i:=0;i<sendTimes;i++{
		//将循环序号转为字符串
		str:=strconv.Itoa(i) //将数字转为string
		//发送封包
		if err:=writePacket(conn,[]byte(str));err!=nil{
			fmt.Println(err)
			break
		}
	}
}

/*
3.接收器

连接器(connector)只能发起到接受器(Acceptor)的连接，一个接受器能接受多个来源连接
*/
//接受器
type Acceptor struct {
	//保存侦听器
	l net.Listener
	//侦听器的停止同步
	wg sync.WaitGroup
	//连接的数据回调
	OnSessionData func(net.Conn,[]byte)bool
}
//异步开始侦听
func (a *Acceptor) Start(address string){
	go a.listen(address)
}
func (a *Acceptor) listen(address string){
	//侦听开始，添加一个任务
	a.wg.Add(1)
	//在退出函数时，结束侦听任务
	defer a.wg.Done()
	var err error
	//根据给定地址进行侦听
	a.l,err=net.Listen("tcp",address)
	//如果侦听发生错误，打印错误并退出
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	//侦听循环
	for{
		//新连接没有到来时，Accept是阻塞的
		conn,err:=a.l.Accept()
		//若发生任何侦听错误，打印错误并退出服务器
		if err !=nil{
			break
		}
		//根据连接开启会话，这个过程需要并行执行
		go handleSession(conn,a.OnSessionData)
	}
}
//停止侦听器
func(a *Acceptor) Stop(){
	a.l.Close()
}
//等待侦听完全停止
func(a *Acceptor) Wait(){
	a.wg.Wait()
}
//实例化一个侦听器
func NewAcceptor() *Acceptor{
	return &Acceptor{}
}

/*
4.封包读取

封包读取是处理TCP封包中最重要的环节

该段TCP粘包代码逻辑使用goroutine+同步系统调用的方式编写。
处理逻辑比使用异步+回调的方式更清晰易读
*/
func readPacket(dataReader io.Reader)(pkt Packet,err error){
	//Size为uint16类型，占2个字节
	var sizeBuffer=make([]byte,2)
	//持续读取size直到读到为止
	_,err=io.ReadFull(dataReader,sizeBuffer)
	//发生错误时返回
	if err!=nil{
		return
	}
	//使用bytes.Reader读取sizeBuffer中的数据
	sizeReader:=bytes.NewReader(sizeBuffer)
	//读取小端的uint16作为size
	err=binary.Read(sizeReader,binary.LittleEndian,&pkt.Size)
	if err !=nil{
		return
	}
	//分配包体大小
	pkt.Body=make([]byte,pkt.Size)
	//读取包体数据
	_,err=io.ReadFull(dataReader,pkt.Body)
	return
}
/*
5.服务器连接会话

服务器在接受一个连接后就会进入会话(session)处理。
会话处理是一个循环，不停地从Socket中接收数据并通过readPacket()函数将数据转换为封包
*/
//连接的会话逻辑
func handleSession(conn net.Conn,callback func(net.Conn,[]byte)bool){
	//创建一个socket读取器
	dataReader:=bufio.NewReader(conn)
	//接收数据的循环
	for{
		//从连接读取封包
		pkt,err:=readPacket(dataReader)
		//回调外部
		if err!=nil || !callback(conn,pkt.Body){
			//回调要求退出
			conn.Close()
			break
		}
	}
}

/*
6.测试粘包处理
*/
func Demo10(){
	//测试次数
	const TestCount=1000
	//测试地址，若发生侦听冲突，请调整端口
	const address="127.0.0.1:8010"
	//接收到的计数器
	var recvCounter int
	//实例化一个侦听器
	acceptor:=NewAcceptor()
	//开启侦听
	acceptor.Start(address)
	//响应封包数据
	acceptor.OnSessionData= func(conn net.Conn, data []byte) bool {
		//转换为字符串
		str:=string(data)
		//转换为数字
		n,err:=strconv.Atoi(str)
		//任何错误或接收错位时，报错
		if err!=nil || recvCounter!=n{
			panic("failed") //panic()停止
		}
		//计数器增加
		recvCounter++
		//完成计数，关闭侦听器
		if recvCounter>=TestCount{
			acceptor.Stop()
			return false
		}
		return true
	}
	//连接器不断发送数据
	connector(address,TestCount)

	//等待侦听器结束
	acceptor.Wait() //测试未结束时，会持续阻塞，直到Acceptor被调用方法停止
}

