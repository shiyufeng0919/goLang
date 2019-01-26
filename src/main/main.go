package main

import (
  "fmt"
  "golandProject/goLang/src/framework/database"
)

//模拟 go build命令
func pkgFunc(){
  fmt.Println("call pkgFunc")
}

func main1() {
  pkgFunc()
  fmt.Println("hello go build")
}


func main()  {
  //web.HttpWebDemo1()
  //文本处理.ParseDemo()
  database.LevelDB()
}