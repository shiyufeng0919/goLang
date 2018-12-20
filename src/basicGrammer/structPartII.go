package basicGrammer

import (
	"fmt"
	"math"
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