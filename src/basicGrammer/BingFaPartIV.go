package basicGrammer

import (
	"sync/atomic"
	"fmt"
	"sync"
	"net/http"
)

/*
第9章 并发

9.4同步--保证并发环境下数据访问的正确性

Go程序可以使用channel进行多个goroutine间的数据交换，但这仅仅是数据同步中的一种方法
channel内部的实现依然使用了各种锁。
*/
/*
9.4.1竞态检测--检测代码在并发环境下可能出现的问题

当多线程并发运行的程序竞争访问和修改同一块资源时，会发生竞态问题。
*/
//序列号生成器
var(
	//序列号
	seq int64
)
func genId() int64{
	//尝试原子的增加序列号
	atomic.AddInt64(&seq,1) //未将返回值作为genId函数返回值，目的：造成一个竞态问题
	return seq
}
func SyncDemo1(){
	//生成10个并发序列号
	for i:=0;i<10;i++{
		go genId()
	}
	fmt.Println(genId()) //单独调用一次genId()
}

/*
9.4.2 互斥锁(sync.Mutex)--保证同时只有一个goroutine可以访问共享资源
*/
//下述保证修改count值是一个原子过程。为修改count值操作加锁
var(
	count int //逻辑中使用的某个变量
	countGuard sync.Mutex //与变量对应的使用互斥锁
)
func getCount() int{
	//锁定
	countGuard.Lock()
	//在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}
func setCount(c int){
	countGuard.Lock()
	count=c
	countGuard.Unlock()
}
func syncDemo2(){
	setCount(1)//可以进行并发安全的设置
	fmt.Println(getCount())//可以进行并发安全的获取
}

/*
9.4.3 读写互斥锁(sync.RWMutex)--在读比写多的环境下比互斥锁更高效

在读多写少环境中，可优先使用读写互斥锁
sync包中的RWMutex提供了读写互斥锁的封装
*/
var(
	cnt int //逻辑中使用的某个变量
	cntGuard sync.RWMutex //与变量对应的使用读写互斥锁
)
//读取数据cn过程，适用于读写互斥锁
func getCn() int{
	cntGuard.RLock() //锁定(Rlock:表示将读写互斥锁标记为读状态,若此时别一个goroutine并发访问cnGuard，同时也调用了Rlock时，并不会发生阻塞)
	defer cntGuard.RUnlock() //在函数时退出锁定 (使用读模式解锁)
	return cnt
}

/*
9.4.4 等待组(sync.WaitGroup)--保证在并发环境中完成指定数量的任务

除了可使用channel和互斥锁进行两个并发程序间的同步外，还可以使用等待组进行多个任务的同步。
*/
func SyncDemo3(){
	var wg sync.WaitGroup //声明一个等待组,对一组等待任务只需要一个等待组
	var urls=[]string{
		"http://github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}
	//遍历这些地址
	for _,url:=range urls{
		//每一个任务开始时，将等待组加1
		wg.Add(1) //等待组的计数器加1
		//开启一个并发
		go func(url string) {
			//使用defer表示函数完成时将等待组减1
			defer wg.Done() //等待组的计数器减1
			//使用http访问提供的地址
			_,err:=http.Get(url)
			fmt.Println(url,err)
		}(url) //通过参数传递url地址(为了避免url变量通过闭包放入匿名函数后又被修改的问题)
	}
	//等待所有的任务完成
	wg.Wait() //等待组的计数器不等于0时阻塞直到变0
	fmt.Println("over")
}

