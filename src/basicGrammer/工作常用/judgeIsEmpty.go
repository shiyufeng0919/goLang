package 工作常用

import (
	"fmt"
	"strconv"
)

/*
####本节知识点：
1。判断是否为空
2。...用法
*/

type data1 struct {
	Id int
	Name string
	Comment string
}
/*
   该示例：可以传入任意个string类型的参数
*/
/*
知识点: ...用法
...string:表示可以接收任意个string类型的参数
array...:表示数组被打散为各元素
*/
func judgeEmpty(param ...string) bool{ //可接收任意个string类型的参数
	fmt.Println("传入参数:",param) //传入参数: [1 isEmpty judge isEmpty]
	for _,value:=range param{
		fmt.Println("value值:",value)
		if value == ""{
			return true
		}
	}
	return false
}

func JudgeIsEmptyDemo1(){
	// 1.结构体参数
	data1:=&data1{ //实例化结构体并赋值
		Id:1,
		Name:"isEmpty",
		Comment:"仅限string类型",
	} //实例化并赋值

	isEmpty:=judgeEmpty(strconv.Itoa(data1.Id),data1.Name,data1.Comment)
	fmt.Println("结构体中有空参数:",isEmpty)

	// 2.数组参数
	array:=[]string{"A","B","C"}
	isEmpty=judgeEmpty(array...) //切片被打散传入
	fmt.Println("数组中有空参数",isEmpty)

	// 3.将两个数组合并
	array2:=[]string{"D","E"}
	array=append(array,array2...) //将array2数组打散，append到array数组中
	fmt.Println("array新数组值:",array) //[A B C D E]
}

/*
   该示例，可传入任意多任意类型参数，不限制
*/
func judgeIsEmpty(param ...interface{}) bool{
	fmt.Println("传入参数:",param) //传入参数: [1 isEmpty judge isEmpty]
	for _,value:=range param{
		fmt.Println("value值:",value)
		if value == ""{
			return true
		}
	}
	return false
}

func JudgeIsEmptyDemo2(){
	data1:=&data1{
		Id:2,
		Name:"interface",
		Comment:"无限类型",
	}
	isEmpty:=judgeIsEmpty(data1.Id,data1.Name,data1.Comment)
	fmt.Println("isEmpty:",isEmpty)
}
