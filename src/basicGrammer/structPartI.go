package basicGrammer

import (
	"fmt"
	"encoding/json"
)

/*
第6章 结构体(struct)

结构体是类型中带有成员的复合类型，GO语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性

Go语言中的类型可以被实例化，使用new或&构造的类型实例的类型是类型的指针
*/

//6.1定义结构体
/*
type 类型名 struct{
  字段1 字段1类型    //#成员变量，字段拥有自己的类型和值，字段名必须唯一，字段的类型可以为结构体
  字段2 字段2类型
  ......
}

//##同类型变量可以写为一行
type Color struct{
  R,G,B byte
}
*/

//6.2实例化结构体--为结构体分配内存并初始化
/*
结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正的分配内存。
因此必须在定义结构体并实例化后才能使用结构体的字段。

实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，
结构体实例与实例间的内存是完全独立的。
*/
//####6.2.1基本的实例化形式
/*
结构体本身是一种类型，可像int,string等类型一样，以var方式声明结构体即可完成实例化
var ins T //T为结构体类型  ins为结构体实例
*/
func structDemo1(){
	type point struct {
		X int
		Y int
	}
	//实例化结构体
	var p point //point为结构体类型，p为结构体实例
	p.X=10 //用.访问结构体成员变量
	p.Y=20
}


//####6.2.2创建指针类型的结构体
/*
GO语言中，可使用new对类型(包括struct,int,float,string等)进行实例化
结构体在实例化后会形成指针类型的结构体

ins:=new(T)
T:为类型，可为struct,int,string...
ins:T类型被实例化后保存到ins变量中，ins的类型是*T，属于指针
访问成员变量:ins.xxx <=> (*ins).xxx
*/
func structDemo2(){
	type player struct {
		Name string
		HealthPoint int
		MagicPoit int
	}
	thank:=new(player) //实例化结构体player
	thank.Name="Canon" //可用.直接访问结构体成员变量 <=>(*thank).Name
	thank.HealthPoint=100
}


//####6.2.3 取结构体的地址实例化(应用广泛)
/*
GO语言中，对struct进行"&"取地址操作时，视为对该类型进行一次new的实例化操作

ins:=&T{}
T:结构体类型
ins:结构体实例，类型为*T,是指针类型
*/
type command struct {
	Name string //指令名称
	Var *int //指令绑定的变量
	Comment string //指令的注释
}
func structDemo3(){
	var version int=1
	cmd:=&command{} //
	cmd.Name="version"
	cmd.Var=&version
	cmd.Comment="show version"
}
//函数封装上述初始化过程
func newCommand(name string,varref *int,comment string) *command{
	return &command{
		Name:name,
		Var:varref,
		Comment:comment,
	}
}
func newCommandDemo(){
	var version int=1
	cmd:=newCommand("version",&version,"showversion")
	fmt.Println(cmd)
}



//6.3初始化结构体的成员变量
/*
结构体实例化时可直接对成员变量进行初始化。
初始化两种形式：
（1）字段"键值对"形式：适合选择性填充字段较多的结构体
（2）多个值的列表形式：适合填充字段较少的结构体
*/

//####6.3.1 使用"键值对"初始化结构体
/*
1.键值对初始化结构体的书写格式
ins:=结构体类型名{
  field1:value1,
  field2:value2,
  ...
}
*/
func initStructDemo1(){
	//eg:爸爸的爸爸是爷爷
	type People struct {
		name string
		child *People //结构体的结构体指针字段，类型是*People
	}
	relation:=&People{  //relation为*People的实例
		name:"爷爷",
		child:&People{ //child在初始化时需要*People的值，使用取地址初始化一个People
			name:"爸爸",
			child:&People{
				name:"我",
			},
		},
	}
	jsonData,err:=json.Marshal(relation)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println(jsonData)

}


//####6.3.2使用多个值的列表初始化结构体
/*
Go语言可在"键值对"初始化的基础上忽略"键"
即可使用多个值的列表初始化结构体的字段
*/
/*
1。多个值列表初始化结构体的书写格式
ins:=结构体类型名{
 filed1Value,
 filed2Value,
 ...

注意：应用该格式初始化时注意
（1）必须初始化结构体的所有字段
（2）每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致
（3）键值对与值列表的初始化形式不能混用
}
*/

//2。多个值列表初始化结构体的例子
func initStructDemo2(){
	type Address struct {
		Province string
		City string
		ZipCode int
		Phonenumber string
	}
	addr:=Address{
		"heilongjiang",
			"haerbin",
			101,
		"0",
	}
	fmt.Println(addr)
}



//6.3.3初始化匿名结构体
/*
匿名结构体没有类型名称，无须通过type定义就可直接使用
*/

/*
1.匿名结构体定义格式和初始化方法
ins:=struct{
 //匿名结构体字段定义
  field1 type1
  field2 type2
  ...
}{
 //字段值初始化    ##可选
  initField1:field1Value,
  initField2:field2Value,
  ...
}
*/
//2.使用匿名结构体的例子
func printMsgType(msg *struct{
	id int
	data string
}){
	//使用动词%T打印msg的类型
	fmt.Printf("%T\n",msg) //*struct { id int; data string }
}
func noNameStructDemo(){
	//对匿名结构体进行实例化，同时初始化成员
	/*
	注：匿名结构体的类型名是结构体包含字段成员的详细白工木言。
	匿名结构体在使用时需要重新定义，造成大量重复代码，因此少用。
	*/
	msg:=&struct {
		//定义部分
		id int
		data string
	}{
		//值初始化部分
		1024,
		"hello",
	}
	printMsgType(msg)
}

func StructPartI(){
	initStructDemo1()
	initStructDemo2() //{heilongijang haerbin 101 0}
	noNameStructDemo()
}


//6.4构造函数--结构体和类型的一系列初始化操作的函数封装
/*
GO语言的类型或结构体没有构造函数的功能
结构体的初始化过程可以使用函数封装实现。
*/
//6.4.1多种方式创建和初始化结构体--模拟构造函数重载
type Cat struct{
	Color string
	Name string
}
func newCatByName(name string) *Cat{
	//取地址实例化猫的结构体
	return &Cat{
		Name:name,
	}
}
func newCatByColor(color string) *Cat{
	return &Cat{
		Color:color,
	}
}

//6.4.2 带有父子关系的结构体的构造和初始化--模拟父级构造调用
/*
GO语言中没有提供构造函数相关的特殊机制，用户根据自己需求，将参数使用函数传递到结构体构造参数中
即可完成构造函数的任务。
*/
type BlackCat struct {
	Cat //嵌入Cat，类似于派生,拥有Cat的所有成员，实例化后可访问Cat的所有成员
}
//构造基类
//newCat函数定义了Cat的构造过程
func newCat(name string) *Cat{
	return &Cat{
		Name:name,
	}
}
//构造子类
func newBlackCat(color string) *BlackCat{
	cat:=&BlackCat{} //实例化BlackCat结构，此时Cat也同时被实例化
	cat.Color=color
	return cat
}
