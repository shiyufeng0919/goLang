package reflect

import (
	"reflect"
	"fmt"
	"encoding/asn1"
)

/*
第10章 反射

反射是指在程序运行期对程序本身进行访问和修改的能力

GO程序在运行期使用reflect包访问程序的反射信息
*/

/*
##10.1反射的类型对象(reflect.Type)
*/
func ReflectDemo1(){
	var a int
	typeOfA:=reflect.TypeOf(a) //该函数可获得任意值的类型对象（typeOfA类型为reflect.Type()）
	fmt.Println(typeOfA.Name(),typeOfA.Kind()) //程序通过类型对象可访问任意值的类型信息。类型名:int 种类:int
}
/*
10.1.1理解反射类型(Type)与种类(Kind)

编程中使用最多的是Type，在反射中需要区分一个大品种类型时，会有Kind(eg:需要统一判断类型中的指针)

1.反射种类(Kind)的定义
eg:
type A struct{} A是struct{}类型，定义的结构体属于Struct种类，*A属于指针

另：Map,slice,chan属于引用类型，使用起来类似于指针，但在种类常量定义中仍然属于独立种类，不属于指针

2.从类型对象中获取类型名称和种类的例子
*/
type Enum int //定义一个Enum类型
const(
	Zero Enum=0
)
func ReflectDemo2(){
	//声明一个空结构体
	type cat struct {
	}
	//获取结构体"实例"的反射类型对象
	typeOfCat:=reflect.TypeOf(cat{})
	//显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(),typeOfCat.Kind()) //cat struct

	//获取Zero常量的反射类型对象
	typeofZero:=reflect.TypeOf(Zero)
	//显示反射类型对象的名称和种类
	fmt.Println(typeofZero.Name(),typeofZero.Kind()) //Enum int
}

/*
10.1.2 指针与指针指向的元素

GO程序中对指针获取反射对象时，通过reflect.Elem()方法获取这个指针指向的元素类型。
这个获取过程被称为取元素，等效于对指针类型变量做了一个"*"操作。
*/
func ReflectDemo3(){
	//声明一个空结构体
	type cat struct {
	}
	//创建cat的实例,ins是*cat的指针类型
	ins:=&cat{}
	//获取结构体实例(指针变量)的反射类型对象
	typeofCat:=reflect.TypeOf(ins)
	//显示反射类型对象(指会变量类型)的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n",typeofCat.Name(),typeofCat.Kind()) //name:'' kind:'ptr' ,指针变量类型名称是''而不是*cat

	//取类型的元素
	typeofCat=typeofCat.Elem()

	//显示反射类型对象名称和种类
	fmt.Printf("element name:'%v',element kind:'%v'\n",typeofCat.Name(),typeofCat.Kind())//element name:'cat',element kind:'struct'
}

/*
10.1.3 使用反射获取结构体成员类型

1。结构体字段类型
*/
type structField struct {
	Name string //字段名
	PkgPath string  //字段路径
	Type reflect.Type //字段反射类型对象
	Tag reflect.StructTag //字段的结构体标签
	Offset uintptr //字段在结构体中的相对偏移
	Index []int //Type.FieldByIndex中的返回的索引值
	Anonymous bool //是否为匿名字段
}
/*
2.获取成员反射信息
*/
func ReflectDemo4(){
	type cat struct {
		Name string
		Type int `json:"type" id:"100"` //``这个字符串在GO语言中被称为Tag(标签)
	}
	//创建cat实例
	ins:=cat{Name:"mimi",Type:1} //注：结构体标签属于类型信息，无须且不能赋值
	//获取结构体实例的反射类型对象
	typeOfCat:=reflect.TypeOf(ins)
	//遍历结构体所有成员
	for i:=0;i<typeOfCat.NumField();i++{ //typeOfCat.NumField()获取一个结构体类型共有多少个字段，注：类型必须为结构体，否则会宕机
		//获取每个成员的结构体字段类型
		fileType:=typeOfCat.Field(i) //fileType为struct Field结构体
		//输出成员名和tag
		fmt.Printf("name:%v tag:'%v' \n",fileType.Name,fileType.Tag)
	}
	//通过字段名，找到字段类型信息
	if catType,ok:=typeOfCat.FieldByName("Type");ok{
		//从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"),catType.Tag.Get("id"))
	}
	/*
	打印结果:
	name:Name tag:''
	name:Type tag:'json:"type" id:"100"'
	type 100
	*/
}
/*
10.1.4 结构体标签(Struct Tag) --对结构体字段的额外信息标签

通过reflect.Type获取结构体成员信息reflect.StructField结构中的Tag被称为结构体标签(Struct Tag)

JSON,BSON等格式进行序列化及ORM对象关系映射系统都会用到结构体标签

<=> C#
[Conditional("DEBUG")]
public static void Message(string msg){
   Console.WriteLine(msg)
}

1.结构体标签的格式

`key1:"value1" key2:"value2"` ##注：key与value间:分隔无空格

2.从结构体标签中获取值

（1）func (tag reflect.StructTag) Get(key string) string ：根据tag中的键获取对应的值
（2）func (tag reflect.StructTag) Lookup(key string)(value string,ok bool)：根据tag中的键，查询值是否存在

3.结构体标签格式错误导致的问题 ##注：key与value间:分隔无空格
*/

