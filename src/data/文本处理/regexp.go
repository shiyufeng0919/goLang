package 文本处理

import (
	"regexp"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

/*
正则: regexp标准包
API:https://golang.org/pkg/regexp/
*/
/*
1.MatchString:判断字符串是否符合我们的描述需求
*/
//判断是否为IP地址
func isIp(ip string)(b bool){
	//192.168.0.1
	if m,_:=regexp.MatchString("^[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1-3}\\.[0-9]{1,3}$",ip);!m{
		return false
	}
	return true
}

//判断是否为合法输入
func isValideInput(){
	if len(os.Args) ==1{
		fmt.Println("Usage:regexp[string]")
		os.Exit(1)
	}else if m,_:=regexp.MatchString("^[0-9]+$",os.Args[1]);m{
		fmt.Println("数字")
	}else{
		fmt.Println("非数字")
	}
}

func MatchStringDemo(){
	isIp:=isIp("127.0.0.1")
	fmt.Println("isIp:",isIp)

	isValideInput()
}

/*
2。通过正则获取内容

以爬虫为例说明如何使用正则来过滤或截取抓取到的数据
*/
func RegexpDemo1(){
	resp,err:=http.Get("http://www.baidu.com")
	if err!=nil{
		fmt.Println("http get error",err)
	}
	defer resp.Body.Close()

	body,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		fmt.Println("http read error",err)
		return
	}
	src:=string(body)

	//将HTML标签全部了转换成小写
	re,_:=regexp.Compile("\\<[\\S\\s]+?\\>")
	src=re.ReplaceAllStringFunc(src,strings.ToLower)

	//去除STYLE
	re,_=regexp.Compile("\\<style[\\S\\s]+?\\</style\\>>")
	src=re.ReplaceAllString(src,"")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")
	fmt.Println(strings.TrimSpace(src))
}

/*
Find相关函数
*/
func RegexpDemo2(){
	a := "I am learning Go language"
	re, _ := regexp.Compile("[a-z]{2,4}")

	//查找符合正则的第一个
	one := re.Find([]byte(a))
	fmt.Println("Find:", string(one))  //am

	//查找符合正则的所有slice,n小于0表示返回全部符合的字符串，否则返回指定的长度
	all := re.FindAll([]byte(a), -1)
	fmt.Println("FindAll", all)

	//查找符合条件的index位置、开始位置和结束位置
	index := re.FindIndex([]byte(a))
	fmt.Println("FindIndex", index)

	//查找符合条件的所有的index位置，n同上
	allindex := re.FindAllIndex([]byte(a), -1)
	fmt.Println("FindAllIndex", allindex)
	re2, _ := regexp.Compile("am(.*)lang(.*)")

	//查找Submatch,返回数组，第一个元素是匹配的全部元素，第二个元素是第一个()中的， 第三个是第二个()中的
	//下面的输出中，第一个元素是"am learning Go language"
	//第二个元素是" learning Go "，注意包含空格的输出
	//第三个元素是"uage"
	submatch := re2.FindSubmatch([]byte(a))
	fmt.Println("FindSubmatch", submatch)
	for _, v := range submatch {
		fmt.Println(string(v))
	}

	//定义和上面的FindIndex一样
	submatchindex := re2.FindSubmatchIndex([]byte(a))
	fmt.Println(submatchindex)

	//FindAllSubmatch,查找所有符合条件的子匹配
	submatchall := re2.FindAllSubmatch([]byte(a), -1)
	fmt.Println(submatchall)

	//FindAllSubmatchIndex,查找所有字匹配的index
	submatchallindex := re2.FindAllSubmatchIndex([]byte(a), -1)
	fmt.Println(submatchallindex)
}

/*
替换相关函数
*/
func RegexpDemo3(){
	src := []byte(`
            call hello alice
            hello bob
            call hello eve
        `)
	pat  :=  regexp.MustCompile(`(?m)(call)\s+(?P<cmd>\w+)\s+(?P<arg>.+)\s*$`)
	res := []byte{}
	for _, s := range pat.FindAllSubmatchIndex(src, -1) {
		res = pat.Expand(res, []byte("$cmd('$arg')\n"), src, s)
	}
	fmt.Println(string(res))
	/*
	打印结果:
	hello('alice')
	hello('eve')
	*/
}