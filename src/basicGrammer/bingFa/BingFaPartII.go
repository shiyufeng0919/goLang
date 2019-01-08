package bingFa

import (
	"fmt"
	"time"
	"errors"
)

/*
第9章 并发

9.2 通道(channel)--在多个goroutine间通信息的管道

GO语言提倡使用通信的方法代替共享内存，这里通信的方法就是使用通道channel

channel就是一种队列一样的结构
*/
/*
9.2.1 通道的特性

GO语言中的channel是一种特殊类型。在任何时候，同时只能有一个goroutine访问通道进行发送和获取数据。
goroutine间通过channel即可通信

通道(channel)像一个传送带或队列，总是先入先出规则，保证收发数据顺序.
*/
/*
9.2.2声明通道类型

通道的元素类型就是其内部传输的数据类型
var 通道变量 chan 通道类型
(1)通道类型：通道内的数据类型 (2)通道变量:保存通道的变量
chan类型的空值为nil，声明后需配合make后才能使用
*/
/*
9.2.3创建通道

通道是引用类型，需要使用make创建.格式：
通道实例:=make(chan数据类型)
(1)数据类型：通道内传输的元素类型
(2)通道实例:通过make创建的通道句柄
*/
func channelDemo1(){
	ch1:=make(chan int) //创建一个int类型的通道
	ch2:=make(chan interface{})//创建一个空接口类型的通道，可以存放任意格式
	type Equip struct {
		name string
		model string
	}
	ch3:=make(chan *Equip)//创建Equip指针类型的通道,可以存放*Equip
	fmt.Println(ch1,ch2,ch3)
}
/*
9.2.4 使用通道channel发送数据

1.通道发送数据格式：
  通道变量<-值    //<-通道的发送使用特殊操作符
(1)通道变量：通过make创建好的通道实例
(2)值：可以为变量，常量，表达式，函数返回值等，值的类型必须与ch通道的元素类型一致

2.通过通道发送数据的例子
*/
func ChannelDemo2(){
	//创建一个空接口通道
	ch:=make(chan interface{})
	//将0放入channel中
	ch <- 0
	//将hello字符串放入通道中
	ch <- "hello"
}

/*
3.发判断将持续阻塞直到数据被接收
channelDemo2只发送向channel发送了数据，并未接收，则发判断操作将持续阻塞。会报错：
*/