/*
10.2 反射的值对象 (reflect.Value)

反射不仅可以获取值的类型信息，还可动态获取或设置变量的值

10.2.1 使用反射值对象包装任意值
*/
func ReflectDemo5(){
	value:=reflect.ValueOf(asn1.RawValue{}) //返回reflect.Value类型，包括RawValue的值信息
	//reflect.Value是一些反射操作的重要类型,如反射调用函数
	fmt.Println(value)
}
/*
10.2.2 从反射值对象获取被包装的值

GO语言可以通过reflect.Value重新获取原始值

1.从反射值对象(reflect.Value)中获取值的方法

2.从反射值对象(reflect.Value)中获取值的例子
*/
func ReflectDemo6(){
	var a int=1024
	//获取变量a的反射值对象
	valueOfA:=reflect.ValueOf(a)
	//获取interface{}类型的值，通过类型断言转换
	var getA int=valueOfA.Interface().(int) //将valueofA反射值对象以interface{}类型取出，通过类型断言转换为int类型并赋值给getA
	//获取64位的值，强制类型转换为int类型
	var getA2 int=int(valueOfA.Int()) //将valueOfA反射值对象通过Int方法以int64类型取出，通过强制类型转换，转换为原本的int类型
	fmt.Println(getA,getA2) //1024 1024
}
/*
10.2.3 使用反射访问结构体的成员字段的值
*/
//定义结构体,每个字段类型不一致
type dummy struct {
	a int
	b string
	//嵌入字段
	float32
	bool

	next *dummy
}
func ReflectDemo7(){
	//实例化结构体，并包装结构体为reflect.Value类型
	d:=reflect.ValueOf(dummy{
		next:&dummy{},
	})

	//获取结构体字段数量
	fmt.Println("NumField",d.NumField()) //NumField 5

	//获取索引为2的字段(float32字段)
	floatField:=d.Field(2)
	//输出字段类型
	fmt.Println("Field",floatField.Type()) //Field float32

	//根据名字查找字段(根据b字符串查找b字段的类型)
	fmt.Println("FieldByName(\"b'\").Type",d.FieldByName("b").Type()) //FieldByName("b'").Type string

	//根据索引查找值中next字段的int字段的值 (4:表示在dummy结构体中索引值为4的成员，即next;0:继续在next值的基础上索引)
	fmt.Println(d.FieldByIndex([]int{4,0}).Type()) //int
}

/*
10.2.4 反射对象的空和有效性判断

IsNil:常被用于判断指针是否为空
IsValid：常被用于判断返回值是否有效
*/
func ReflectDemo8(){
	//*int的空指针 初始值为nil
	var a *int
	fmt.Println("var a *int:",reflect.ValueOf(a).IsNil()) //true

	//nil值
	fmt.Println("nil:",reflect.ValueOf(nil).IsValid()) //false

	//*int类型的空指针 ((*int)(nil):将nil转换为*int,即*int类型的空指针)
	fmt.Println("(*int)(nil):",reflect.ValueOf((*int)(nil)).Elem().IsValid()) //false

	//实例化一个结构体
	s:= struct {

	}{}

	//尝试从结构体中查找一个不存在的字段 (空字段)
	fmt.Println("不存在的结构体成员:",reflect.ValueOf(s).FieldByName("").IsValid()) //false

	//尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的方法:",reflect.ValueOf(s).MethodByName("").IsValid()) //false

	//实例化一个map
	m:=map[int]int{}

	//尝试从map中查找一个不存在的键
	fmt.Println("不存在的键:",reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid()) //false
}

