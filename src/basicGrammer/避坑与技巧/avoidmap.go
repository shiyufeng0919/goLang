package 避坑与技巧

import "fmt"

/*
第12章 "避坑"与技巧

12.4 map的多键索引--多个数值条件可以同时查询
*/
//人员档案
type Profile struct {
	Name string //名字
	Age int //年龄
	Married bool //已婚
}
func Demo7(){
	list:=[]*Profile{
		{Name:"A",Age:30,Married:true},
		{Name:"B",Age:21},
		{Name:"C",Age:22},
		{Name:"A",Age:30,Married:false},
	}
	buildIndex(list)
	queryData("A",30)
}
//12.4.1 基于哈希值的多键索引及查询
//1.字符串转哈希值
func simpleHash(str string)(ret int){
	//遍历字符串中的每一个ASCII字符
	for i:=0;i<len(str);i++{
		//取出字符
		c:=str[i] //c变量类型为uint8
		//将字符的ASCII码相加
		ret+=int(c)
	}
	return
}
//2.查询键
type classicQueryKey struct {
	Name string //要查询的名字
	Age int //要查询的年龄
}
//计算查询键的哈希值
func (c *classicQueryKey) hash()int{
	//将名字的hash和年龄hash合并
	return simpleHash(c.Name)+c.Age*1000000
}
/*
3.构建索引

本例需快速查询，因此需要提前对已有数据构建索引
*/
//创建哈希值到数据的索引关系
var mapper=make(map[int][]*Profile) //key:hash值 value:*Profile
//构建数据索引
func buildIndex(list []*Profile){
	//遍历所有的数据
	for _,profile:=range list{
		//构建数据的查询索引
		key:=classicQueryKey{profile.Name,profile.Age}
		//计算数据的哈希值，取出已经存在的记录
		existValue:=mapper[key.hash()]
		//将当前数据添加到已经存在的记录切片中
		existValue=append(existValue,profile)
		//将切片重新设置到映射中
		mapper[key.hash()]=existValue
	}
}
/*
4.查询逻辑
*/
func queryData(name string,age int){
	//根据给定查询条件构建查询键
	keyToQuery:=classicQueryKey{name,age}
	//计算查询键的哈希值并查询，获得相同哈希值的所有结果集合
	resultList:=mapper[keyToQuery.hash()]
	fmt.Println(resultList)
	//遍历结果集合
	for _,result:=range resultList{
		//与查询结果比对，确认找到打印结果
		if result.Name==name && result.Age==age{
			fmt.Println(result)
			return
		}
	}
	//没有查询到，打印结果
	fmt.Println("no found")
}


/*
12.4.2 利用map特性的多键索引及查询

使用结构体进行多键索引和查询比传统的写法更为简单，最主要区别是无须准备哈希函数及相应的字段无须做哈希合并
*/
/*
1.构建索引
*/
type Prof struct {
	Name string //名字
	Age int //年龄
	Married bool //已婚
}
//查询键
type queryKey struct{
	Name string
	Age int
}
//创建查询键到数据映射
var mapp=make(map[queryKey]*Prof)
//构建查询索引
func buildindex(list []*Prof){
	//遍历所有数据
	for _,profile:=range list{
		//构建查询键
		key:=queryKey{
			Name:profile.Name,
			Age:profile.Age,
		}
		//保存查询键
		mapp[key]=profile
	}
}
//2.查询逻辑
func querydata(name string,age int){
	//根据查询条件构建查询键
	key:=queryKey{name,age}
	//根据键值查询数据
	result,ok:=mapp[key]
	//找到数据打印出来
	if ok{
		fmt.Println(result)
	}else{
		fmt.Println("no found")
	}
}
func Demo8(){
	list:=[]*Prof{
		{Name:"A",Age:30,Married:true},
		{Name:"B",Age:21},
		{Name:"C",Age:22},
		{Name:"A",Age:30,Married:false},
	}
	buildindex(list)
	querydata("A",30)
}

/*
以上注意：代码量减少
GO语言的底层会为map键自动构建哈希值。能够构建哈希值的类型必须是非动态类型，非指针，函数，闭包
1。非动态类型：可用数组，不能用切片
2。非指针：每个指针数值都不同，失去哈希意义
3。函数，闭包不能作为map的键
*/