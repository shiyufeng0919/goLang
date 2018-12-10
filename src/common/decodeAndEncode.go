package common

import (
	"fmt"
	"net/url"
)

/*
url编码与解码(谷歌浏览器会自动请求的url中的中文编码，因此需要解码。)
*/
func UrlDecode(){
	originUrl := "http://www.baidu.com/s?wd=自由度"
	fmt.Println("originUrl:",originUrl) // http://www.baidu.com/s?wd=自由度
	encodeUrl:= url.QueryEscape(originUrl)//url编码
	fmt.Println("encodeUrl:",encodeUrl) //http%3A%2F%2Fwww.baidu.com%2Fs%3Fwd%3D%E8%87%AA%E7%94%B1%E5%BA%A6
	decodeUrl,err := url.QueryUnescape(encodeUrl)//url解码
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("decodeUrl:",decodeUrl)//http://www.baidu.com/s?wd=自由度
}
