package basic

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"strings"
	"unicode/utf8"
	"os"
	"bufio"
	"time"
	"reflect"
)

//1。变量
//2。数据类型
//3。转换不同的数据类型
//4。指针
func Pointer() {
	/*
	   （1）指针地址：
	   	每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置
	       "&"操作符放在变量前对变量进行"取地址"操作
	       格式：ptr := &v //v的类型为T,ptr类型为*T,T为指针类型，*为指针
	   （2）指针类型
	   （3）指针取值

	     变量，指针，地址关系：每个变量都拥有地址，指针的值就是地址。
	*/
	//eg1:认识指针地址和指针类型
	var cat int = 1                 //整型cat变量
	var str string = "apple"        //字符串str变量
	fmt.Printf("%p %p", &cat, &str) //0xc000012060 0xc00000e1e0

	//eg2:从指针获取指针指向的值
	var name = "kaixin yufeng"
	ptr := &name                         //对name取地址 ptr为*T类型
	fmt.Printf("ptr type:%T\n", ptr)     //打印ptr类型 *string
	fmt.Printf("address:%p\n", ptr)      //打印ptr指针地址 0xc00000e1f0
	value := *ptr                        //对指针进行取值
	fmt.Printf("value type:%T\n", value) //取值后的类型 string
	fmt.Printf("value:%s\n", value)      //指针取值后就是指向变量的值 kaixin yufeng

	//eg3:使用指针修改值(应用指针进行数值交换)
	x, y := 1, 2
	swap(&x, &y)      //交换变量
	fmt.Println(x, y) //2,1

	//eg4:使用指针变量获取命令行的输入信息 (flag包应用)
	var mode = flag.String("mode", "", "process mode") //应用flag.String定义mode变量,类型为*String
	flag.Parse()                                       //解析命令行参数
	fmt.Println(*mode)                                 //打印mode指针所指向的变量

	//eg5:创建指针的另一种方法--new(类型)函数
	str2 := new(string)
	*str2 = "pointer"
	fmt.Println(*str2) //*在变量左边，代表指向变量的值pointer
}

func swap(a, b *int) { //a,b两个参数均为*int（*T）类型
	t := *a //取a指针的值，赋值给变量t
	*a = *b //取b指针的值，赋给a指针指向的变量
	*b = t  //将a指针的值赋给b指针指向的变量
}

//5。变量生命周期--变量能够使用的代码范围
/*
1)。栈：拥有特殊规则的线性表数据结构。后进先出（栈用于内存分配，栈的分配的回收速度非常快，出栈释放内存）
2)。堆：散列分布，适合不可预知大小的内存分配（分配速度慢，且会形成内存碎片）
3)。变量逃逸--自动决定变量分配方式，提高运行效率
GO将选用怎样的内存分配方式来适应不同的算法需求（eg:函数局部变量使用栈，结构体成员使用堆）这个过程整合到编译器中，命名为"变量逃逸分析"
4)。取地址发生逃逸
命令：进行内存分配分析(go run -gcflags "-m -l" main.go)

另：编译器判断变量应该分配在堆/栈上的原则：
A。变量是否被取地址  B。变量是否发生逃逸
*/

