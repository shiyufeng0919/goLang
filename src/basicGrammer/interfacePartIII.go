package basicGrammer

import (
	"io"
	"fmt"
)

/*
第7章 接口
*/
/*
7.6接口的嵌套组合--将多个接口放在一个接口内

在GO语言中，不仅结构体与结构体间可以嵌套，接口与接口间也可以通过嵌套创造出新接口
*/
/*
1。系统包中的接口嵌套组合
GO语言io包中定义了写入器(Writer)，关闭器(Closer)，写入关闭器(WriteCloser)三个接口
*/

/*
2.在代码中使用接口嵌套组合
*/
//声明一个设备结构
type device struct { //该结构会实现io包中三个接口(Writer/Close/WriteCloser)
}
//实现io.Writer的Write()方法
func (d *device) Write(p []byte)(n int,err error){ //实现io.Writer的Write方法
	return 0,nil
}
//实现io.Closer的Close()方法
func (d *device) Close() error{ //实现io.Closer的Close()方法
	return nil
}
func InterfaceDemo10(){
	//声明写入关闭器，并赋予device的实例
	var wc io.WriteCloser=new(device) //device实例化，由于device实现了io.WriteCloser的所有嵌入接口，因此device指针会被隐式转换为io.WriteCloser接口
	//写入数据
	wc.Write(nil) //调用io.WriteCloser接口的Write方法,由于wc被赋值*device，因此最终会调用device的Write()方法
	//关闭设备
	wc.Close()
	//声明写入器，并赋予device的新实例
	var writeOnly io.Writer=new(device) //writeOnly是一个io.Writer接口，这个接口只有Write方法
	//写入数据
	writeOnly.Write(nil) //只能调用Write方法，没有Close方法
}

/*
7.7在接口和类型间转换

GO语言使用接口断言(typeassertions)将接口转换成另外一个接口，也可将接口转换为另外的类型
*/
/*
7.7.1类型断言的格式
t:=i.(T)

i:代表接口变量
T:代表转换的目标类型
t:代表转换后的变量

注：若i没有完全实现T接口的方法，则这个语句会触发宕机

更友好写法：
t,ok := i.(T)
ok可以被认为是：i接口是否实现T类型的结果
未实现时:t置为T类型的0值，ok置为false
*/

/*
7.7.2 将接口转换为其他接口
*/
//示例：鸟和猪
//定义飞行动物接口
type Flyer interface {
	Fly()
}
//定义行走动物接口
type Walker interface {
	Walk()
}
//定义鸟类
type Bird struct {

}
//实现飞行动物接口
func (b *Bird) Fly(){
	fmt.Println("bird:fly")
}
//为鸟添加walk()方法，实现行走动物接口
func (b *Bird) Walk(){
	fmt.Println("bird:walk")
}
//定义猪
type Pig struct {

}
//为猪添加Walk()方法，实现行走动物接口
func (p *Pig) Walk(){
	fmt.Println("pig:walk")
}
func InterfaceDemo11(){
	//创建动物的名字到实例的映射
	animals:=map[string]interface{}{ //map，映射对象名字和对象实例
		"bird":new(Bird),
		"pig":new(Pig),
	}
	//遍历映射
	for name,obj:=range animals{
		//判断对象是否为飞行动物
		f,isFlyer:=obj.(Flyer) //应用类型断言获取f
		//判断对象是否为行走动物
		w,isWalker:=obj.(Walker)
		fmt.Printf("name:%s isFlayer:%v isWalker:%v\n",name,isFlyer,isWalker)

		//如果飞行动物，则调飞行动物接口
		if isFlyer{
			f.Fly()
		}
		//如果行走动物，则调行走动物接口
		if isWalker{
			w.Walk()
		}
	}
	/*执行结果
		name:pig isFlayer:false isWalker:true
		pig:walk
		name:bird isFlayer:true isWalker:true
		bird:fly
		bird:walk
	*/
}

/*
7.7.3将接口转换为其他类型
*/
//上述代码：实现将接口转换为普通的指针类型
func InterfaceDemo12(){
	p1:=new(Pig)
	var a Walker=p1 //由于pig实现了Walker接口，因此可隐式转换为Walker接口类型保存于a中
	p2:=a.(*Pig) //由于a中保存的本身就是*Pig本体，因此可转换为*Pig类型
	//p2:=a.(*Bird) //报错：接口转换时，Walker接口内部保存的是*Pig，而不是*Bird。因此接口在转换为其他类型时，接口内保存的实例对应的类型指针必须是要转换的对应的类型指针
	fmt.Printf("p1=%p p2=%p",p1,p2) //p1和p2的指针是相同的
	/*
	注：接口断言类似于流程控制中的if。但大量类型断言出现时，应使用更为高效的switch分支
	*/
}


