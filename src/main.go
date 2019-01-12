package main

import "fmt"

func pkgFunc(){
  fmt.Println("call pkgFunc")
}

func main() {
  pkgFunc()
  fmt.Println("hello go build")
}


