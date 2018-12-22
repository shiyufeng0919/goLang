package basicGrammer

import (
	"fmt"
	"encoding/json"
)

/*
6.6 类型内嵌和结构体内嵌 & 分离json数据

结构体允许其成员字段在声明时没有字段名而只有类型。这种形式字段被称为类型内嵌或匿名字段
*/
type data struct {
	int //定义结构体的匿名字段.类型内嵌依然有自己的字段名，只是字段名即为其类型本身（同种类型字段只能有一个）
	float32
	bool
}
func structDemo11(){
	/*
	结构体实例化后，若匿名字段类型为结构体，则可直接访问匿名结构体里的所有成员。此种方式称结构体内嵌
	*/
	ins:=&data{ //实例化data。并为data赋值
	 int:10,
	 float32:3.14,
	 bool:true,
	}
	fmt.Println(ins)
}

/*
6.6.1 声明结构体内嵌
*/
//基础颜色
type BasicColor struct {
	//三种颜色分量
	R,G,B float32
}
//完整颜色定义
type Color struct {
	//将基本颜色作为成员
	Basic BasicColor
	//透明度
	Alpha float32
}
func StructDemo12(){
	var c Color
	//设置基本颜色分量
	c.Basic.R=1
	c.Basic.G=1
	c.Basic.B=0
	//设置透明度
	c.Alpha=1 //1不透明 0完全透明
	//显示整个结构体内容
	fmt.Printf("%+v",c)//{Basic:{R:1 G:1 B:0} Alpha:1}
}

//优化上述代码：应用结构体内嵌写法
type basicColor struct {
	R,G,B float32
}
type color struct {
	basicColor //结构体内嵌(basicColor结构体嵌入到color结构体，没有字段名，只有类型。叫结构体内嵌)
	alpha float32
}
func StructDemo13(){
	var c color
	c.R=1
	c.G=1
	c.B=0
	c.alpha=1
	fmt.Printf("%+v",c)//{basicColor:{R:1 G:1 B:0} alpha:1}
}

/*
6.6.2结构体内嵌特性

1.内嵌的结构体可以直接访问其成员变量

2。内嵌结构体的字段名是它的类型名
注：一个结构体只能嵌入一个同类型的成员
*/

/*
6.3.3使用组合思想描述对象特性

在面向对象思想中，实现对象关系需要使用"继承"特性（禁多重继承）

GO语言结构体内嵌特性是一种组合特性，可快速构建对象的不同特性
*/
//可飞行的
type flying struct {

}
func (f flying) fly(){ //为结构体flying添加方法fly()
	fmt.Println("can fly")
}
//可行走的
type walkable struct {

}
func (w walkable) walk(){ //为结构体walkable添加方法walk()

}
//人类
type human struct {
	walkable //内嵌结构体
}
//鸟类
type bird struct {
	walkable //内嵌结构体
	flying
}
func StructDemo14(){
	//实例化鸟类
	b:=new(bird)
	fmt.Println("bird:")
	b.fly() //调用鸟类能使用的功能
	b.walk()

	//实例化人类
	h:=new(human)
	fmt.Println("human:")
	h.walk() //调用人类能使用的功能
}

/*
6.6.4初始化结构体内嵌
*/
//车轮
type Wheel struct {
	Size int
}
type Engine struct {
	Power int  //功率
	Type string //类型
}
//车
type Car struct {
	Wheel
	Engine
}
func StructDemo15(){
	c:=Car{
		Wheel:Wheel{
			Size:18,
		},
		Engine:Engine{
			Type:"1.4T",
			Power:143,
		},
	}
	fmt.Printf("%+v\n",c) //{Wheel:{Size:18} Engine:{Power:143 Type:1.4T}}
}

/*
6.5.5 初始化内嵌匿名结构体
*/
//车轮
type wheel struct {
	size int
}
//车
type car struct {
	wheel
	//引擎
	engine struct{ //被直接定义在car结构体内部。这种嵌入写法即将原来结构体类型转换为struct{...}
		Power int
		Type string
	}
}
func StructDemo16(){
	c:=car{
		wheel:wheel{ //初始化轮子
			size:18,
		},
		engine: struct { //初始化引擎(由于engine字段类型未单独定义，此初始化时需要填写struct{...}声明其类型)
			Power int
			Type  string
		}{Power: 143, Type:"1.4T"},
	}
	fmt.Printf("%+v\n",c)//{wheel:{size:18} engine:{Power:143 Type:1.4T}}
}


/*
6.6.6成员名字冲突
*/
type a struct {
	a int
}
type b struct {
	a int //与a结构体字段同名
}
type c struct {
	a
	b
}
func StructDemo17(){
	c:=&c{}
	c.a.a=1 //可正常输出
	//c.a=1 //报错，编译器不知赋给a中的a还是b中的a
	fmt.Println(c) //&{{1} {0}}
}

/*
6.7 示例：使用匿名结构体分离JSON数据
*/
/*
1。定义数据结构
注意必须大写，否则Marshal的数据为{}
*/
//定义手机屏幕
type Screen struct {
	Size float32 //屏幕尺寸
	Resx,Resy int //屏幕小平和垂直分辨率
}
//定义电池
type Battery struct {
	Capcacity int //容量
}
//准备json数据
func genJsonData() []byte{
	//完整数据结构
	raw:=&struct {
		Screen
		Battery
		HashTouchId bool //序列化时添加的字段，是否有指纹识别
	}{
		//屏幕参数
		Screen:Screen{
			Size:5.5,
			Resx:1920,
			Resy:1080,
		},
		//电池参数
		Battery:Battery{
			2910,
		},
		//是否有指纹识别
		HashTouchId:true,
	}
	//将数据序列化为json
	jsonData,err:=json.Marshal(raw)
	if err !=nil{
		return nil
	}
	fmt.Println("jsonData:",string(jsonData))//jsonData: {"Size":5.5,"Resx":1920,"Resy":1080,"Capcacity":2910,"HashTouchId":true}
	return jsonData
}
//分离json数据
func StructDemo18(){
	//生成一段json数据
	jsonData :=genJsonData()
	fmt.Println(string(jsonData))//{"Size":5.5,"Resx":1920,"Resy":1080,"Capcacity":2910,"HashTouchId":true}
	//只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch:= struct {
		Screen
		HasTouchId bool
	}{}
	//反序列化到screenAndTouch
	json.Unmarshal(jsonData,&screenAndTouch)
	//只需要电池和指纹识别信息的结构和实例
	batteryAndTouch:= struct {
		Battery
		HasTouchId bool
	}{}
	//反序列化到batteryAndTouch
	json.Unmarshal(jsonData,batteryAndTouch)
	//输出batteryAndTouch详细信息
	fmt.Printf("%+v\n",batteryAndTouch)//{Battery:{Capcacity:0} HasTouchId:false}

}