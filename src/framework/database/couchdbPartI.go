package database

import (
	"github.com/joelnb/sofa"
	"time"
	"fmt"
)

/*
CouchDB是Apache组织发布的一款开源的、面向文档类型的NoSQL数据库。由Erlang编写，使用json格式保存数据。CouchDB以RESTful的格式提供服务
可以很方便的对接各种语言的客户端
CouchDB最大的竞争对手就是熟悉的MangoDB。

docker中couchdb安装配置图解:https://www.linuxidc.com/Linux/2017-03/142405.htm

优秀的Go存储开源项目和库:https://studygolang.com/articles/9434
*/

func CouchDB(){
	conn,err:=sofa.NewConnection("http://127.0.0.1:32769",10*time.Second,sofa.NullAuthenticator())

	if err !=nil{
		panic(err)
	}

	db:=conn.Database("mydb")
	doc:=&struct {
		sofa.DocumentMetadata
		Name string `json:"name"`
		Type string `json:"type"`
	}{
		DocumentMetadata:sofa.DocumentMetadata{
			ID:"fruit",
		},
		Name:"apple",
		Type:"fruit",
	}
	rev,err:=db.Put(doc)
	if err!=nil{
		panic(err)
	}
	fmt.Println(rev)
}