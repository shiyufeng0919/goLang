package basicGrammer

import (
	"fmt"
	"math"
	"net/http"
	"strings"
	"os"
	"io/ioutil"
	"time"
)

/*
6.5 方法
*/
/*
6.5.1 为结构体添加方法

GO语言中的方法（Method）是一种作用于特定类型变量的函数。
这种特定类型变量叫做接收器（Receiver）

若将特定类型理解为结构体/类时，接收器概念就类似于this/self
*/
/*
1.面向过程实现方法
面向过程中没有"方法"概念，只能通过结构体和函数，由使用者使用函数参数和调用关系来形成接近"方法的概念
*/
type Bag struct {
	items []int
}
//将一个特品放入背包的过程
func insertBag(b *Bag,itemid int){
	b.items=append(b.items,itemid) //将物品放入背包的过程作为"方法"
}
func methodDemo1(){
	bag:=new(Bag) //实例化结构体Bag,bag为实例
	insertBag(bag,1001)
}
/*
2.GO语言的结构体方法
*/
/*
为*Bag创建一个方法
(b *Bag)表示接收器，即insertBag作用的对象实例
注：每个方法只能有一个接收器
*/
func (b *Bag) insertBag(itemid int){
	b.items=append(b.items,itemid)
}
func methodDemo2(){
	bag:=new(Bag)
	bag.insertBag(1001)
}

/*
6.5.2 接收器--方法作用的目标
*/
/*
接收器的格式：
func (接收器变量 接收器类型) 方法名(参数列表)(返回参数){
   函数体
 }
命名规则：
接收器变量名为接收器类型名第一个小写字母。如Socket类型接收器变量应该命名为s

接收器类型和参数类似，可为指针/非指针类型

接收器根据接收器类型可分为指针接收器和非指针接收器
*/
/*
1。理解指针类型的接收器
*/
type Property struct {
	//属性值
	value int
}
//设置属性值
func (p *Property) setValue(v int){
	//修改p的成员变量
	p.value=v
}
//取属性值
func (p *Property) getValue() int{
	return p.value
}
func methodDemo3(){
	p:=new(Property) //实例化
	p.setValue(1001) //设置值
	fmt.Println(p.getValue()) //取值
}
/*
2。理解非指针类型的接收器

当方法作用于非指针接收器时，GO语言会在代码运行时将接收器的值复制一份。
在非指针接收器的方法中可以获取接收器成员的值，但修改无效。
*/
type Point struct { //定义点结构
	x int
	y int
}
//非指针接收器的加方法
//使用了非指针接收器，add()方法变得类似于只读方法，add()方法内部不会对成员进行任何修改
func (p Point) add(other Point) Point{
	//成员值与参数相加后返回新的结构
	return Point{p.x+other.x,p.y+other.y}
}
func methodDemo4(){
	//初始化点
	p1:=Point{1,1}
	p2:=Point{2,2}

	//与另外一个点相加
	result:=p1.add(p2)
	//输出结果
	fmt.Println(result)
}

/*
3。指针和非指针接收器的使用

小对象由于值复制速度较快，适合使用非指针接收器
大对象由于复制性能较低，适合使用指针接收器，在接收器和参数间传递不进行复制，只是传递指针
*/


/*
6.5.3 示例：二维矢量模拟玩家移动
*/
//1.实现二维矢量结构
type Vec2 struct {
	X,Y float32
}
//使用矢量加上另外一个矢量，生成新的矢量
func (v Vec2) add(other Vec2) Vec2{
	return Vec2{ //返回相加后结果，不会修改Vec2成员值
		v.X+other.X,
		v.Y+other.Y,
	}
}
//使用矢量减去另外一个矢量，生成新的矢量
func (v Vec2) sub(other Vec2) Vec2{
	return Vec2{
		v.X-other.X,
		v.Y-other.Y,
	}
}
//使用矢量乘以另外一个矢量，生成新的矢量
func (v Vec2) scale(s float32) Vec2{ //等比缩放
	return Vec2{
		v.X*s,
		v.Y*s,
	}
}
//计算两个矢量的距离
func (v Vec2) distanceTo(other Vec2) float32{
	dx:=v.X-other.X
	dy:=v.Y-other.Y
	return float32(math.Sqrt(float64(dx*dx+dy*dy)))
}
//返回当前矢量的标准化矢量
func (v Vec2) normalize() Vec2{
	mag:=v.X*v.X+v.Y*v.Y
	if mag>0{
		oneOverMag:=1/float32(math.Sqrt(float64(mag)))
		return Vec2{v.X*oneOverMag,v.Y*oneOverMag}
	}
	return Vec2{0,0}
}

//2.实现玩家对象
type Player struct {
	currPos Vec2 //当前位置
	targetpos Vec2 //目标位置
	speed float32 //移动速度
}
//设置玩家移动的目标位置
func (p *Player) moveTo(v Vec2)  {
	p.targetpos=v
}
//获取当前位置
func (p *Player) pos() Vec2{
	return p.currPos
}
//判断是否到达目的地
func (p *Player) isArrived() bool{
	//通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.distanceTo(p.targetpos) < p.speed
}
//更新玩家位置
func (p *Player) update(){
	if !p.isArrived(){
		//计算出当前位置指向目标的朝向
		dir:=p.targetpos.sub(p.currPos).normalize()
		//添加速度矢量生成新的位置
		newPos:=p.currPos.add(dir.scale(p.speed))
		//移动完成后，更新当前位置
		p.currPos=newPos
	}
}
//创建新玩家
func newPlayer(speed float32) *Player{
	return &Player{
		speed:speed,
	}
}

