package 算法

import (
	"crypto/sha256"
	"fmt"
	"io"
	"crypto/sha1"
	"crypto/md5"
	"golang.org/x/crypto/scrypt"
)

/*
########如何存储密码#####
加密算法：sha族(sha-1,sha-256,md5...)
实际应用：利用和已有的哈希算法进行多次哈希
安全性比较好的网站应用"加盐"方式存储密码(先将用户输入的密码进行一次MD5(或其他哈希算法)加密，将得到的MD5值前后加上一些只有管理员知道的随机串，再进行一次MD5加密)
https://golang.org/pkg/crypto/sha256/
*/
func AlgorithmDemo1(){
	/*
	单向哈希，value->hash 但hash 不能推出->value
	*/
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	fmt.Printf("%x", h.Sum(nil)) //a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447


	fmt.Println("")

	io.WriteString(h,"hello world\n")
	fmt.Printf("%x",h.Sum(nil))//ec498a36221dd860c6f24ea26cb29cec68a38479496f78e54ce35f34c8106847

	fmt.Println()

	/*
	有碰撞：即输入不同的值得出的哈希值相同
	*/
	h=sha1.New()
	io.WriteString(h,"hello world\n")
	fmt.Printf("%x",h.Sum(nil))//22596363b3de40b06f981fb85d82312e8c0ed511

	fmt.Println()

	/*
	双向，可逆
	*/
	h=md5.New()
	io.WriteString(h,"需要加密的pwd")
	fmt.Printf("%x",h.Sum(nil)) //4b6dee2a5d68413b674448020cf6af7c
}

/*
加盐：salt
*/
func AlgorithmDemo2(){
	/*
	eg:username=yufeng,pwd=123456
	*/
	h:=md5.New()
	io.WriteString(h,"123456")
	pwmd5:=fmt.Sprintf("%x",h.Sum(nil))
	fmt.Printf("%x",pwmd5)
	//指定两个salt
	salt1:="@#$%"
	salt2:="^&*()"
	//salt1+username+salt2+md5拼接
	io.WriteString(h,salt1)
	io.WriteString(h,"yufeng")
	io.WriteString(h,salt2)
	io.WriteString(h,pwmd5)
	last:=fmt.Sprintf("%x",h.Sum(nil))
	fmt.Printf("%x",last)
}

/*
增加密码计算所需耗费的资源和时间，使得任何人都不可获得足够的资源建立所需的rainbow table
推荐scrypt方案，scrypt是由著名FreeBSD黑客Colin Percival为他的备份服务Tarsnap开发
https://godoc.org/golang.org/x/crypto/scrypt
*/
func AlgorithmDemo3()  {
	salt:=[]byte("@#$%")
	dk, err := scrypt.Key([]byte("some password"), salt, 32768, 8, 1, 32)
	if err !=nil{
		fmt.Println(err)
	}
	fmt.Printf("%x",dk) //1590f9b4155924f2d335072fb9e17f84b6a5c835c3149ddaf3fbc70893ac6d53
}