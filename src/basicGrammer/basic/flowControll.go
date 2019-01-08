package basic

import (
	"fmt"
	"encoding/base64"
)

/*
流程控制:
常用:if和for
switch和goto主要为了简化代码，降低重复代码而生的结构--扩展类流程控制
*/
//1.if
func If(){
	array:=make([]string,2)
	if len(array)==0{ //注：该{必须与if同行，否则报错
		fmt.Println("no value")
	}else if len(array) > 0{
		fmt.Println("have value")
	}else{ //注：{必须与else同行，否则报错
		fmt.Println("error")
	}

	//可在if表达式前添加一个执行语句
	if err := returnVal(); err != "" {
		fmt.Println(err)
		return
	}
}
func returnVal() string{
	decodedMsg, err := base64.StdEncoding.DecodeString("c2VuZCBpbmZvIHRvIGthaXhpbnl1ZmVuZw==")
	if err!=nil{
		fmt.Println(err)
		return ""
	}
	return string(decodedMsg) //send info to kaixinyufeng
}

//2.构建循环(for)
func For(){
	//for循环可通过break,goto,return,panic语句强制退出循环

	//1)for中的初始语句--开始循环时执行的语句
	step:=2 //step放在for前面初始化(具有最大作用域)
	for ; step>0;step--{
		fmt.Println(step) //2,1
	}
	for step:=1;step>0;step--{
		fmt.Println(step) //1
	}

	//2)for中的条件表达式--控制是否循环的开关
	var i int
	for ; ; i++{ //无限循环
		if i>2{
			break //跳出for循环
		}
	}

	//更美观写法--无限循环 <=>与上for等价
	for{
		if i>2{
			break
		}
		i++
	}

	//只有一个循环条件的循环
	for i<=2 {
		i++
	}

	//九九乘法表
	for y:=1;y<=9;y++{
		for x:=1;x<=y;x++{
			fmt.Printf("%d*%d=%d ",x,y,x*y)
		}
		fmt.Println()//手动生成回车
	}

}

//3.键值循环(for range)--直接获得对象的索引和数据
/*
Go语言可使用for range遍历数组，切片，字符串，map及通道(channel)
for range遍历返回值规律：
（1）数组，切片，字符串返回索引和值
（2）map返回键和值
（3）通道(channel)只返回通道内的值
*/
func ForRange(){
	//(1)遍历数组，切片--获得索引和元素
	for key,value:=range[]int{1,2,3}{
		fmt.Printf("key:%d value:%d\n",key,value) //0 1 / 1 2 / 2 3
	}

	//(2)遍历字符串获得字符
	var str="kaixin玉凤"
	for key,value:=range str{
		fmt.Printf("key:%d value:0x%x\n",key,value)
		fmt.Printf("unicode:%c %d\n", value, value)//中文
	}

	//(3)遍历map--获得map的键和值
	mapVal:=map[string]int{
		"A":100,
		"B":200,
	}
	for key,value:=range mapVal{
		fmt.Println(key,value) //A 100 / B 200
	}

	//(4)遍历通道(channel)--接收通道数据
	//注：在通道遍历时只输出一个值，即管道内的类型对应的数据
	channelData:=make(chan int)
	go func() {
		//向通道添加数据1,2,3
		channelData <- 1
		channelData <- 2
		channelData <- 3
		close(channelData) //关闭通道
	}()
	for v:=range channelData{ //从通道中取数据，直到通道被关闭
		fmt.Println(v) // 1 / 2 /3
	}

	//(5)在遍历中选择希望获得的变量
	for _,value := range mapVal{
		fmt.Println(value) //只获取值，不需要索引，可用_(_即为匿名变量:一种占位符，不占用内存)
	}

}

//4.分支选择(switch)--拥有多个条件分支的判断
func Switch(){
	//(1)GO语言中的switch不仅可以基于常量进行为判断还可基于表达式进行判断
	var a="A"
	switch a {
	case "A":
		fmt.Println("A")
	case "B":
		fmt.Println("B")
	default: //只能有一个
		fmt.Println(0)
	}

	//(2)一分支多值
	switch a {
	case "A","B":
		fmt.Println("AB") //AB
	}

	//(3)分支表达式
	//case后不仅可以为常量也可以为表达式
	var r int=11
	switch { //***注意：此时switch后不能判断变量
	case r >10 && r<20:
		fmt.Println(r)
	}

	//(4)跨越case的fallthrough--兼容C语言的case设计
	var s = "kaixin"
	switch{
	case s == "kaixin":
		fmt.Println("kaixin")
		//使用fallthrough关键字，控制执行下一case语句（否则不会再执行下面的case语句）
		fallthrough //注：新编写的代码不建议使用fallthrough
	case s != "kaixin":
		fmt.Println("yufeng")
	}

}

//5.跳转到指定代码标签（goto）
//goto语句通过标签进行代码间的无条件跳转
func Goto(){
	//(1)使用goto退出多层循环
	for x:=0;x<10;x++{
		for y:=0;y<10;y++{
			if y == 2{
				//跳转到指定标签处
				goto breakHere
			}
		}
	}
	return //手动返回，避免执行进入标签
	//标签
	breakHere:
		fmt.Println("done")

	//(2)统一错误处理
	err:=returnVal()
	if err!=""{
		fmt.Println(err)
		goto exitProcess
		return
	}
	exitProcess:
	fmt.Println("done")
	//exit Process() //退出进程
}

//6.跳出指定循环break--可以跳出多层循环
func Break(){
	//(1)跳出指定循环
	OuterLoop:
	for i:=0;i<2;i++{
		for j:=0;j<5;j++{
			switch j{
			case 2:
				fmt.Println(i,j) //0 2
			break OuterLoop //退出outer loop对应的循环之外
			case 3:
				fmt.Println(i,j)
			break OuterLoop
			}
		}
	}
}

//7.继续下一次循环（continue）
//continue:结束当前循环，开始下一次循迭代过程(仅限在for循环内使用)
func Continue(){
OuterLoop:
	for i:=0;i<2;i++{
		for j:=0;j<5;j++{
			switch j{
			case 2:
				fmt.Println(i,j) //0 2   / 1 2
				continue OuterLoop //退出本次循环，进行下一次循环(最外层循环)
			case 3: //不会执行
				fmt.Println(i,j)
				continue OuterLoop
			}
		}
	}
}