//3.处理移动逻辑
func methodDemo5(){
	//实例化玩家对象，并设速度为0.5
	p:=newPlayer(0.5)
	//让玩家移动到3,1点
	p.moveTo(Vec2{3,1})
	//如果没有到达就一直循环
	for !p.isArrived(){
		//更新玩家位置
		p.update()
		//打印每次移动后的玩家位置
		fmt.Println(p.pos())
	}
}

/*
6.5.4为类型添加方法

GO语言可为任何类型添加方法
*/
/*
1。基本类型添加方法
*/
//将int定义为myInt类型
type myInt int
//为myInt类型添加方法isZero()
func (m myInt) isZero() bool{
	return m==0
}
//为myInt添加add()方法 (非指针接收器)
func (m myInt) add(other int) int{
	return other+int(m)
}
func methodDemo6(){
	var b myInt
	fmt.Println(b.isZero()) //true
	b=1
	fmt.Println(b.add(2)) //3
}

/*
2.http包中的类型方法
*/
func methodDemo7(){
	//实例化http客户端,请求需要通过该client实例发送
	client:=&http.Client{}
	//创建一个http请求 strings.NewReader创建一个字符串读取器
	//仅构造一个请求对象，不会连接网络
	req,err:=http.NewRequest("POST","http://www.163.com",strings.NewReader("key=value"))
   //发现错误打印并退出
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//为标头添加信息
	req.Header.Add("User-Agent","my Client")
	//开始请求
	resp,err:=client.Do(req)//client将http请求发送到服务器,服务器响应后将信息返回并保存到resp中
	//处理请求错误
	if err !=nil{
		fmt.Println(err)
		os.Exit(1)
		return
	}
	//读取服务器返回的内容
	data,err:=ioutil.ReadAll(resp.Body)//读取响应的body部分并打印
	fmt.Println(string(data))
	defer resp.Body.Close()
}

/*
3.time包中的类型方法

time包用于时间的获取和计算
*/
func methodDemo8(){
	//time.Second是一个常量
	//Second的类型是Duration,Duration实际是一个int64类型
	//Duration.String可将Duration的值转为字符串
	fmt.Println(time.Second.String())
}


/*
6.5.5 示例：使用事件系统实现事件的响应和处理
*/
/*
1.方法和函数的统一调用

下述例子：结构体和普通函数参数/签名完全一致
然后用与它们签名一致的函数变量分别赋值方法与函数
则该函数变量可以保存普通函数/结构体方法
*/
//声明一个结构体
type class struct {
}
//给结构体添加Do方法
func (c *class) do(v int){
	fmt.Println("call method do:",v)
}
//普通函数的do()方法
func funcDo(v int){
	fmt.Println("call function do:",v)
}
func methodDemo9(){
	//声明一个函数回调
	var delegate func(int)
	//创建结构体实例
	c:=new(class)
	//将回调设为c的do方法
	delegate=c.do
	//调用
	delegate(100) //call method do:100
	//将回调设为普通函数
	delegate=funcDo
	delegate(100) //call function do:100
}
/*
2.事件系统基本原理
事件系统可将事件派发者与事件处理者解耦
事件调用      事件响应代码
      <--注册---
      --调用---->
*/

/*
3.事件注册

事件注册的过程就是将事件名称和响应函数关联并保存起来
*/
//实例化一个通过字符串映射函数切片的map
var eventByName=make(map[string][]func(interface{}))//通过事件名string关联回调列表func(interface{})
//注册事件，提供事件名和回调函数
func RegisterEvent(name string,callback func(interface{})){
	//通过名字查找事件列表
	list:=eventByName[name]
	//在列表切片中添加函数
	list=append(list,callback)
	//保存修改的事件列表切片
	eventByName[name]=list
}

/*
4.事件调用

事件调用方和注册方是事件处理中完全不同的两个角色
*/
//调用事件  name:提供事件名字，param：事件参数表示描述事件具体的细节
func callEvent(name string,param interface{}){
	//通过名字找到事件列表
	list:=eventByName[name]
	//遍历这个事件的所有回调
	for _,callback:=range list{
		//传入参数调用回调
		callback(param)//触发事件实现方的逻辑处理
	}
}


/*
5.使用事件系统
将事发现场和事件处理方联系起来
*/
//声明角色的结构体
type Actor struct {

}
//为角色添加一个事件处理函数
func (a *Actor) onEvent(param interface{}){
	fmt.Println("actor event:",param)
}

//全局事件
func globalEvent(param interface{}){
	fmt.Println("global event:",param)
}
func methodDemo10(){
	//实例化一个角色
	a:=new(Actor)
	//注册名为onSkill回调
	RegisterEvent("onSkill",a.onEvent)
	//再次在onSkill上注册全局事件
	RegisterEvent("onSkill",globalEvent)
	//调用事件，所有注册的同名函数都会被调用
	callEvent("onSkill",100)
	/*
	执行结果：（角色和全局的事件会按注册顺序顺序地触发，事件系统认为所有函数都是平等的）
	actor event: 100
	global event: 100
	*/
}

func StructPartII(){
	methodDemo10()
}