//6。字符串应用
func Str() {
	//1)计算字符串长度 len(string)->返回值type为int
	str1 := "abc"
	fmt.Println(len(str1)) //3
	str2 := "开心"
	fmt.Println(len(str2))                    //6 统计ASCII字符串长度用len()函数
	fmt.Println(utf8.RuneCountInString(str2)) //2 统计unicode字符串长度用utf8.RuneCountInString()函数

	//2)遍历字符串--获取每个字符串元素
	//A。遍历每一个ASCII字符（ascii字符遍历直接使用下标）
	str3 := "开心yufeng"
	for i := 0; i < len(str3); i++ {
		fmt.Printf("ascii:%c %d\n", str3[i], str3[i]) //动词%c %d ,汉字乱码
	}
	//B。按Unicode字符遍历字符串（unicode字符遍历用for range）
	str4 := "开心yufeng"
	for _, s := range str4 {
		fmt.Printf("unicode:%c %d\n", s, s) //中文不会乱码
	}

	//3)获取字符串的某一段字符(子串，substring)
	//A.strings.index() 正向搜索子字符串
	originStr := "开心玉凤,smile happy"           //定义原字符串
	indexStr := strings.Index(originStr, ",") //搜索,出现的位置
	fmt.Println(indexStr)                     //12(从0计算)
	newStr := strings.Index(originStr[indexStr:], "happy")
	fmt.Println(newStr) //7
	newStr1 := originStr[indexStr+newStr:]
	fmt.Println(newStr1) //happy

	//B.strings.lastIndex()反向搜索子字符串
	indexStr1 := strings.LastIndex(originStr, "玉凤")
	fmt.Println(indexStr1) //6

	//C.搜索的起始位置可通过切片偏移来制作
	//originStr[indexStr1:] 从indexStr1开始到后，形成slice
	fmt.Println(originStr[indexStr1:]) //玉凤,smile happy
	indexStr2 := strings.Index(originStr[indexStr1:], "smile")
	fmt.Println(indexStr2) //7

	//4)修改字符串 ：
	// 注：字符串无法直接修改每个元素，只能通过重新构造新的字符串并赋值给原来的字符串变量实现
	oldStr := "kaixin yufeng lala" //**字符串不可变（线程安全，只读对象，无须加索，内存共享）
	byteStr := []byte(oldStr)      //**定义byte数组，可变，本身是一个切片
	fmt.Println(string(byteStr))   //**string()将[]byte转为字符串，kaixin yufeng lala

	for i := 15; i < 18; i++ {
		byteStr[i] = 'A'
	}
	fmt.Println(string(byteStr)) //kaixin yufeng lAAA

	//5)连接字符串
	//A.使用+号连接字符串，低效
	str5 := "kaixin"
	str6 := "yufeng"
	str7 := str5 + " " + str6
	fmt.Println(str7) //kaixin yufeng

	//B.类似string builder机制进行高效字符串连接
	var sb bytes.Buffer  //声明字符串缓冲
	sb.WriteString(str5) //把字符串写入缓冲
	sb.WriteString(str6)
	str8 := sb.String() //将缓冲以字符串形式输出
	fmt.Println(str8)   //kaixinyufeng

	//6)格式化
	/*
	  fmt.Sprintf(格式化样式，参数列表...)
	     格式化样式：字符串形式，格式化动词以%开头
	     参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应
	*/
	//A.两个参数格式化
	var cnt1 = 2
	var cnt2 = 6
	title := fmt.Sprintf("共%d个,卖出%d个", cnt2, cnt1)
	fmt.Println(title) //共6个,卖出2个

	//B.将数值本身的格式输出
	pi := 3.1415926
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)
	fmt.Println(variant) //月球基地 3.1415926 true

	//C.匿名结构体声明，并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "yufeng",
		HP:   100000,
	}
	fmt.Printf("使用'%%+v' %+v\n", profile) //使用'%+v' &{Name:yufeng HP:100000}
	fmt.Printf("使用'%%#v' %#v\n", profile) //使用'%#v' &struct { Name string; HP int }{Name:"yufeng", HP:100000}
	fmt.Printf("使用'%%T' %T\n", profile)   //使用'%T' *struct { Name string; HP int }

	/*
	  常见动词：
	  %v:按值的本来值输出
	  %+v:在%v基础上，对结构体字段名和值进行展开
	  %#v:输出go语言语法格式的值
	  %T:输出GO语言语法格式的类型和值
	  %%:输出%本体
	  %b / %o / %d / %x / %X:类型以二/八/十/十六进制/十六进制字母大写方式显示
	  %U:unicode字符
	  %f:浮点数
	  %p:指针，十六进制方式显示
	*/
	//7)Base64编码--电子邮件的基础编码格式
	//base64编码是常见的对8比特字节码的编码方式之一。
	message := "send info to kaixinyufeng"
	encodedMsg := base64.StdEncoding.EncodeToString([]byte(message)) //编码消息(string转[]byte)
	fmt.Println(encodedMsg)                                          //c2VuZCBpbmZvIHRvIGthaXhpbnl1ZmVuZw==

	decodedMsg, err := base64.StdEncoding.DecodeString(encodedMsg) //解码消息
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decodedMsg)) //send info to kaixinyufeng
	}

	//8)从INI配置文件中查询需要的值
	//Step1:获取当前目录
	dirCurrent,err:=os.Getwd() //获取当前目录 **os包应用
	if err != nil{
		fmt.Println("获取当前目录错误：",err)
	}
	fmt.Println("dirCurrent:",dirCurrent)// /Users/shiyufeng/Documents/kaixinyufeng/goworkspace/src/golandProject/goLang
	//Step2:以当前目录为准，获取文件路径
	filename:=dirCurrent+"/src/source/go.ini"

	fetch:=getValue(filename,"remote \"origin\"","fetch")
	fmt.Println("fetch:",fetch)//+refs/heads/＊:refs/remotes/origin/＊

	core:=getValue(filename,"core","hide Dot Files")
	fmt.Println("core:",core) //dot Git Only

}

