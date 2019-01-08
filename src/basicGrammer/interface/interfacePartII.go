package _interface

import (
	"sort"
	"fmt"
)

/*
第7章 接口
*/
/*
7.5示例：使用接口进行数据的排序

Go语言在排序时，需要使用者通过sort.Interface接口提供数据的一些特性和操作方法
*/
type Interface interface {
	//获取无素的数量
	Len() int
	//小于比较
	Less(i,j int) bool
	//交换元素
	Swap(i,j int)
}
/*
7.5.1使用sort.Interface接口进行排序
*/
//将[]string定义为MyStringList类型

type MyStringList []string
//实现sort.Interface接口的获取元素的数量
func (m MyStringList) Len() int{
	return len(m)
}
//实现sort.Interface接口的比较元素方法
func (m MyStringList) Less(i,j int) bool{
	return m[i]<m[j]
}
//实现sort.Interface接口的交换元素方法
func (m MyStringList) Swap(i,j int){
	m[i],m[j]=m[j],m[i]
}
func InterfaceDemo5(){
	//准备一个内容被打乱的字符串切片
	names:=MyStringList{
		"red",
		"color",
		"yellow",
	}
	//使用sort包进行排序
	sort.Sort(names)
	//遍历打印结果
	for _,v:=range names{
		fmt.Printf("%s\n",v)
	}
}

/*
7.5.2常见类型的便捷排序
*/
/*
1.字符串切片的便捷排序 sort包有个StringSLice类型
*/
func InterfaceDemo6(){
	//字符串切片的便捷排序
	names:=sort.StringSlice{
		"red",
		"color",
		"yellow",
	}
	sort.Sort(names)
	fmt.Println(names) //[1 2 3 4 5 A B C D E]
}
/*
2.对整型切片进行排序
*/
func InterfaceDemo7(){
	names:=[]string{
		"3",
		"2",
		"1",
		"4",
	}
	sort.Strings(names)//直接对字符串切片进行排序
	for _,v:=range names{
		fmt.Printf("%s\n",v)
	}
}
/*
sort包内建的类型排序接口一览
   类型             实现sort.Interface的类型      直接排序方法                      说明
字符串(string)      StringSlice                    sort.Strings(a []string)     字符ascii码升序
整型(int)		   IntSlice						  sort.Ints(a []int)           数值升序
双精度浮点(float64)  float64Slice                   sort.Float64s(a []float64)   数值升序
*/


/*
7.5.3对结构体数据进行排序
*/
/*
1.完整实现sort.Interface进行结构体排序
*/
/*
示例：选按kind分类，kind一致，再按name分类
*/
//声明英雄的分类
type HeroKind int
//定义HeroKind常量，类似于枚举
const (
	NoneHeroKind=iota
	Tank
	Assassin
	Mage
)
//定义英雄名单的结构
type Hero struct {
	Name string //英雄名字
	Kind HeroKind //英雄分类
}
//将英雄指针的切片定义为Heros类型
type Heros []*Hero
//实现sort.Interface接口取元素数量方法
func (s Heros) Len() int{
	return len(s)
}
//实现sort.Interface接口比较元素方法
func (s Heros) Less(i,j int) bool{
	//若英雄分类不一致，则优先对分类进行排序
	if s[i].Kind!=s[j].Kind{
		return s[i].Kind<s[j].Kind
	}
	//默认按英雄名字字符升序排列
	return s[i].Name<s[j].Name
}
//实现sort.Interface接口交换元素方法
func (s Heros) Swap(i,j int){
	s[i],s[j]=s[j],s[i]
}
func InterfaceDemo8(){
	heros:=Heros{
		&Hero{"吕布",Tank},
		&Hero{"李白",Assassin},
		&Hero{"妲已",Mage},
		&Hero{"关羽",Tank},
		&Hero{"诸葛亮",Mage},
	}
	//使用sort包进行排序
	sort.Sort(heros)
	for _,v:=range heros{
		fmt.Printf("%+v\n",v)
	}
}

/*
2.使用sort.slice进行切片元素排序
*/
type Herokind int
const(
	None1=iota
	Tank1
	Assassin1
	Mage1
)
type Hero1 struct {
	Name string
	Kind Herokind
}
func InterfaceDemo9(){
	heros:=[] *Hero1{
		{"吕布",Tank},
		{"李白",Assassin},
		{"妲已",Mage},
		{"关羽",Tank},
		{"诸葛亮",Mage},
	}
	sort.Slice(heros,func(i,j int)bool{
		if heros[i].Kind!=heros[j].Kind{
			return heros[i].Kind<heros[j].Kind
		}
		return heros[i].Name<heros[j].Name
	})
	for _,v:=range heros{
		fmt.Printf("%+v\n",v)
	}
}