package 文本处理

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

/*
文本：字符串 / 数字 / JSON / XML
*/
/*
XML处理：编/解码XML文件
*/
/*
解析XML(解析XML成对应的struct对象)

注：为正确解析，GO语言XML包要求struct定义中的所有字段必须是可导出的（即首字线大写）
*/
type Recurlyservers struct {
	XMLName xml.Name `xml:"servers"`     //<servers version="1">
	Version string `xml:"version,attr"`
	Svs []server `xml:"server"`         //<server></server><server></server>...
	Description string `xml:",innerxml"` //<?xml version="1.0" encoding="utf-8"?>
}
type server struct{
	XMLName xml.Name `xml:"server"`
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
func XMLDemo1(){
	file,err:=os.Open("src/data/文本处理/serverTest.xml")
	if err !=nil{
		fmt.Printf("error:%v",err)
		return
	}
	defer file.Close()
	data,err:=ioutil.ReadAll(file)
	if err!=nil{
		fmt.Printf("error:%v",err)
		return
	}
	v:=Recurlyservers{}
	err=xml.Unmarshal(data,&v) //解析
	if err!=nil{
		fmt.Printf("error:%v",err)
		return
	}
	fmt.Println(v)

	/*
	打印结果:
	{{ servers} 1 [{{ server} Shanghai_VPN 127.0.0.1} {{ server} Beijing_VPN 127.0.0.2}]
    <server>
        <serverName>Shanghai_VPN</serverName>
        <serverIP>127.0.0.1</serverIP>
    </server>
    <server>
        <serverName>Beijing_VPN</serverName>
        <serverIP>127.0.0.2</serverIP>
    </server>
    }
	*/
}

/*
输出xml:

xml包-> Marshal & MarshalIndent(会增加前缀和缩进)两个函数
*/
type Servers struct{
	XMLName xml.Name `xml:"servers"`
	Version string `xml:"version,attr"`
	Svs []Server `xml:"server"`
}
type Server struct {
	ServerName string `xml:"serverName"`
	ServerIP string `xml:"serverIP"`
}
func XMLDemo2(){
	v:=&Servers{Version:"1"}
	v.Svs=append(v.Svs,Server{"shanghai_vpn","0.0.0.1"})
	v.Svs=append(v.Svs,Server{"beijing_vpn","0.0.0.2"})
	output,err:=xml.MarshalIndent(v," ","   ")
	if err !=nil{
		fmt.Printf("error:%v\n",err)
	}
	os.Stdout.Write([]byte(xml.Header))//因Marshal & MarshalIndent不带xml头，加xml包预定义的header变量
	os.Stdout.Write(output)

	/*
	输出结果:
	<?xml version="1.0" encoding="UTF-8"?>
		 <servers version="1">
			<server>
			   <serverName>shanghai_vpn</serverName>
			   <serverIP>0.0.0.1</serverIP>
			</server>
			<server>
			   <serverName>beijing_vpn</serverName>
			   <serverIP>0.0.0.2</serverIP>
			</server>
		 </servers>
	*/
}