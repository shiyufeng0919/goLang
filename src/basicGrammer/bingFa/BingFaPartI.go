package bingFa

import (
	"fmt"
	"time"
	"runtime"
)

/*
第9章 并发

并发：指在同一时间内可以执行多个任务

并发编程包含：多线程编程(本示例)，多进程编程，分布式程序等

GO语言的并发通过goroutine特性完成。goroutine类似于线程，可根据需要创建多个goroutine并发工作

goroutine是由GO语言运行时调度完成，而线程是由操作系统调度完成

GO语言还提供channel在多个goroutine间通信

goroutine和channel是GO语言秉承的CSP(Communicating Sequential Process)并发模式的重要实现基础。
*/

/*
9.1轻量级线程(goroutine)--根据需要随时创建"线程"

goroutine类似于线程，但goroutine由GO程序运行时的调度和管理。
Go程序会智能地将gorountine中的任务合理地分配给每个CPU

go程序从main包的main()函数开始，在程序启动时，GO程序就会为main()函数创建一个默认的goroutine

所有goroutine在main()函数结束时会一同结束。

终止goroutine的最好方法就是自然返回goroutine对应的函数。
*/

/*
9.1.1 使用普通函数创建goroutine

Go程序中使用go关键字为一个函数创建一个goroutine。
一个函数可以被创建多个goroutine，一个goroutine必定对应一个函数
*/
/*
1。格式--为一个普通函数创建goroutine写法
go函数名(参数列表)
（1）函数名：要调用的函数名
（2）参数列表：调用函数需要传入的参数
使用go创建goroutine时，被调用函数的返回值会被忽略。
注：若要goroutine返回数据，则可利用通道channel特性，通过channel把数据从goroutine中人作为返回值传出
*/
/*
2。例子：使用go关键字，将running()函数并发执行，每隔1s打印一次计数器，而main的goroutine则等待用户输入，两个行为可同时进行。
*/
func running(){
	var times int
	//构建一个无限循环
	for{
		times++
		fmt.Println("tick",times)
		//延时1s
		time.Sleep(time.Second)
	}
}
func RunningDemo(){
	//并发执行程序
	go running()
	//接受命令行输入，不做作何事情
	var input string
	fmt.Scanln(&input)
}

/*
9.1.2 使用匿名函数创建goroutine

go关键字后也可为匿名函数或闭包启动goroutine

1.使用匿名函数创建goroutine的格式

go func(参数列表){
   函数体 //匿名函数代码
}(调用参数列表) //启动goroutine时需要向匿名函数传递的调用参数
*/
/*
2。使用匿名函数创建goroutine的例子
*/
func BingfaDemo1(){
	//创建匿名函数，并为匿名函数启动goroutine
	go func() {
		var times int
		for{
			times++
			fmt.Println("tick",times)
			time.Sleep(time.Second)
		}
	}() //因为匿名函数无参数，因此，此处参数列表为空
	var input string
	fmt.Scanln(&input)
}

/*
9.1.3 调整并发的运行性能(GOMAXPROCS)
*/
/*
runtime.GOMAXPROCS(逻辑CPU数量)
CPU数量值：
（1）<1:不修改任何值 （2）=1:单核心执行 （3）>1:多核并发执行
*/
func BingfaDemo2(){
	runtime.NumCPU() //查询CPU数量
	runtime.GOMAXPROCS(runtime.NumCPU())//设置
}

/*
9.1.4 理解并发和并行

1.并发(concurrency)：
把任务在不同的时间点交给处理器进行处理。在同一时间点，任务并不会同时执行(eg:先接电话，再吃饭)
2.并行(parallelism)
把每一个任务分配给每一个处理器独立完成，在同一时间点，任务一定是同时运行(eg:边接电话边吃饭)
*/

/*
9.1.5 GO语言的协作程序(goroutine)和普通协作程序(coroutine)

goroutine：可能发生在多线程环境,goroutine间使用channel通信，运行机制属于抢占式任务处理
coroutine：始终发生在单线程,coroutine使用yield和resume操作，运行机制属于协作式任务处理
*/