/*
9.2.5 使用通道channel接收数据

channel接收特性：
1。通道的收发操作在不同的两个goroutine间进行
(由于channel的数据在没有接收方处理时，数据发送方会持续阻塞，因此channel的接收必定在别外一个goroutine中进行)
2。接收将持续阻塞直到发送方发送数据
3。每次接收一个元素

channel的数据接收如下4种写法:
1.阻塞接收数据
data:=<-ch //执行该语句会阻塞，直到接收到数据并赋值给data变量

2.非阻塞接收数据
data,ok:= <-ch //使用较少，占用CPU

3.接收任意数据，忽略接收的数据
<-ch //会发生阻塞,接收到的数据被忽略(该方式实际上只是通过channel在goroutine间阻塞收发实现并发同步)
*/
func ChannelDemo3(){
	ch:=make(chan int) //创建一个int类型的channel
	go func() { //开启一个并发匿名函数
		fmt.Println("start goroutine")
		ch <- 0 //通过channel通知main的goroutine
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	<- ch //等待匿名goroutine
	fmt.Println("all done")
	/*
	运行结果:
	wait goroutine
	start goroutine
	exit goroutine
	all done
	*/
}

/*
4。循环接收
for data:=range ch{ ... }
*/
func ChannelDemo4(){
	ch:=make(chan int) //创建一个channel
	go func() { //开启一个并发匿名函数
		for i:=3;i>=0;i--{ //从3循环到0
			ch <- i //发送3-0间的数值
			time.Sleep(time.Second) //每次发送完时等待
		}
	}()
	//遍历接收通道数据
	for data:=range ch{
		fmt.Println(data) //打印channel数据(3 2 1 0)
		if data==0{ //当遇到数据0时，退出循环
			break; //终止接收
		}
	}
}

/*
9.2.6 示例：并发打印(goroutine和channel放在一起用法)

上面例子创建的都是无缓冲channel，无缓冲channel，消息发送方与接收方均会出现阻塞

设计模式：生产者与消费者
*/
func printer(c chan int){ //消费者
	for{ //开始循环等待数据
		data:= <-c //从channel中获取一个数据
		if data == 0{ //将0视为数据结束
			break
		}
		fmt.Println(data)
		c <- 0 //通知main已结束循环（我搞定啦！）
	}
}
func ChannelDemo5(){
	c:=make(chan int) //创建一个channel(实现printer中的goroutine & ChannelDemo5中的goroutine通信)
	go printer(c) //并发执行printer,传入channel
	for i:=1;i<=10;i++{ //生产者
		c <-i //将数据通过channel投送给printer
	}
	c <- 0 //通知并发的printer结束循环(没数据啦！)
	<-c //等待printer结束(搞定喊我！)
}

/*
9.2.7单向通道--通道中的单行道

GO的通道可在声明约束其操作方向，如只发送/只接收。这种被约束方向的channel被称为单向channel

1.单向通道的声明格式

var 通道实例chan<-元素类型 //只能发送通道

var 通道实例<-chan元素类型 //只能接收通道

2。单向通道使用例子
*/
func ChannelDemo6(){
	ch:=make(chan int)
	var chSendOnly chan<-int=ch //声明一个只能发送的通道类型，并赋值为ch
	var chRecvOnly<-chan int=ch //声明一个只能接收的通道类型，并赋值为ch
	fmt.Println(chSendOnly,chRecvOnly)

	/*一个不能填充数据(发送)只能读取的channel是无意义的*/
}
/*
3.time包中的单向通道
*/
func ChannelDemo7(){
	timer:=time.NewTimer(time.Second)
	fmt.Println(timer)
}

/*
9.2.8 带缓冲的通道

在无缓冲通道基础上，为通道增加一个有限大小的存储空间形成带缓冲通道

带缓冲通道在发送时无需等待接收方接收即可完成发送过程，且不会发生阻塞，只有当存储空间满时才会发生阻塞。
同理，若缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读取时，通道将会再度阻塞。

带缓冲通道类比：快递柜，提高效率
*/
/*
1.创建带缓冲通道

通道实例:=make(chan通道类型,缓冲大小)
*/
func ChannelDemo8(){
	ch:=make(chan int,3) //创建一个3个元素缓冲大小的整型通道(使用缓冲channel，即便没有goroutine接收，发送者也不会发生阻塞)
	fmt.Println(len(ch)) //当前通道大小 0
	//发送3个int型元素到channel
	ch <-1
	ch <-2
	ch <-3
	fmt.Println(len(ch)) //查看通道大小 3
}

/*
2.阻塞条件

无缓冲通道=长度永远为0的带缓冲通道
(1)带缓冲通道被填满时，尝试再次发送数据时发生阻塞
(2)带缓冲通道为空时，尝试接收数据时发生阻塞

为什么GO语言对通道要限制长度？
限制通道长度有利于约束数据提供方的供给速度，供给数据量必须在消费方处理量+通道长度的范围内，才能正常处理数据。
*/

/*
9.2.9通道的多路复用--同时处理接收和发送多个通道的数据

"多路复用"：是通信和网络中的一个专业术语。通常表示在一个信道上传输多路信号或数据流的过程和技术。如电话：在一条通信线路上，即可接收也可发送数据

channel在接收数据时，若没有数据则会发生阻塞。可运用遍厣，但运行性能极差
for{
  data,ok:=<-ch1 //尝试接收ch1通道
  data,ok:=<-ch2 //尝试接收ch2通道
  ... //接收后绪通道
}

***GO语言提供了关键字select,可以同时响应多个通道,
select的每个case都会对应一个通道的收发过程,当收发完成时，就会触发case中响应的语句
多个操作在每次select中挑选一个进行响应.格式：

select{
   case操作1:
     响应操作1
   case操作2:
	 响应操作2
   ...
default:
   没有操作情况
}

select多路复用中可以接收的样式
操作				语句示例
接收任意数据		case<-ch
接收变量			case d:=<-ch
发送数据			case ch<-100
*/

/*
9.2.10 示例：模拟远程过程调用(RPC)

服务器开发中会使用RPC(Remote Procedure Call)远程过程调用简化进程间通信的过程

RPC能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数据调用一样

本示例：使用channel替代Socket实现RPC的过程。客户端与服务器运行在同一个进程，服务器和客户端在两个goroutine(线程)中运行

1.客户端请求和接收封装
*/
func RPCClient(ch chan string,req string)(string,error){
	//模拟socket向服务器发送请求（一个字符串信息）,服务器接收后结束阻塞执行下一行
	ch <- req
	//等待服务器返回
	select{ //select做多路复用，注：select后不接任何语句,而是将要复用的多个通道语句写在每一个case上
	   //下述两个case同时开始，看谁先返回就先执行谁
	   case ack:=<-ch: //取数据
	   	return ack,nil //接收到服务器返回数据
	   case <-time.After(time.Second) : //超时
	    return "",errors.New("time out")
	}
}
/*
2.服务器接收和反馈数据
*/
func RPCServer(ch chan string){
	for{ //构造一个无限循环，模拟执行完一个客户端继续执行下一个客户端
		data:=<-ch //接收客户端请求
		fmt.Println("server received:",data) //打印接收到的数据
		ch<-"success" //向客户端反馈已收到
	}
}

/*
3.模拟超时
*/
func RPCServer2(ch chan string){
	for{
		data:=<-ch
		fmt.Println("server received:",data)
		//模拟服务器延时，造成客户端超时
		time.Sleep(time.Second*2) //通过睡眠函数让程序(goroutine)执行阻塞2秒的任务
		ch<-"success"
	}
}
/*
4.主流程
*/
func ChannelDemo9(){
	ch:=make(chan string) //创建一个无缓冲字符串通道(模拟网络和socket概念,即可从channel接收也可发送数据)
	//并发执行服务器逻辑
	go RPCServer(ch)
	//客户端请求数据和接收数据
	recv,err:=RPCClient(ch,"hi")
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("client received",recv)
	}
	/*
	执行结果:
	server received: hi
    client received success
	*/
}

