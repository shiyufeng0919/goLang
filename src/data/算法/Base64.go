package 算法

import (
	"encoding/base64"
	"fmt"
)

/*
加密和解密数据

对称加密算法----Base64(对敏感数据加密与解密)
*/
func base64Encode(src []byte)[]byte{
	return []byte(base64.StdEncoding.EncodeToString(src))
}
func base64Decode(src []byte)([]byte,error){
	return base64.StdEncoding.DecodeString(string(src))
}
func Base64Demo(){
	//encode:；加密
	h:="hello base64"
	debyte:=base64Encode([]byte(h))
	fmt.Println(debyte)

	//decode：解密
	enbyte,err:=base64Decode(debyte)
	if err!=nil{
		fmt.Println(err.Error())
	}
	if h!=string(enbyte){
		fmt.Println("h!=enbyte")
	}
	fmt.Println(string(enbyte)) //hello base64
}
