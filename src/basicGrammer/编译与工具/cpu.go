package 编译与工具

import (
	"github.com/pkg/profile"
	"time"
)

/*
性能分析
*/
func joinSlice()[]string{
	var arr []string
	for i:=0;i<100000;i++{
		/*
		故意造成多次的切片添加(append操作)，由于每次操作可能会有内存重新分配和移动，性能较低
		*/
		arr=append(arr,"arr") //消耗性能
	}
	return arr
}

/*
首先下载包依赖:go get github.com/pkg/profile
*/
func main(){
	//开始性能分析，返回一个停止接口
	stopper:=profile.Start(profile.CPUProfile,profile.ProfilePath(".")) //指定输出的分析文件路径，当前文件夹
	//在main()函数结束时停止性能分析
	defer stopper.Stop()
	//分析的核心逻辑
	joinSlice()
	//让程序至少运行1s
	time.Sleep(time.Second)
}
/*
性能分析需要可执行配合才能生成分配的结果

$ go build -o cpu cpu.go  #将cpu.go编译为可执行文件cpu
$ ./cpu #运行可执行文件，在当前目录输出cpu.pprof文件
$ go tool pprof --pdf cpu cpu.pprof > cpu.pdf #使用go tool工具链输入cpu.pprof和cpu可执行文件，生成pdf格式的输出文件
#将输出文件重定向为cpu.pdf文件.这个过程会调用Graphviz工具。windows下需将Graphviz的可执行目录添加到PATH中
*/

//优化上述代码
func joinSlicenew() []string{
	const count=100000 //将切片预分配count数量，避免append()函数多次分配
	var arr []string=make([]string,count)
	for i:=0;i<count;i++{
		arr[i]="arr" //预分配后直接对每个元素进行赋值
	}
	return arr
}