package 避坑与技巧

import (
	"fmt"
	"runtime"
	"net"
	"time"
	"sync"
)

/*
第12章 "避坑"与技巧

12.1 合理地使用并发特性

GO语言原生支持并发

12.1.1 了解goroutine的生命周期再创建goroutine

在GO语言中，开发者习惯将并发内容与goroutine一一对应地创建goroutine。

开发者很少考虑goroutine在什么时候能退出和控制goroutine生命周期。这会造成goroutine失控情况
*/
/*
一段耗时的计算函数
该函数模拟平时业务中放到goroutine中执行的耗时操作，该函数从其他goroutine中获取和接收数据/指令，处理后返回结果
*/
func consumer(ch chan int){
	//无限获取数据的循环
	for{
		//从通道获取数据
		data:=<-ch
		//打印数据
		fmt.Println(data)
	}
}
func Demo1(){
	//创建一个传递数据用的通道
	ch:=make(chan int)
	for{
		//空变量，什么也不做
		var dummy string
		//获取输入，模拟进程持续运行
		fmt.Scan(&dummy) //fmt.Scan()函数接收数据时，需要提供变量地址
		//启动并发执行consumer()函数
		go consumer(ch)
		//输出现在的goroutine数量
		fmt.Println("goroutines:",runtime.NumGoroutine())//每启动一个goroutine使用runtime.NumGoroutine()检查进程创建的goroutine数量总数
	}
	/*
	执行main函数,输入a,得到输出，输入b，得到输出...
	a
	goroutines: 2
	b
	goroutines: 3
	c
	goroutines: 4

	问题点：随着输入的字符串越来越多，goroutine将会无限制地被创建，但并不会结束。
	若在生产环境，会造成内存大量分配，最终使进程崩溃。
	*/
}

//优化Consumer
func consumernew(ch chan int){
	//无限获取数据的循环
	for{
		//从通道获取数据
		data:=<-ch
		if data == 0{ //添加合理的退出条件
			break
		}
		//打印数据
		fmt.Println(data)
	}
	fmt.Println("goroutine exit")
}
func Demo2(){
	//传递数据用的通道
	ch:=make(chan int)
	for{
		//空变量，什么也不做
		var dummy string
		//获取输入，模拟进程持续执行
		fmt.Scan(&dummy)
		if dummy == "quit"{
			for i:=0;i<runtime.NumGoroutine()-1;i++{ //main的goroutine也被算入,需要被减掉
				ch <- 0 //并发开启的goroutine都在竞争获取通道中的数据，因此只要知道有多少个goroutine需要退出，就给通道发多少个0
			}
			continue
		}
		//启动并发执行consumer()函数
		go consumernew(ch)
		//输出现在的goroutine数量
		fmt.Println("goroutines:",runtime.NumGoroutine())
	}
	/*
	执行结果: a/b/quit/c为输入

	a
	goroutines: 2
	b
	goroutines: 3
	quit
	goroutine exit
	goroutine exit
	c
	goroutines: 2
	*/
}


/*
12.1.2 避免在不必要的地方使用通道

通道channel与map,切片一样，由go源码编写而成。为了保证两个goroutine并发访问安全性，通道也需要做一些锁操作。
因此通道其实并不比锁高效
*/
/*
展示套接字的接收和并发管理.对于TCP来说，一般是接收过程创建goroutine并发处理。当套接字结束时，就要正常退出这些goroutine
*/
/*
1。套接字接收部分:套接字在连接后，就需要不停地接收数据
*/
//套接字接收过程
func sockeRecv(conn net.Conn,exitChan chan string){ //net.Conn为套接字接口，exitChan为退出发送同步通道
	//创建一个接收的缓冲
	buff:=make([]byte,1024)
	//不停地接收数据
	for{
		//从套接字中读取数据
		_,err:=conn.Read(buff)
		//需要结束接收，退出循环
		if err!=nil{
			break
		}
	}
	//函数己经结束，发送通知
	exitChan <- "recv exit"
}
/*
2。连接，关闭，同步goroutine主流程部分
*/
func Demo3(){
	//连接一个地址
	conn,err:=net.Dial("tcp","www.163.com:80") //net.Dial发起tcp协议连接
	//发生错误时打印错误并退出
	if err!=nil{
		fmt.Println(err)
		return
	}
	//创建退出通道
	exit:=make(chan string)
	//并发执行套接字接收
	go sockeRecv(conn,exit)
	//在接收时，等待1s
	time.Sleep(time.Second)
	//主动关闭套接字
	conn.Close() //此时会触发套接字接收错误
	//等待goroutine退出完毕
	fmt.Println(<-exit)
	/*
	注：goroutine退出使用通道来通知，此做法可以解决问题，但实际上通道中的数据并没有完全被使用
	*/
}

/*
优化：使用等待组替代通道简化同步
*/
//套接字接收过程
func sockeRecvnew(conn net.Conn,wg *sync.WaitGroup){
	//创建一个接收的缓冲
	buff:=make([]byte,1024)
	//不停地接收数据
	for{
		//从套接字中读取数据
		_,err:=conn.Read(buff)
		//需要结束接收，退出循环
		if err!=nil{
			break
		}
	}
	//函数已经结束，发送通知
	wg.Done() //将等待组计数器减1
}
func Demo4(){
	//连接一个地址
	conn,err:=net.Dial("tcp","www.163.com:80")
	//发生错误时打印错误并退出
	if err!=nil{
		fmt.Println(err)
		return
	}
	//退出通道
	var wg sync.WaitGroup
	//添加一个任务
	wg.Add(1)
	//并发执行接收套接字
	go sockeRecvnew(conn,&wg)
	//在接收时，等待1s
	time.Sleep(time.Second)
	//主动关闭套接字
	conn.Close()
	//等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}