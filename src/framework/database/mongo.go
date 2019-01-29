package database

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

/*
MongoDB是Nosql中常用的一种数据库

docker安装mongo及开启用户认证:https://blog.csdn.net/yori_chen/article/details/81036149?utm_source=blogxgwz7

Step1: docker pull mongo
Step2: docker run --name mongo-master -p 27017:27017 -v ~/myLearn/mongo:/data/db -d mongo --auth
Step3: docker exec -it mongo-master bash //进入mongo容器
Step4: mongo
      >use admin  //切换
      >db.createUser({user:'admin',pwd:'admin',roles:[{role:'userAdminAnyDatabase',db:'admin'}]});
      >db.auth("admin","admin");
      >exit
    exit //退出容器
Step5:docker exec -it mongo-master mongo admin
Step6:
     >db.auth("admin","admin");
     >use mymongo
     >db.createUser({user:'syf',pwd:'syf',roles:[{role:'readWrite',db:'mymongo'}]});
     >db.auth("syf","syf");


######
ZBMAC-C02VX5Y7H:mongo shiyufeng$ docker exec -it fb8cf574967b mongo syf
MongoDB shell version v4.0.5
connecting to: mongodb://127.0.0.1:27017/syf?gssapiServiceName=mongodb
Implicit session: session { "id" : UUID("04549558-2fa2-4967-870e-f5d185607532") }
MongoDB server version: 4.0.5
######
Step6:use mymongo //创建/切换DB(若DB不存在则创建，否则切换)
Step7:db.dropDatabase() //删除数据库


golang操作mongo
Step1:go get gopkg.in/mgo.v2
Step2:
*/
type Persons struct {
	Name string
	Phone string
}
func MongoDB(){
	session,err:=mgo.Dial("127.0.0.1:27017")
	if err!=nil{
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic,true)

	c:=session.DB("mymongo").C("people") /*集合*/
	err=c.Insert(&Persons{"yufeng","136"},&Persons{"kaixin","139"}) /*向集合中加入文档*/

	if err!=nil{
		log.Fatal(err)
	}

	result:=Persons{}
	err=c.Find(bson.M{"name":"kaixin"}).One(&result)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("name:",result.Name)
	fmt.Println("phone:",result.Phone)
}