//从INI文件中取值的函数(filename:文件名,expectSection:期望读取的字段,expectKey期望读取段中的键)
//函数名小写，则文件内可见
func getValue(filename,expectSection,expectKey string) string{
  //B.读取文件
  //GO语言os包提供了文件打开函数os.Open()
  file,err:=os.Open(filename) //打开文件,成功打开会返回文件句柄
  if err!=nil{ //文件找不到，则getValue()会返回一空串,表无法从给定的INI文件中获取到需要的值
    return ""
  }
  defer file.Close() //defer:函数结束时，关闭文件(否则文件会发生占用，系统无法释放缓冲资源)

  //C.读取行文本
  reader:=bufio.NewReader(file) //使用读取器读取文件
  var sectionName string //当前读取的段的名字
  //读取循环，不断地读取文件中的每一行
  for {
	  //读取字符串，直到碰到\n即行结束
	  linestr, err := reader.ReadString('\n')
	  if err != nil {
		  break
	  }
	  //切掉行左右两边的空白字符
	  linestr = strings.TrimSpace(linestr)
	  //忽略空行
	  if linestr == "" {
		  continue
	  }
	  //忽略注释（ini文件行首为;表示注释）
	  if linestr[0] == ';' {
		  continue
	  }

	  //D.读取段和键值的代码
	  //读取段
	  //行首和行尾分别是方括号，说明是段标记的起始符
	  if linestr[0] == '[' && linestr[len(linestr)-1] == ']' {
		  //将段名取出(去掉[ 和 ])
		  sectionName = linestr[1 : len(linestr)-1]
	  } else if sectionName == expectSection { //读取键值
		  //切开等号分割的键值对 ,切开后值[key,value]
		  pair := strings.Split(linestr, "=")
		  //保证切开只有1个等号分割的键值情况
		  if len(pair) == 2 {
			  //去掉键的多余空白字符
			  key := strings.TrimSpace(pair[0])
			  if key == expectKey {
				  //去掉空白字符
				  keyVal := strings.TrimSpace(pair[1])
				  return keyVal
			  }
		  }
	  }
   }
	return ""
}

