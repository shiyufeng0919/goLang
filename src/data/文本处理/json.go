package 文本处理

import (
	"encoding/json"
	"fmt"
)

/*
json(javascript object notation)处理
*/
/*
解析json -> struct
*/
type Server1 struct{
	ServerName string
	ServerIP string
}
type Serverslice struct{
	Servers []Server1
}
/*
json -> struct
*/
func JsonDemo1(){
	var s Serverslice
	str:=`{"Servers":[{"ServerName":"shanghai_vpn","ServerIP":"0.0.0.1"},{"ServerName":"beijing_vpn","ServerIP":"0.0.0.2"}]}`
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s) //{[{shanghai_vpn 0.0.0.1} {beijing_vpn 0.0.0.2}]}
}

/*
struct -> json
*/
func JsonDemo2(){
	var s Serverslice
	s.Servers=append(s.Servers,Server1{"shanghai_vpn","0.0.0.1"})
	s.Servers=append(s.Servers,Server1{"beijing_vpn","0.0.0.2"})

	b,err:=json.Marshal(s)
	if err!=nil{
		fmt.Println("json err:",err)
	}
	fmt.Println(string(b))
	/*
	结果：{"Servers":[{"ServerName":"shanghai_vpn","ServerIP":"0.0.0.1"},{"ServerName":"beijing_vpn","ServerIP":"0.0.0.2"}]}
	*/
}
