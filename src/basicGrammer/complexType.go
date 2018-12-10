package basicGrammer

import (
	"fmt"
	"sort"
	"sync"
	"container/list"
)

/*
复杂类型具有各种形式的存储和处理数据的功能，将它们称为"容器"
容器以标准库的方式提供
以下介绍：数组，切片，映射及列表的增，删，改，遍历使用方法
Go的数组和切片从C语言延续过来。
*/

//1.数组--固定大小的连续空间
func ArrayType(){
	//A.声明数组： var 数组变量名 [元素数量]T
	var team [3]string
	team[0]="A"
	team[1]="B"
	team[2]="C"
	fmt.Println(team)//[A B C]

	//B.初始化数组
	var team2=[3]string{"a","b","c"} //指定数组大小
	var team3=[...]string{"a1","b1","c1"} //不指定数组大小（...表示让编译器确定数组大小）
	fmt.Println(team2,team3)//[a b c] [a1 b1 c1]

	//C.遍历数组--访问每一个数组元素
	for k,v:=range team{ //k为数组索引，v为数组的每个元素
		fmt.Println(k,v) //0 A  - 1 B - 2 C
	}
}

//2.切片（slice）--动态分配大小的连续空间
func SliceType(){
	/*
	GO语言切片的内部结构包含：地址（从哪里开始切），大小（切多大），容量（装切片口袋大小）
	切片用于快速操作一块数据集合
	*/

	//A。从数组或切片生成新的切片
	//切片默认指向一段连续内存区域，可以是数组，也可以是切片本身（从连续内存区域生成切片是常见操作）
	//slice [开始位置:结束位置]
	var array = [3]int{1,2,3}
	arraySlice:=array[1:2]  //从array切片，[第1个索引-第2个索引）取出的元素数量=结束位置-开始位置
	fmt.Println(arraySlice)//[2] 表示从索引为1开始到索引为2结束位置（不包括2）
	fmt.Println(array[:2]) //[1 2]  表示从开始位置到索引为2位置
	fmt.Println(array[1:]) //[2 3]  表示从索引为1位置到结束位置
	fmt.Println(array[1:len(array)]) //[2 3] 结速位置可用len(array)，因为不包含，所以不会越界
	fmt.Println(array[len(array)-1]) //3 切片最后一个元素获取
	fmt.Println(array[:]) //[1 2 3]  与切片本身等效
	fmt.Println(array[0:0]) //[] 空切片，一般用于切片复位

	//B。声明切片 var name []T  #name:切片类型的变量名  T:切片类型对应的元素类型
	var strList []string //声明字符串切片
	var numList []int //声明整型切片
	var numListEmpty = []int{} //声明一个空切片
	fmt.Println(strList,numList,numListEmpty)//输出3个切片 [][][]
	fmt.Println(len(strList),len(numList),len(numListEmpty))//输出3个切片大小 0 0 0
	//切片判定空的结果（切片是动态结构，只能与nil判断相等，不能互相判定相等）
	fmt.Println(strList==nil) //true
	fmt.Println(numList==nil) //true
	fmt.Println(numListEmpty==nil) //false numListEmpty已经被分配到了内存，但没有元素

	//C.使用make()函数构造切片 make([]T,size,cap) #T切片的元素类型 size为这个类型分配多少个元素,cap预分配的元素数量，此值不影响size,只是能提前分配空间，降低多次分配空间造成的性能问题
	slice1:=make([]int,2)
	slice2:=make([]int,2,10)//预分配10个空间，但实际仅用了2个
	fmt.Println(slice1,slice2) //[0,0] [0,0]
	fmt.Println(len(slice1),len(slice2))//2,2（容量cap不影响当前元素个数）

	//D.使用append()函数为切片添加元素
	/*
	注：每个切片都会指向一片内存空间，这片空间能容纳一定数量的元素。当空间不足时，切片就会"扩容"
	扩容操作发生在append()函数调用时。空量扩展规律以2的倍数扩容，即1,2,4,8...
	类比：
	员工和工位即切片中的元素
	办公地址就是分配好的内存
	搬家就是重新分配内存
	无论搬多少次家，公司名称不会变，即变量名不会变
	*/
	var numbers []int //声明一个整型切片
	for i:=0;i<10;i++{
		numbers=append(numbers,i)
		fmt.Printf("len:%d cap:%d pointer:%p\n",len(numbers),cap(numbers),numbers)//输出元素个数，切片容量，指针变化
	}

	//append()一次性添加很多元素
	var car []string //声明一个字符串切片
	car =append(car,"red") //添加一个元素
	car=append(car,"blue","yellow","black")//添加多个元素
	//添加切片
	team:=[]string{"A","B","C"} //声明新切片
	car=append(car,team...) //***注：...表示将team整个添加到car的后面
	fmt.Println(car)//[red blue yellow black A B C]

	//E.复制切片元素到另一个切片copy()
	//copy(destSlice,srcSlice []T) int #srcSlice为数据来源切片，destSlice为复制的目标,返回值表示实际发生复制的元素个数
	const elementCount=1000 //设置元素数量为1000
	srcData:=make([]int,elementCount)//预分配足够多的元素切片
	//将切片赋值
	for i:=0;i<elementCount;i++{
		srcData[i]=i;//[0,1,2,3...]
	}
	//引用切片数据
	refData:=srcData
	//预分配足够多的元素切片
	copyData:=make([]int,elementCount)
	//将数据复制到新的元素切片空间中
	copy(copyData,srcData)
	//修改原始数据的第一个值
	srcData[0]=999
	//打印引用切片的第一个元素
	fmt.Println("refData[0]:",refData[0]) //refData[0]: 999
	//打印复制切片的第一个和最后一个元素
	fmt.Println("copyData[0]:",copyData[0],"copyData[elementCount-1]:",copyData[elementCount-1])//0 999
    //复制原始数据从4-6（不包含）
    copy(copyData,srcData[4:6]) //srcData[4:6] #4,5
    //fmt.Println(srcData[4:6],copyData)//[4,5] [4,5,2,3,4...]
	for i:=0;i<5;i++{
		fmt.Printf("%d",copyData[i])//45234
	}

	//F.从切片中删除元素
	//GO中切片删除元素本质：以被删除元素为分界点，将前后两个部分的内存重新连接起来
	seq:=[]string{"a","b","c","d"}
	index:=2 //指定删除位置 "c'
	//查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index],seq[index:]) //[a,b][c,d]
	//将删除点前后的元素连接起来
	seq=append(seq[:index],seq[index+1:]...) //...表示将seq[index+1:]的数组整个加入到seq[:index]
	fmt.Println(seq) //[a b d]
}

