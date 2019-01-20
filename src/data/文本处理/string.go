package 文本处理

import (
	"strings"
	"fmt"
	"strconv"
)

/*
字符串处理: 标准库strings / strconv
*/
func StringDemo1(){
	str:="hello string"
	/*
	strings.Contains(原字符串，目标字符串) //包含
	*/
	bool:=strings.Contains(str,"hello")
	fmt.Println(bool) //true
	fmt.Println(strings.Contains("","")) //true

	/*
	连接字符串
	*/
	str1:=[]string{"A","B","C"}
	str11:=strings.Join(str1,",")
	fmt.Println(str11) //A,B,C

	/*
	查找字符串
	*/
	loc:=strings.Index(str,"string")
	fmt.Println(loc) //6

	/*
	重复字符串
	*/
	repeat:=strings.Repeat("hello",2)
	fmt.Println(repeat) //hellohello

	/*
	替换字符串
	-1<0表示全部替换  >0表示替换次数
	*/
	replace:=strings.Replace(str,"hello","hi",-1)
	fmt.Println(replace) //hi string
	fmt.Println(strings.Replace("ok in ok in ok","ok","ha",2)) //ha in ha in ok

	/*
	分割字符串
	*/
	split:=strings.Split("a,b,c",",")
	fmt.Println(split) //[a b c]

	/*
	去除指定字符串
	*/
	trimstr:=strings.Trim("!AB!C!","!")
	fmt.Println(trimstr) //AB!C

	/*
	去除s字符串的空格符,并按空格分割返回slice
	*/
	fields:=strings.Fields(" A B  C  ")
	fmt.Println(fields) //[A B C]
}

/*
字符串转换
*/
func StrconvDemo(){

	str:=make([]byte,0,10)
	fmt.Println(str, len(str),cap(str)) //[] 0 10
	str=strconv.AppendInt(str,10,4)

	fmt.Println(string(str))

	str=strconv.AppendBool(str,false)

	str=strconv.AppendQuote(str,"abc")

	str=strconv.AppendQuoteRune(str,'单')

	fmt.Println(string(str))

}

/*
format系统函数将其他类型转换为字符串
*/
func FormatDemo(){
	a:=strconv.FormatBool(false) //false

	b:=strconv.FormatFloat(123.23,'g',12,64) //123.23

	c:=strconv.FormatInt(1234,10) //1234

	d:=strconv.FormatUint(12345,10) //12345

	e:=strconv.Itoa(1023) //1023

	fmt.Println(a,b,c,d,e) //false 123.23 1234 12345 1023
}

/*
parse系列函数把字符串转换为其他类型
*/
func ParseDemo(){
	a,err:=strconv.ParseBool("false")
	if err!=nil{
		fmt.Println(err)
	}
	b,err:=strconv.ParseFloat("123.23",64)
	if err!=nil{
		fmt.Println(err)
	}

	c,err:=strconv.ParseInt("1234",10,64)
	if err!=nil{
		fmt.Println(err)
	}
	d,err:=strconv.ParseUint("12345",10,64)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(a,b,c,d) //false 123.23 1234 12345

}