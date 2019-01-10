package reflect

import (
	"encoding/json"
	"fmt"
	"bytes"
	"reflect"
	"strconv"
	"errors"
)

/*
10.3 示例：将结构体的数据保存为JSON格式的文本数据

JSON格式是一种用途广泛的对象文本格式

通过反射获取结构体成员及各种值的过程，使用反射将结构体序列化为文本数据

注：避免字节数组和字符串的转换可以提高一些性能

1.数据结构及入口函数
*/
func ReflectDemo15(){
	//声明技能结构
	type Skill struct {
		Name string
		Level int
	}
	//声明角色结构
	type Actor struct {
		Name string
		Age int
		Skills []Skill
	}
	//填充基本角色数据
	a:=Actor{
		Name:"coolBoy",
		Age:18,
		Skills:[]Skill{
			{Name:"A",Level:1},
			{Name:"B",Level:2},
			{Name:"C",Level:3},
		},
	}

	//应用json.Marshal序列化
	if result,err:=json.Marshal(a);err==nil{
		fmt.Println(string(result)) //{"Name":"coolBoy","Age":18,"Skills":[{"Name":"A","Level":1},{"Name":"B","Level":2},{"Name":"C","Level":3}]}
	}else{
		fmt.Println(err)
	}


	//应用自定义函数序列化
	if result,err:=MarshalJson(a);err==nil{
		fmt.Println(result)//{"Name":"coolBoy","Age":18,"Skills":[{"Name":"A","Level":1},{"Name":"B","Level":2},{"Name":"C","Level":3}]}
	}else{
		fmt.Println(err)
	}

}
/*
2.序列化主函数

MarshalJson是对函数writeAny封装，将外部interface{}类型转换为内部reflect.Value类型，同时构建输出缓冲，
将一些复杂的操作简化，方便外部使用.
*/
func MarshalJson(v interface{})(string,error){ //接收任意类型值，转换为json字符串返回
	//准备一个缓冲
	var b bytes.Buffer //<=>StringBuilder,在大量字符串拼接时推荐使用该结构

	//将任意值转换为JSON并输出到缓冲中
	//&b:以指针方式传入，方便各种类型的数据都写入到bytes.Buffer中
	//v:转换为反射值对象并传入
	if err:=writeAny(&b,reflect.ValueOf(v));err==nil{
		return b.String(),nil //将bytes.Buffer内容转换为字符串返回
	}else{
		return "",err
	}
}

/*
3.任意值序列化

param:字节缓冲 和 反射值对象

将反射值对象转换为JSON格式并写入到字节缓冲
*/
func writeAny(buff *bytes.Buffer,value reflect.Value)error{
	/*
	可扩充switch种类扩充序列化可能识别的类型
	*/
	switch value.Kind() {
	case reflect.String:
		//写入带有双引号括起来的字符串
		/*
		value.String()：使用reflect.Value的String()函数将传入值转换为字符串
		strconv.Quote():该函数体供比较正规的封装
		buff.WriteString()：应用bytes.Buffer的WriteString()函数，将前面输出的字符串写入缓冲中
		*/
		buff.WriteString(strconv.Quote(value.String()))
	case reflect.Int:
		//将整型转换为字符串并写入缓冲中
		/*
		strconv.FormatInt(value.Int(),10):将传入值转换为整型，再将整型以10进制格式使用strconv.FormatInt()函数格式化为字符串
		最后写入缓冲中
		*/
		buff.WriteString(strconv.FormatInt(value.Int(),10))
	case reflect.Slice:
		//切片序列化为JSON操作
		return writeSlice(buff,value)
	case reflect.Struct:
		//结构体序列化为JSON操作
		return writeStruct(buff,value)
	default:
		//遇到不认识的种类，返回错误
		return errors.New("unsupport kind: "+value.Kind().String())
	}
	return nil
}
/*
4.切片序列化
*/
func writeSlice(buff *bytes.Buffer,value reflect.Value) error{
	//写入切片开始标记
	buff.WriteString("[")
	//遍历每个切片元素
	for s:=0;s<value.Len();s++{
		sliceValue:=value.Index(s)
		//写入每个切片元素
		writeAny(buff,sliceValue)
		//每个元素尾部都写入逗号，最后一个字段不添加
		if s<value.Len()-1{
			buff.WriteString(",")
		}
	}
	//写入切片结束标记
	buff.WriteString("]")
	return nil
}
/*
5.结构体序列化
*/
func writeStruct(buff *bytes.Buffer,value reflect.Value) error{
	//取值的类型对象
	valueType:=value.Type()
	//写入结构体左大括号
	buff.WriteString("{")
	//遍历结构体的所有值
	for i:=0;i<value.NumField();i++{
		//获取每个字段的字段值(reflect.value)
		fieldValue:=value.Field(i)
		//获取每个字段的类型(reflect.StructField)
		fieldType:=valueType.Field(i)
		//写入字段名左双引号
		buff.WriteString("\"")
		//写入字段名
		buff.WriteString(fieldType.Name)
		//写入字段名右双引号和冒号
		buff.WriteString("\":")
		//写入每个字段值
		writeAny(buff,fieldValue)
		//每个字段尾部写入逗号，最后一个字段不添加
		if i<value.NumField()-1{
			buff.WriteString(",")
		}
	}
	//写入结构体右大括号
	buff.WriteString("}")
	return nil
}
