package web

import (
	"net/http"
	"fmt"
	"strings"
	"log"
)

/*

原生web http.net

########一.web工作方式

1。用户访问一个web站点过程

Client->访问google.com->DNS服务网络->返回IP地址->client端->请求IP地址向服务端->服务端返回内容给客户端

2。名词
(1)。URL：统一资源定位符，用于描述一个网络上的资源
(2)。DNS：(domain name system)域名系统，用于管理域名与IP的映射
(3)。HTTP协议：是WEB工作的核心，是一种让web服务器与浏览器(C端)通过internet发送与接收数据的协议，它建立在TCP协议之上，一般采用TCP的80端口
(4)。Cookie机制：web程序引入cookie机制来维护连接的可持续状态(同一客户端的本次与上次请求，对HTTP服务器它并不知道这两个请求是否来自同一客户端)
(5)。Dos(拒绝服务攻击) / DDos(分布式拒绝服务攻击)：是一种利用TCP协议缺陷，发送大量伪造TCP连接，从而使被攻击方资源耗尽(CPU满负荷或内存不足)的攻击方式
(6)。HTTP请求包(浏览器信息)
    Request包分为三部分：
    A。Request line:请求行  B。Request header(请求头) C。Body(主体)
    注：header & body间有个空行，用于分割请求头和消息体
(7)。fiddler抓取Get和POST信息
(8)。HTTP协议定义了与服务器交互的请求方法：GET | POST | PUT | DELETE /查|改|增|删
(9)。HTTP协议与Connection:keep-alive区别
     HTTP是一个无状态的面向连接的协议。无状态指协议对于事务处理没有记忆能力，服务器不知道客户端是什么状态。
     HTTP/1.1起默认开启了Keep-Alive保持连接特性（可设置保存时间）

*/

/*
########1.http包建立web服务器
*/
func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:",r.URL.Path)
	fmt.Println("scheme:",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v:=range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"hello http web") //这个写入到w的是输出到客户端的
}
func HttpWebDemo1(){
	http.HandleFunc("/",sayhelloName) //设置访问的路由
	err:=http.ListenAndServe(":8181",nil) //设置监听的端口
	if err !=nil{
		log.Fatal("listenAndServe:",err)
	}
}
/*
上述访问方式：
1。postman GET|POST 请求 ->http://localhost:8181/
2。go build->生成.exe-—》./main运行程序->地址栏直接访问 http://localhost:8181/
*/


/*
########2.GO语言如何使用web工作
########3.Go语言的http包详解
Go语言http包两个核心功能：conn / ServeMux
*/