package basicGrammer

import (
	"fmt"
	"sync"
	"os"
	"errors"
)

/*
函数第II部分主要内容：
5.8:延迟执行语句(defer)
5.9:处理运行时发生的错误
5.10:宕机(panic)--
5.11：宕机恢复(recover)--
*/
//############1。延迟执行语句(defer)
/*
GO语言的defer语句会将其后面跟随的语句进行延迟处理。
*/
//(1)多个延迟执行语句的处理顺序
/*
在defer归属函数即将返回时，将延迟处理的语句按defer的逆序进行执行。
*/
func deferOrder(){
	fmt.Println("defer begin...")
	//将defer放入延迟调用栈
	defer fmt.Println(1) //先被defer的语句最后被执行
	defer fmt.Println(2)
	//最后一个放入，位于栈顶，最先调用
	defer fmt.Println(3) //最后被defer的语句最先被执行
	fmt.Println("defer end...")
	/*
     执行结果：
		defer begin...
		defer end...
		3
		2
		1
	*/
}

//(2)使用延迟执行语句在函数退出时释放资源
//如成对操作：打开/关闭文件，接收/回复请求，加/解锁等
//(2-1).使用延迟并发解锁
/*
演示在函数中并发使用map，为防止竞态问题，使用sunc.Mutex进行加锁
*/
var(
	valueByKey=make(map[string]int) //一个演示用的映射(map默认不是并发安全的)
	valueByKeyGuard sync.Mutex //保证使用映射时的并发安全的互斥锁
)
//根据键读取值
func readValue(key string) int{
	valueByKeyGuard.Lock()//对共享资源加锁
	v:=valueByKey[key] //取值
	valueByKeyGuard.Unlock()//对共享资源解锁
	return v //返回值
}
//使用defer语句简化readValue()
func deferReadValue(key string) int{
	valueByKeyGuard.Lock()
	//defer后面的语句不会马上调用，而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}


//(2-2)使用延迟释放文件句柄
/*
文件操作：打开文件，获取和操作文件资源，关闭资源等。
每一步操作均要有错误处理，每一步处理均会造成一次可能的退出。因此退出时需要释放资源
*/
//根据文件名查询其大小
func fileSize(filename string) int64{
   f,err := os.Open(filename) //根据文件名打开文件，返回文件句柄和错误
   if err != nil{ //若打开文件时发生错误，返回文件大小为0
   	 return 0
   }
   info,err:=f.Stat() //取文件状态信息
   if err != nil{ //若获取信息时发生错误，关闭文件并返回文件大小为0
   	 f.Close() //关闭文件（否则会发生资源泄漏）
   	 return 0
   }
   size:=info.Size()//取文件大小
   f.Close() //关闭文件
   return size //返回文件大小
}
//使用defer简化fileSize()
func deferFileSize(filename string) int64{
	f,err:=os.Open(filename)
	/*注意，defer f.Close()不能放在该位置,因为一旦文件打开错误，f将为空，在延迟语句触发时，将触发宕机错误*/
	if err != nil{
		return 0
	}
	defer f.Close() //延迟调用Close，此时Close不会被调用
	info,err:=f.Stat()
	if err != nil{
		//defer机制触发，调用Close关闭文件
		return 0
	}
	size:=info.Size()
	//defer机制触发，调用Close关闭文件
	return size
}
//demo
func DeferDemo(){
	deferOrder()
}



//############2.处理运行时发生的错误
//(2-1)net包中的例子
/*
net.Dial()是GO语言系统包net中的一个函数，一般用于创建一个socket连接
net.Dial有两个返回值，Conn和error。此函数时阻塞的。
*/
//(2-2)错误接口的定义格式
type erro interface {
	Error() string
}
//(2-3)自定义一个错误(erros包定义错误)
/*
注：错误字符串由于相对固定，一般在包作用域声明，应尽量减少在使用时直接使用erros.New()返回
*/
var err=errors.New("this is an error")
//(2-4)代码中使用错误定义
var errDivisionByZero=errors.New("division by zero")
func div(dividend,divisor int)(int,error){
	if divisor == 0{ //除数为0的情况并返回
		return 0,errDivisionByZero
	}
	//正常计算，返回空错误
	return dividend/divisor,nil
}
//(2-5)示例，在解析中使用自定义错误
/*
说明：使用erros.New定义的错误字符串的错误类型是无法提供丰富的错误信息的。
若要携带错误信息返回，则需借助自定义结构体实现错误接口
*/
type parseError struct {
	filename string //文件名
	line int  //行号
}
//实现error接口，返回错误描述
func (e *parseError) Error() string{
	return fmt.Sprintf("%s:%d",e.filename,e.line)
}
//创建一些解析错误
func newParseError(filename string,line int) error{
	return &parseError{filename,line}
}
func ErrorDemo(){
	var e error
	//创建一个错误实例，包含文件名和行号
	e=newParseError("main.go",1)
	//通过error接口查看错误描述
	fmt.Println(e.Error())
	//根据错误接口的具体类型，获取详细的错误信息
	switch detail:=e.(type) {
	case *parseError:
		//解析错误
		fmt.Println("Filename:%s line:%d\n",detail.filename,detail.line)
	default:
		fmt.Println("other error")

	}
}