//3。映射(map)--建立事物关联的容器
/*
大多数语言中，映射关系容器使用两种算法：散列表和平衡树
(1)散列表：
简单描述为一个数组（俗称"桶"），数组的每个元素为一个列表
(2)平衡树：
类似有父子关系的一棵数据树
*/
//map使用散列表hash实现
func MapType(){
	//[1]。添加关联到map并访问关联和数据
	//map定义：map[keyType]valueType
	kv:=make(map[string]int) //key为string类型，值为int类型
	kv["a"]=1 //为key=a赋值1
	fmt.Println(kv["a"],kv["b"]) //1 0

	//B.查看某个键是否在map中存在
	v,ok:=kv["b"]
	fmt.Println(v,ok) //0 false

	//[2].map声明即填充内容
	m:=map[string]string{
		"a":"A",
		"b":"B",
		"c":"C",
	}
	fmt.Println(m) //map[a:A b:B c:C]

	//[3].遍历map的"键值对"--访问每一个map中的关联关系
	for k,v:=range m{
		fmt.Println(k,v) //a A / b B / c C
	}
	for _,v:=range m{ //仅遍历值，_代表垃极桶，丢掉key值
		fmt.Println(v)//A B C
	}

	//[4]。特定顺序的遍历结果(排序:sort.Strings(slice切片):对传入的字符串切片进行字符串的升序排列)
	kv2:=make(map[string]int)
	kv2["C"]=66
	kv2["B"]=8
	kv2["A"]=88
    var kvList []string
	for k:=range kv2{
		kvList=append(kvList,k)
	}
	fmt.Println("kvList:",kvList) //[C B A]
	//sort.Strints作用是对传入的字符串切片进行字符串的升序排列
	sort.Strings(kvList) //对切片进行排序
	fmt.Println("order kvList:",kvList)//[A B C]

	//[5]。使用delete()函数从map中删除键值对
	//delete(map，键) #map为要删除的map实例，键为要删除的map键值对中的键
	delete(kv2,"C") //从kv2集合中删除C
	for k,v:=range kv2{
		fmt.Println(k,v) //B 8 / A 88
	}

	//[6].清空map中的所有元素
	//方法：重新make一个新的map  make(map[string]int)

	//[7]能够在并发环境中使用的map--sync.Map
	//注：GO语言中的map在并发情况下，只读是线程安全的，同时读写线程不安全
	m2:=make(map[int]int) //创建一个int到int的映射

	//开启一段并发代码
	go func() {
		//不停地对map进行写入
		for{
			m2[1]=1
		}
	}()

	//开启一段并发代码
	go func() {
		//不停地对map进行读取
		for{
			_=m2[1]
		}
	}()

	//无限循环，让并发程序在后台执行
	//for {} //报错，并发的map读写两个并发函数不断地对map进行读写，产生了竞态

	//并发读写一般做法是加锁，但效率低。

	//[8]sync.Map:效率较高且并发安全 --go1.9提供
	//注意：并发情况下应用sync.Map，因会对性能造成损失，所以非并发情况下用map
	var syncMap sync.Map
	//将键值对保存到sync.Map
	syncMap.Store("C",3)
	syncMap.Store("B",2)
	syncMap.Store("A",1)

	fmt.Println("syncMap:",syncMap)
	//从sync.Map中根据键取值
	val,ok:=syncMap.Load("B")
	fmt.Println("val:",val,"ok:",ok) //2 true
	//根据键删除对应的键值对
	syncMap.Delete("C")
	//遍历所有sync.Map中的键值对
	syncMap.Range(func(k,v interface{}) bool{
		fmt.Println("iterater:",k,v) // B 2 / A 1
		return true
	})


}

