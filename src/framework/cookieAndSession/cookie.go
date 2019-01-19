package cookieAndSession

import (
	"net/http"
	"time"
	"fmt"
)

/*
GO语言cookie中文网-https://studygolang.com/articles/5905

cookie:客户端机制
*/

func cookieDemo1(w http.ResponseWriter, req *http.Request){
	cookie:=http.Cookie{
		Name:"username",
		Value:"val",
		Expires:time.Time{},
	}
	//####设置cookie
   http.SetCookie(w,&cookie)

   //####读取cookie
   username,_:=req.Cookie("username")
   fmt.Fprint(w,username)

   //####循环读取方式
   for _,cookie:=range req.Cookies(){
   	  fmt.Fprint(w,cookie.Name)
   }
}