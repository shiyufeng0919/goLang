package main

import (
  "golandProject/goLang/src/framework/database"
)

/*
GO语言部署工具：Supervisord,upstart,daemontools (进程管理软件)
1.Supervisord(python编写)它会帮读者反映管理的应用程序转成daemon程序。可通过命令执行开启，关闭，重启等操作。
*/

func main()  {
  database.MongoDB()
}