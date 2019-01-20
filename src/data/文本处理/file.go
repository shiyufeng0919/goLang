package 文本处理

import (
	"os"
	"fmt"
)

/*
文件操作:生成文件目录，文件(夹)编辑等操作

os包
*/
func FileDemo1()  {
	os.Mkdir("dir1",0777) //创建dir1目录，权限0777
	os.Mkdir("dir11/dir2/dir3",0777)//根据path创建多级子目录,权限0777

	err:=os.Remove("dir1") //移除目录
	if err !=nil{
		fmt.Println(err)
	}
	os.RemoveAll("dir11") //移除dir11目录包下的所有子目录
}

//文件操作--建立与打开文件
func FileDemo2(){
	userFile:="file.txt"
	fout,err:=os.Create(userFile) //创建文件
	if err!=nil{
		fmt.Println(userFile,err)
		return
	}
	defer fout.Close()
	for i:=0;i<10;i++{
		fout.WriteString("just a test!\r\n") //写文件
		fout.Write([]byte("just a test.\r\n"))
	}
}

//读文件
func FileDemo3(){
	userFile:="file.txt"
	file,err:=os.Open(userFile)
	defer file.Close()

	if err != nil{
		fmt.Println(userFile,err)
		return
	}
	buf:=make([]byte,1024)
	for{
		n,_:=file.Read(buf)
		if 0==n{
			break
		}
		os.Stdout.Write(buf[:n])
	}
}

/*
删除文件
GO语言中删除文件和删除文件夹是同一函数
*/
func FileDemo4(){
	os.Remove("file.txt")
}
