package 避坑与技巧

import "fmt"

/*

第12章 "避坑"与技巧

12.3 接口的nil判断

nil在GO语言中只能被赋值给指针和接口

1。接口与nil不相等
*/
type Stringer interface {
	String() string
}
//定义一个结构体
type myImplement struct {

}
//实现fmt.Stringer的String方法
func (m *myImplement) String() string{
	return "hi"
}
//在函数中返回fmt.Stringer接口
func GetStringer() fmt.Stringer{
	//赋nil
	var s *myImplement=nil
	//返回变量
	return s
}
func Demo6(){
	//判断返回值是否为nil
	if GetStringer()==nil{
		fmt.Println("GetStringer()==nil")
	}else{
		fmt.Println("GetStringer()!=nil")
	}
	/*
	注：上述value为nil，但type为*MyImplement。使用==判断，不相等
	*/
}
/*
2.发现nil类型值返回时直接返回nil

为避免误判断的问题，发现带有nil的指针时直接返回nil
*/
func GetStringernew() fmt.Stringer{
	var s *myImplement=nil
	if s==nil{
		return nil
	}
	return s
}