/*
7.8 空接口类型（interface{}） 能保存所有值的类型

空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口

空接口类型可以保存任何值，也可从空接口中取出原值

类似于java中的Object

注：空接口内部实现保存了对象的类型和指针，使用空接口保存一个数据过程会比直接用数据对应类型的变量保存要慢。
*/
/*
7.8.1 将值保存到空接口
*/
func InterfaceDemo13(){
	//声明any为interface{}类型变量
	var any interface{}
	any=1
	fmt.Println(any) //1

	any="interface"
	fmt.Println(any) //interface

	any=false
	fmt.Println(any) //false
}
/*
7.8.2 从空接口获取值
*/
func InterfaceDemo14(){
	//声明a变量，类型为int，初始值为1
	var a int =1
	//声明i变量，类型为interface{}，初始值为a，此时i的值变为1
	var i interface{} =a //i虽然赋值为int，但i类型依然为interface{}
	//声明b变量，尝试赋值i
	//var b int=i //编译错误,不能将i变量视为int类型赋值给b
	var b int=i.(int) //应用类型断言
	fmt.Println(b) //1
}
/*
7.8.3 空接口的值比较(==)
*/
/*
1。类型不同的空接口间的比较结果不相同
*/
func CompareInterface1(){
	//a保存整型
	var a interface{} =100
	//b保存字符串
	var b interface{} = "compare"
	//两个空接口不相等
	fmt.Println(a==b) //false
}
/*
2.不能比较空接口中的动态值
*/
func CompareInterface2(){
	//c保存包含10的整型切片
	var c interface{} =[]int{10}
	//d保存包含20的整型切片
	var d interface{} = []int{20}
	fmt.Println(c==d) //崩溃,运行时错误，[]int是不可比较的类型
}
/*
类型及可比较情况：
类型 				说明
map					宕机错误，不可比较
切片([]T)			宕机错误，不可比较
通道(channel)		可比较，必须由同一make生成，即同一个通道才会是channel，否则为false
数组([容量]T)			可比较，编译期知道两个数组是否一致
结构体				可比较，可逐个比较结构体的值
函数					可比较
*/



/*
7.9 示例：使用空接口实现可以保存任意值的字典

字典同java中的map类似,可将任意类型的值做成键值对保存，然后进行找回，遍历操作
*/
/*
1。值设置和获取
*/
//字典结构
type Dictionary struct {
	data map[interface{}]interface{} //键值均为interface{}类型
}
//根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key] //若key不存在，则返回nil
}
//设置键值
func (d *Dictionary) Set(key interface{},value interface{}) {
	d.data[key]=value
}

/*
2.遍历字段的所有键值关联数据
*/
//遍历所有的键值，若回调返回值为false，则停止遍历
func (d *Dictionary) Visit(callback func(k,v interface{}) bool) { //回调函数类型 func(k,v interface{}) bool
	if callback == nil{
		return
	}
	for k,v:=range d.data{ //遍历map中的所有元素
		if !callback(k,v){ //根据callback返回值决定是否遍历
			return
		}
	}
}
/*
3.初始化和清除
*/
//清空所有数据
func (d *Dictionary) Clear(){
	d.data=make(map[interface{}]interface{})
}
//创建一个字典
func NewDictionary() *Dictionary{
	d:=&Dictionary{} //实例化一个Dictionary,实例化对象为d
	//初始化map
	d.Clear()
	return d
}

/*
4.使用字典
*/
func DictionaryDemo1(){
	//创建字典实例
	dict:=NewDictionary()
	//添加游戏数据
	dict.Set("A",60)
	dict.Set("B", 36)
	dict.Set("C",24)
	//获取值及打印值
	favorite:=dict.Get("B")
	fmt.Println(favorite) //36
	//遍历所有的字典元素
	dict.Visit(func(key, value interface{}) bool {
		//将值转为int类型，并判断是否大于40
		if value.(int) >40{
			//输出
			fmt.Println(key,"is expensive")
			return true
		}
		fmt.Println(key,"is cheap")
		return true
	})
}

/*
7.10 类型分支--批量判断空接口中变量的类型

GO语言的switch不仅可以像其他语言一样实现数值，字符串的判断，还可以：
判断一个接口内保存或实现的类型
*/
/*
7.10.1 类型断言的书写格式

switch 接口变量.(type){
  case 类型1:
     //变量是类型1时的处理
  case 类型2:
    //变量是类型2的处理
...
default:
   //变量不是所有case中列举的类型时的处理
}
*/
/*
7.10.2 使用类型分支判断基本类型
*/
func PrintType( v interface{}){
	switch v.(type) { //类型分支典型写法
	case int:
		fmt.Println(v,"is int")
	case string:
		fmt.Println(v,"is string")
	case bool:
		fmt.Println(v,"is bool")
	}
}
func DemoPrintType(){
	PrintType(123)
	PrintType("hello")
	PrintType(true)
}
/*
7.10.3使用类型分支判断接口类型
*/
//电子支付方式
type Alipay struct {
}

//为Alipay添加CanUseFaceId方法，能够刷脸
func (a *Alipay) CanUseFaceId(){
}

//现金支付方式
type Cash struct {
}
//为Cash添加Stolen方法，容易偷窃
func (c *Cash) Stolen(){
}

//具备刷脸特性的接口
type ContainCanUseFaceId interface {
	CanUseFaceId()
}
//具备被偷特性的接口
type ContainStolen interface {
	Stolen()
}
//打印支付方式具备的特点
func Print(payMethod interface{}){
	switch payMethod.(type){
	case ContainCanUseFaceId:
		fmt.Printf("%T canUseFaceId\n",payMethod)
	case ContainStolen:
		fmt.Printf("%T maybe stolen",payMethod)
	}
}
func DemoPrintType2(){
	Print(new(Alipay)) //*basicGrammer.Alipay canUseFaceId
	Print(new(Cash)) //*basicGrammer.Cash maybe stolen
}


/*
7.11 示例：实现有限状态机(FSM)

有限状态机(Finite-State Machine,FSM)：
表示有限个状态及在这些状态间的转移和动作等行为的数学模型
*/