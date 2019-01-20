package 文本处理

import (
	"net/http"
	"html/template"
)

/*
模版
*/

func handler(w http.ResponseWriter,r *http.Request){
	/*
	GO语言模版：先获取数据，再渲染数据
	*/
	t:=template.New("some template") //创建模版
	t,_=t.ParseFiles("src/templates/welcome.html") //解析模版文件
	//user:=GetUser() //获取当前用户信息
	//t.Execute(w,user) //执行模版的merge操作
}
