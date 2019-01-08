package bingFa

import (
	"net"
	"fmt"
	"bufio"
	"strings"
	"os"
)

/*
第9章 并发

9.3 示例：Telnet回音服务器 -- TCP服务器的基本结构

Telnet协议是tcp/ip协议族中的一种，它允许用户(telnet客户端）通过一个协商过程与一个远程设备进行通信
*/
/*
1。接受连接

回音服务器能同时服务于多个连接
*/
//服务逻辑，传入地址和退出的通道
func server(address string,exitChain chan int){
	//根据给定地址进行侦听
	l,err:=net.Listen("tcp",address)
	//若侦听错误，则打印错误并退出
	if err != nil{
		fmt.Println(err.Error())
		exitChain <-1
	}
	//打印侦听地址，表侦听成功
	fmt.Println("listen:"+address)
	//延迟关闭侦听器
	defer l.Close()
	//侦听循环
	for{
		//新连接没有到来时，Accept是阻塞的
		conn,err :=l.Accept()
		//发生任何的侦听错误，打印错误并退出服务器
		if err!=nil{
			fmt.Println(err.Error())
			continue
		}
		//根据连接开启会话，这个过程需要并行执行
		go HandleSession(conn,exitChain)
	}
}
/*
2.会话处理

每个客户端连接处理业务的过程称为会话。

GO可根据实际会话数量创建多个goroutine，并自动调度它们的处理。
*/
//连接会话逻辑
func HandleSession(conn net.Conn,exitChain chan int){//HandleSession函数并发执行
	fmt.Println("session started。。。")
	//创建一个网络连接数据的读取器
	reader:=bufio.NewReader(conn)
	//接收数据的循环
	for{  //读取器读取封包，处理完封包后需要继续读取从网络发送过来的下一个封包
		//读取字符串，直到碰到回车返回
		str,err:=reader.ReadString('\n') //封包读取,内部会自动处理粘包过程,直到下一个回车符到达后返回数据
		//数据读取正确
		if(err == nil){
			//去掉字符串尾部回车
			str=strings.TrimSpace(str)
			//处理telnet命令
			if !processTelnetCommand(str,exitChain){
				conn.Close()
				break
			}
			//echo逻辑，发什么数据，原样返回
			conn.Write([]byte(str+"\r\n"))
		}else{
			//发生错误
			fmt.Println("session closed")
			conn.Close()
			break
		}
	}
}
/*
3.Telnet命令处理

telnet是一种协议，在OS中可在命令行使用telnet命令发起tcp连接
*/
func processTelnetCommand(str string,exitChain chan int) bool{
	//@close指令表示终止本次会话
	if strings.HasPrefix(str,"@close"){
		fmt.Println("session closed...")
		//告诉外部需要断开连接
		return false
		//@shutdown指令表示终止服务进程
	}else if strings.HasPrefix(str,"@shutdown"){
		fmt.Println("server shutdown")
		//往通道中写入0，阻塞等待接收方处理
		exitChain <-0
		//告诉外部需要断开连接
		return false
	}
	fmt.Println(str)
	return true
}
/*
4.程序入口 Telnet回音处理主流程
*/
func BingFaDemo1(){
	//创建一个程序结束码的通道
	exitChan:=make(chan int)
	//将服务器并发运行
	go server("127.0.0.1:7001",exitChan)
	//通道阻塞，等待接收返回值
	code:=<-exitChan
	//标记程序返回值并退出
	os.Exit(code)
}
/*
5.测试输入字符串

telnet 127.0.0.1:7001
hello
*/

/*
6.测试关闭会话
@close
*/
/*
7.测试关闭服务器
@shutdown
*/