//7。常量--恒定不变的值
func Const(){

	//const定义一个常量值
	const i=1

	//const定义多个常量
	const(
		pi=3.1415926
		e=2.81
	)

	fmt.Println("pi:",pi,"e:",e)//pi: 3.1415926 e: 2.81

	//常量因为在编译期确定，所以可用数组声明
	const size=4
	var arr [size]int
	fmt.Println("arr:",arr)//arr:[0 0 0 0]

	//枚举--一组常量值(Go语言中没有枚举，可用常量配合iota模拟枚举)
	type Color int
	const(
		red=iota //开始生成枚举值，默认从0开始(iota常量值自动生成,起始为0)
		blue
		yellow
		black
	)
	fmt.Println(red,blue,yellow,black) //输出所有枚举值(0 1 2 3)
	var color Color=yellow //使用枚举值并赋初值
	fmt.Println(color)//2

	//应用iota来做强大的枚举常量值生成器---生成标志位常量
	const(
		FlagNone = 1 << iota //移位操作，每次将上一次的值左移一位
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n",FlagRed,FlagGreen,FlagBlue)//2 4 8 （按常量输出）
	fmt.Printf("%b %b %b\n",FlagRed,FlagGreen,FlagBlue)//10 100 1000（按二进制格式输出）

	//将枚举值转换为字符串
	fmt.Printf("%s %d",CPU,CPU)//CPU 1
}

type ChipType int //声明芯片类型
const(
	None ChipType=iota
	CPU //中央处理器
	GPU //图形处理器
)
//函数名必须为String()，当这个类型需要显示为字符串时，会自动寻找String90方法并进行调用
func(c ChipType) String() string{
	switch c {
	case None:
		return "None"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

//8。类型别名（Type Alias）
/*
Go1.9版本前：内建类型定义
type byte uint8
type rune int32
GO1.9版本后：内建类型定义
type byte=unit8
type rune=int32
*/
func TypeAlias(){
	//A.区分类型别名与类型定义
	/*
	  类型别名写法：type TypeAlias=Type
	*/
	//将NewInt定义为int类型(***没有等号)
	type NewInt int
	//将int取一个别名叫IntAlias
	type IntAlias=int //(***有等号，类型别名)
	//将a声明为NewInt类型
	var a NewInt
	//查看a的类型名(%T:输出GO语言语法格式的类型和值)
	fmt.Printf("a type:%T\n",a)//a type:basicGrammer.NewInt
	//将a2声明为IntAlias类型
	var a2 IntAlias
	//查看a2的类型名
	fmt.Printf("a2 type:%T\n",a2)//a2 type:int

	//B.非本地类型不能定义方法
	/*
	非本地方法：指的就是不能为不在一个包中的类型定义方法
	time.Durationd在time包，而本示例在basicGrammer包
	*/

	//C.在结构体成员嵌入时使用别名
	var veh Vehicle //声明变量veh为Vehicle类型
	veh.FakeBrand.Show() //指定调用FakeBrand的show
	reflectObj:=reflect.TypeOf(veh) //取veh的类型反射对象
	//遍历所有veh的成员
	for i:=0 ; i<reflectObj.NumField();i++{
		memberInfo:=reflectObj.Field(i) //veh的成员信息
		fmt.Printf("FiledName:%v,FileType:%v\n",memberInfo.Name,memberInfo.Type.Name())
		//FiledName:FakeBrand,FileType:Brand
		//FiledName:Brand,FileType:Brand
	}
}

//B.非本地类型不能定义方法:为MyDuration添加一个函数
//定义time.Duration别名MyDuration
//type MyDuration=time.Duration //此种定义会报错，不能为不在一个包中的类型定义方法
type MyDuration time.Duration //此种为正解，可将MyDuration别名定义在time包中
func (m MyDuration)EasySet(a string){}

//C.在结构体成员嵌入时使用别名
type Brand struct{} //定义商标结构
func(t Brand) Show(){} //为商标结构添加Show()方法
type FakeBrand=Brand //为Brand定义一个别名FakeBrand
type Vehicle struct {//定义车辆结构
    //嵌入两个结构
	FakeBrand
	Brand
}



