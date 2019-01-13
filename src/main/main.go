package main

import (
  "fmt"
  "golandProject/goLang/src/basicGrammer/避坑与技巧"
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
  //避坑与技巧.Demo1()
  //避坑与技巧.Demo2()
  //避坑与技巧.Demo7()
  避坑与技巧.Demo8()
}