//4.列表(list) --可以快速增删的非连续空间的容器
/*
列表是一种非连续存储的容器，由多个节点组成，节点间通过一些变量记录彼此间关系
列表有多种实现方法，如单链表，双链表等
Go语言中，列表使用container/list包实现
*/
func ListType(){
	//(1)初始化列表
	//list初始化两种方式：New(变量名:=list.New())和声明(var 变量名 list.List)
	//注：列表与切片和map不同的是：列表没有具体元素类型限制，列表的元素可以为任意类型

	//(2)在列表中插入元素
	l:=list.New() //list初始化,创建一个列表实例 #变量名:=list.New()
	l.PushBack("backInsert") //后插
	l.PushFront("fontInsert") //前插
	fmt.Println(l.Len()) //2

	//(3)从列表中删除元素
	l2:=list.List{} //创建列表实例 #var 变量名 list.List
	l2.PushBack("back") //尾部添加 #back
	l2.PushFront("front") //头部添加 #front back
	element:=l2.PushBack("end")//尾部添加后保存元素句柄  #front back end
	l2.InsertAfter("ending",element) //在end后添加ending #front back end ending
	l2.InsertBefore("first",element) //在font前添加first #front back first end ending

	l2.Remove(element)//移除element变量对应的元素 #front back first ending
	fmt.Println(l2.Len()) //4

	//(4)遍历列表，访问列表的每一个元素
	for i:=l2.Front();i!=nil;i=i.Next(){
		fmt.Println(i.Value) //front back end ending
	}
}
