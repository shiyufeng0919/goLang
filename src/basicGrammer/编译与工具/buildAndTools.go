package 编译与工具

/*
第11章 编译与工具

GO命令教程:https://www.kancloud.cn/cattong/go_command_tutorial/261347

说明：

$GOPATH工作目录结构，约定有三个子目录（需要自行创建）：

src ——存放源代码文件
pkg——存放编译后的文件
bin ——存放编译后的可执行文件

11.1 编译(go build)

go build : 该指令将源码编译为可执行文件

***注意go安装版本问题。否则执行go build会报一堆错误。

原因：本机安装了两个go版本，需要保留一个

$ which go
$ echo $GOROOT
$ go version


####go build多种编译方法:

11.1.1 go build无参数编译

go build #main包下对main.go生成二进制文件

./main 	#运行可执行文件，输出内容


11.1.2 go build + 文件列表

编译同目录的多个源码文件，可在go build后面指定多个文件名

go build main.go lib.go ...

#若需要指定输出可执行文件名，可使用-o参数

go build -o myexec main.go lib.go


11.1.3 go build + 包

go build + 包在设置GOPATH后，可直接根据包名进行编译，即便包内文件被增/删也不影响编译指令


11.2 编译后运行 go run

go run命令会编译源码，并且直接执行源码的main()函数，不会在当前目录留下可执行文件

go run main.go #直接打印输出

11.3 编译并安装 go install

go install 功能 《=》 go build

go install: 只是将编译的中间文件放在GOPATH的pkg目录下，以及固定地将编译结果放在GOPATH的bin目录下

注：go install 编译过程有如下规律：

（1）go install是建立在GOPATH上的，无法在独立的目录里使用go install
（2）gopath下的bin目录放置的是使用go install生成的可执行文件，可执行文件的名称来自于编译时的包名
（3）go install输出目录始终为gopath/bin目录，无法使用-o附加参数进行自定义
（4）gopath/pkg目录放置的是编译期间的中间文件

11.4 一键获取代码，编译并安装(go get)

go get可借助代码管理工具通过远程拉取/更新代码包及其依赖包，并自动完成编译和安装（整个过程尤如安装APP）

注：使用go get前，需要安装与远程包匹配的代码管理工具，如git/svn/hg等，参数中需要提供一个包名

11.5 测试

GO语言拥有一套单元测试和性能测试系统，仅需要很少的代码即可快速测试一段需求代码

性能测试系统可以给出代码的性能数据，帮助测试者分析性能问题

11.5.1 单元测试 --测试和验证代码的框架

参见unit_test.go

11.5.2 基准测试--获得代码内存占用和运行效率的性能数据

基准测试可以测试一段程序的运行性能及耗费CPU的程度。

GO语言提供了基准测试框架，使用方法类似于单元测试。使用者无须准备高精度计时器和各种分析工具。

基准测试本身可以打印出非常标准的测试报告

参见jizhun_test.go


11.6 性能分析(go pprof)--发现代码性能问题的调用位置

GO语言工具链中的go pprof可以帮助开发者快速分析及定位各种性能问题。如CPU消耗，内存分配及阻塞分析

go pprof工具链配合Graphviz图形化工具可将runtime.pprof包生成的数据转换为PDF格式，以图片的方式展示程序的性能分析结果

11.6.1 安装第三方图形化显示分析数据工具(Graphviz)

Graphviz：是一套通过文本描述的方法生成图形的工具包，描述文本的语言叫做DOT

www.graphviz.org

11.6.2安装第三方性能分析来分析代码包

runtime.pprof提供基础的运行分析的驱动

$ go get github.com/pkg/profile ##对runtime.pprof技术上进行便利封装

参见cpu.go
*/