/*
9.2.11 示例：使用通道响应计时器的事件

GO语言中的time包提供了计时器的封装

1.一段时间后(time.After)
*/
func ChannelDemo10(){
	exit:=make(chan int) //声明一个退出用的通道
	fmt.Println("start...")
	time.AfterFunc(time.Second, func() {
		fmt.Println("one second after")
		exit <- 0 //通知main()的goroutine已经结束
	}) //过1秒后调用匿名函数
	<-exit //等待结束
}
/*
2.定点计时

计时器(Timer)原理和倒计时类似，均为给定多少时间后触发

打点器(Ticker)原理和钟表类似，钟表每到整点就会触发
*/
func ChannelDemo11(){
	//创建一个打点器，每500ms触发一次
	ticker:=time.NewTicker(time.Microsecond*500)
	//创建一个计时器，每2s后触发
	stopper:=time.NewTimer(time.Second*2)
	//声明计数变量
	var i int
	//不断检查通道情况
	for{
		//多路复用通道
		select {
		   case <-stopper.C:
		   	fmt.Println("stop") //计时器到时了
		   	goto StopHere //跳出循环
		   case <-ticker.C:
		   	i++
		   	fmt.Println("tick",i) //打点器触发了
		}
	}
	StopHere: //退出的标签，使用goto跳转
		fmt.Println("done")
}
/*
9.2.12关闭通道后继续使用通道
通道是一个引用对象，和map类似，map在没有任何外部引用时，GO程序在运行时(runtime)会自动对内存进行垃圾回收GC
通道可以被GC，但通道也可被主动关闭。
*/
/*
1.格式：使用close关闭一个通道
close(ch) //关闭的channel依然可以被访问，但会有些问题

2.给被关闭通道发送数据将会触发panic(关闭的channel不会被置为nil)
*/
func ChannelDemo12(){
	ch:=make(chan int) //创建一个int型channel
	close(ch) //关闭channel,不会将ch置为nil
	//打印channel指针，容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n",ch,cap(ch),len(ch)) //ptr:0xc000082120 cap:0 len:0
	ch<-1 //给关闭的channel发送数据
}
/*
3.从已关闭的通道接收数据时将不会发生阻塞
*/
func ChannelDemo13(){
	//创建一个int型带2个缓冲的channel
	ch:=make(chan int,2)
	//给channel放入两个数据,此时channel装满了
	ch <-0
	ch <-1
	//关闭channel,此时带缓冲channel的数据不会被释放，通道也没消失
	close(ch)
	//遍历缓冲所有数据，且多遍历1个,故意造成channel的越界访问
	for i:=0;i<cap(ch)+1;i++{
		//从channel中取出数据
		v,ok:=<-ch
		fmt.Println(v,ok)
	}
	/*
	打印结果:
		0 true
		1 true
		0 false //0为channel的默认值,false表示没有获取成功,因为此时channel已经空了.此时channel关闭，即便channel没有数据，在获取数据时也不会发生阻塞，但此时取出数据失败
	*/
}