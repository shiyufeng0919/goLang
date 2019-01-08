package basic

/*
第8章 包(package)

GO语言包与文件夹一一对应，所有与包相关的操作，必须依赖于工作目录(GOPATH)

GOPATH是GO语言中使用的一个环境变量，它使用绝对路径提供项目的工作目录。

工作目录是一个工程开发的相对参考目录。
*/
/*
8.1.1 使用命令行查看GOPATH信息

$ go env

输出:
GOARCH="amd64"  #表目标处理器架构
GOBIN="/usr/local/go/go1.8.3/bin" #编译器和链接器安装位置
GOEXE=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="darwin" #目标操作系统
GOPATH="/Users/shiyufeng/Documents/kaixinyufeng/goworkspace" #当前工作目录
GORACE=""
GOROOT="/usr/local/go/go1.8.3" #GO开发包的安装目录
GOTOOLDIR="/usr/local/go/go1.8.3/pkg/tool/darwin_amd64"
GCCGO="gccgo"
CC="clang"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/0b/7hwwpn550w72m_r_xx6n06zw1cy3lk/T/go-build368935419=/tmp/go-build -gno-record-gcc-switches -fno-common"
CXX="clang++"
CGO_ENABLED="1"
PKG_CONFIG="pkg-config"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
*/

/*
8.1.2 使用GOPATH的工程结构

在GOPATH指定的工作目录下，代码总是会保存在$GOPATH/src目录下

工程执行go build /go install /go get 等指令后：
会将产生的二进制可执行文件放在$GOPATH/bin目录下
生成的中间缓存文件会被保存在$GOPATH/pkg目录下

若需要将整个源码添加到版本管理工具中，则只需要添加$GOPATH/src目录的源码即可。bin和pkg目录的内容都可以由src目录生成。
*/

/*
8.1.3 设置和使用GOPATH

1.设置当前目录为GOPATH
export GOPATH=`pwd`

2.建立GOPATH中的源码目录
mkdir -p src/hello //连续创建包src/hello

func main(){
 fmt.Println("hello"))
}

3.添加main.go源码文件
vim main.go

4.编译源码并运行
go install test //编译完成的可执行文件会保存在$GOPATH/bin目录下，在bin目录下执行./

*/

/*
8.1.4 在多项目工程中使用GOPATH

GoLand分全局GOPATH和项目GOPATH(建议)

建议在开发时只填写项目GOPATH，不使用多个GOPATH和全局的GOPATH
*/

/*
8.2 创建包package --编写自己的代码扩展

包特性：
1。一个目录下的同级文件归属一个包
2。包名可以与其目录名不同
3。包名为main的包为应用程序的入口包，编译源码没有main包时，将无法编译输出可执行的文件
*/

/*
8.3 导出标识符 --让外部访问包的类型和值

首字母大写即可
*/

/*
8.4导入包(import)--在代码中使用其他的代码

8.4.1默认导入的写法
1.单行导入
import "package1"

2.多行导入
import(
   "package1"
   "package2"
   ......
 )

8.4.2导入包后自定义引用的包名
import newName "xx/xx/package"

8.4.3匿名导入包--只导入包但不使用包内类型和数值
import(
  _ "xx/xx/package"
)

8.4.4 包在程序启动前的初始化入口 init

init()函数特性
1。每个源码可以使用一个init()函数
2。init()会在程序执行前(main()函数执行前)被自动调用
3。调用顺序为main()中引用的包，以深度优化顺序初始化

如：main -> A ->B ->C ，则这些包的init()函数调用顺序为：
C.init() -> B.init() ->A.init() ->main

4。同一个包中的多个init()函数调用顺序不可预期
5。init()函数不能被其他函数调用
*/