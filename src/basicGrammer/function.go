package basicGrammer

import "fmt"

/*
函数:
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
}
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