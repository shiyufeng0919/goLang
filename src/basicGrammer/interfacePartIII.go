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
*/