/*
10.2.5 使用反射值对象修改变量的值

1.判断及获取元素的相关方法

使用reflect.Value取元素，取地址及修改值的属性方法:

Elem() Value
Addr() Value
CanAddr() bool
CanSet() bool

2.值修改相关方法

使用reflect.Value修改值的相关方法
SetInt(x int64)
SetUint(x uint64)
SetFloat(x float64)
SetBool(x bool)
SetBytes(x []byte)
SetString(x string)

3.值可被修改条件之一：可被寻址

通过反射修改变量值的前提条件之一：这个值必须可以被寻址（即这个变量必须能被修改）

注：reflect.Value的Addr()方法类似语言层的&操作，Elem()方法类似语言层的*操作。但这个方法与语言屋操作不等效
*/
func ReflectDemo9(){
	var a int=1024
	//获取变量a的反射值对象
	valueOfA:=reflect.ValueOf(a) //传入的是a的值，而不是a的地址
	//尝试将a修改为1(此时会发生崩溃)
	valueOfA.SetInt(1) //SetInt()正在使用一个不能被寻址的值。
}
//修改值正确方法
func ReflectDemo10(){
	//正确方法
	var b int =1024
	//获取b的反射值对象(a的地址)
	valueOfB:=reflect.ValueOf(&b) //valueOfB持有b的地址
	//取出b地址的元素
	valueOfB=valueOfB.Elem() //Elem取出b地址的元素.返回值类型reflect.Value
	//修改b的值为1
	valueOfB.SetInt(1)
	//打印b的值
	fmt.Println(valueOfB.Int()) //1
}

/*
4.值可修改条件之一：被导出

结构体成员中，若字段没有被导出，即便不使用反射也可被访问，但不能通过反射修改

值的修改表面叫"可寻址"，即值必须"可被设置"。若想修改变量值，一般步骤：
（1）取这个变量的地址/这个变量所在结构体已经是指针类型
（2）使用reflect.ValueOf进行值包装
（3）通过value.Elem()获得指针指向的元素值对象(Value)。因为值对象(Value)内部对象为指针时，使用set设置时会报出宕机错误
（4）使用Value.set设置值
*/
func ReflectDemo11(){
	type dog struct {
		legCount int //注意此处首字母小写
	}
	//获取dog实例的反射值对象
	valueOfDog:=reflect.ValueOf(dog{})
	//获取legCount字段的值
	vLegCount:=valueOfDog.FieldByName("legCount")
	//尝试设置legCount的值（这里会发生崩溃）
	vLegCount.SetInt(4) //报错原因：SetInt使用的值来自于一个未导出的字段
}
//正确写法
func ReflectDemo12()  {
	type dog struct {
		LegCount int //注意首字母大写,导出该字段
	}
	//获取dog实例地址的反射值对象
	valueOfDog:=reflect.ValueOf(&dog{}) //地址
	//取出dog实例地址的元素
	valueOfDog=valueOfDog.Elem()
	//获取LegCount字段的值
	vLegCount:=valueOfDog.FieldByName("LegCount")
	//尝试设置LegCount的值
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int()) //4
}

/*
10.2.6 通过类型创建类型的实例
*/
func ReflectDemo13(){
	var a int
	//取变量a的反射类型对象
	typeOfA:=reflect.TypeOf(a)
	//根据反射类型对象创建类型实例
	aIns:=reflect.New(typeOfA)
	//输出Value的类型和种类
	fmt.Println(aIns.Type(),aIns.Kind()) //*int ptr
}

/*
10.2.7 使用反射调用函数
*/
func add(a,b int) int{
	return a+b
}
//不建议大量使用反射函数调用
func ReflectDemo14(){
	//将add函数包装为反射值对象
	funcValue:=reflect.ValueOf(add)

	//构造函数参数，传入两个整型值
	paramList:=[]reflect.Value{reflect.ValueOf(10),reflect.ValueOf(20)}

	//反射调用函数
	retList:=funcValue.Call(paramList)

	//获取第一个返回值，取整数值
	fmt.Println(retList[0].Int()) //30
}

