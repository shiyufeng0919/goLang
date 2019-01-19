package ginFramework

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"os"
	"log"
	"io"
	"github.com/gin-gonic/gin/binding"
)

/*
#####Gin框架使用:Gin是一个轻巧而强大的golang web框架
*/

func RouterDemo1(){
	router:=gin.Default() //该方法创建一个路由handler
	router.GET("/", func(context *gin.Context) { //request & response都封装到了gin.Context上下文
		context.String(http.StatusOK,"hello Gin")
	})
	router.Run(":8181") //启动路由的Run方法监听接口
}

/*
###restful路由
gin的路由来自于httprouter库，因此httprouter具有的功能，gin也具有。但是gin不支持路由正则表达式
*/
/*
 ########参数传递
*/
func RestfulRouter1(){
	/*
	 Step1：gin.Default()：该方法创建一个路由handler。
	*/
	router:=gin.Default()
	/*
	Step2：通过http方法绑定路由规则和路由函数
	注：不同于net/http库的路由函数，gin进行了封装，把request & response都封装到了gin.Context的上下文环境

	注：/usr/a123 | /usr/abc均可以匹配。但/usr/ 和 /usr/a123/不会被匹配
	*/
	router.GET("/usr/:name", func(context *gin.Context) {
		name:=context.Param("name") //读取路由参数中的值
		context.String(http.StatusOK,"hello %s",name)
	})

	/*
	Step3：启动路由的Run方法监听接口
	*/
	router.Run(":8181") //指定运行端口

	/*
	  postman测试:访问url -> http://localhost:8181/usr/gin ,GET请求，输出hello gin
	*/
}

/*
除了: gin还提供了*号处理参数，*号匹配的规则更多
*/
func RestfulRouter2(){
  router:=gin.Default()
  router.GET("/usr/:name/*action", func(context *gin.Context) {
	  name:=context.Param("name")
	  action:=context.Param("action")
	  message:=name+" is"+action
	  context.String(http.StatusOK,message)
  })
  router.Run(":8181")

  /*
  postman请求 -》 http://localhost:8181/usr/gin/framework
  结果: gin is/framework
  */
}

/*
query string参数与body参数

C端 -> S端 发送请求，参数分为：
1。路由参数 (http://localhost:8181/usr/:name ,则:name为路由参数)
2。查询字符串 query string (http://localhost:8181/usr?key1=val1&key2=val2 .key-value是经过urlencode编码)
3。报文体body参数
*/
func QueryStringRouter(){
	router:=gin.Default() //创建一个路由handler
	router.GET("/querystring", func(context *gin.Context) {
		/*
		DefaultQuery -> 获取参数firstname，如无该参数(只作用于key不存在情况)，则给该参数设置默认值Guest
		*/
		firstname:=context.DefaultQuery("firstname","Guest")
		/*
		Query -> 获取参数 ,即使没有该参数也不会报错
		*/
		lastname:=context.Query("lastname")
		//响应返回结果为string类型数据
		context.String(http.StatusOK,"Hello %s %s",firstname,lastname)
	})
	router.Run() //不指定端口，默认端口8080
	/*
	postman访问 -> http://localhost:8080/querystring?firstname=gin&lastname=framework
	返回结果: Hello gin framework

	postman访问 - > http://localhost:8080/querystring
	返回结果: Hello Guest

	postman访问 -> http://localhost:8080/querystring?firstname=gin
	返回结果：Hello gin
	*/
}

/*
报文体参数,常见格式
1.application/json：JSON格式
2.application/x-www-form-urlencoded：即把query string的内容放在body体里，同样也需要urlencode(中文问题)
3.application/xml：
4.multipart/form-data：主要用于图片上传
默认情况下，context.PostForm解析的是2和4的参数
*/
func BodyRouterDemo1(){
	router:=gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message:=context.PostForm("message")
		nick:=context.DefaultPostForm("nick","anonymous")

		/*
		响应返回结果为JSON格式数据
		gin.H ：封装了生成json的方式，对于嵌套json的实现，嵌套gin.H即可
		*/
		context.JSON(http.StatusOK,gin.H{
			"status":gin.H{
				"status_code":http.StatusOK,
				"status":"ok",
			},
			"message":message,
			"nick":nick,
		})
	})
	router.Run()

	/*
	postman POST请求 -》 http://localhost:8080/form_post
	Body->选择x-www-form-urlencoded ->添加参数message->gin / nick->framework
	返回结果:
	 {
		"message": "gin",
		"nick": "framework",
		"status": {
			"status": "ok",
			"status_code": 200
		}
	 }
	*/
}

/*
同时使用查询字符串query string 与 body参数发送数据给服务器
*/
func BodyRouterDemo2(){
	router:=gin.Default()
	/*
	restful接口，PUT请求 (GET/POST/PUT/DELETE/OPTION)
	*/
	router.PUT("/post", func(context *gin.Context) {
		/*
		query string参数形式
		*/
		id:=context.Query("id")
		page:=context.DefaultQuery("page","0")
		/*
		body报文参数形式
		*/
		name:=context.PostForm("name")
		message:=context.PostForm("message")
		fmt.Printf("id:%s;page:%s;name:%s;message:%s \n",id,page,name,message)
		/*
		返回客户端响应为JSON格式
		*/
		context.JSON(http.StatusOK,gin.H{
			"status_code":http.StatusOK,
			"param":gin.H{
				id:id,
				page:page,
				name:name,
				message:message,
			},
		})
	})
	router.Run() //默认8080端口

	/*
	postman PUT请求 -> http://localhost:8080/post?id=1&page=10
	Body->选择x-www-form-urlencoded ->添加参数 name->gin / message->framework
	返回响应结果:
	 {
		"param": {
			"1": "1",
			"10": "10",
			"framework": "framework",
			"gin": "gin"
		},
		"status_code": 200
	  }
	*/
}

/*
###############文件上传
*/
/*
1.上传单个文件
multipart/form-data
*/
func FileUploadDemo1(){
	router:=gin.Default()
	router.POST("/upload", func(context *gin.Context) {
		name:=context.PostForm("name") //body报文参数
		fmt.Println("name:",name)
		/*
		注意：upload为context.Request.FormFile指定的参数，其值必须为绝对路径
		*/
		file,header,err:=context.Request.FormFile("upload") //解析客户端文件name属性
		if err != nil{
			context.String(http.StatusBadRequest,"Bad Request")
			return
		}
		filename:=header.Filename
		fmt.Println(file,err,filename)

		out,err:=os.Create(filename) //os操作把文件复制到硬盘上
		if err !=nil{
			log.Fatal(err)
		}
		defer out.Close()
		_,err=io.Copy(out,file)
		if err!=nil{
			log.Fatal(err)
		}

		//返回响应
		context.String(http.StatusOK,"上传成功,文件名为 %s",filename)
	})
	router.Run()

	/*
	postman Post请求 ->http://localhost:8080/upload
	Body->选择form-data
	    Key(选择file而非Text)=upload，value上传文件(可以为图片，也可以为文件)
	    Key(Text)=name,value->测试上传
	响应结果: 上传成功,文件名为 1.jpeg
	*/
}

/*
上传多个文件
*/
func FileUploadDemo2(){
	router:=gin.Default()
	router.POST("/multi/upload", func(context *gin.Context) {
		err:=context.Request.ParseMultipartForm(200000)
		if err !=nil{
			log.Fatal(err)
		}
		formdata:=context.Request.MultipartForm //得到文件句柄
		files:=formdata.File["upload"] //获取文件数据
		for i,_:=range files{ //遍历读写
			file,err:=files[i].Open()
			defer file.Close()
			if err!=nil{
				log.Fatal(err)
			}
			out,err:=os.Create(files[i].Filename)
			defer out.Close()
			if err!=nil{
				log.Fatal(err)
			}
			_,err=io.Copy(out,file)
			if err!=nil{
				log.Fatal(err)
			}
			context.String(http.StatusCreated,"upload success")
		}
	})
	router.Run()

	/*
	postman POST请求-》http://localhost:8080/multi/upload
	Body ->选择form-data
		Key(选择file)=upload,value(上传文件)
	    Key(选择file)=ipload,value(上传文件)
	    ... ##上传多个文件，key值一致，均为upload
	请求响应:
	  upload successupload success
	*/
}

/*
web的form表单上传.
HTML表单模版在templates/目录下
*/
func FormUpload(){
	router:=gin.Default()
	/*
     LoadHTMLGlob:定义模版文件路径
	*/
	router.LoadHTMLGlob("src/templates/*")
	router.GET("/upload", func(context *gin.Context) {
		/*
		context.HTML渲染模版,可通过gin.H给模版传值
		*/
		context.HTML(http.StatusOK,"upload.html",gin.H{})
	})
	router.Run()
}

/*
参数绑定: JSON格式数据通信
content-type类型:application/json的格式
*/
type User struct {
	//binding表示必填
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd" json:"passwd" binding:"required"`
	Age      int    `form:"age" json:"age"`
}
func JsonParamDemo(){
	router:=gin.Default()
	router.POST("/login", func(context *gin.Context) {
		var user User
		var err error
		contentType:=context.Request.Header.Get("Content-Type")

		switch contentType{
		case "application/json":
			err=context.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err=context.BindWith(&user,binding.Form)
		}

		if err!=nil{
			fmt.Println(err)
			log.Fatal(err)
		}
		context.JSON(http.StatusOK,gin.H{
			"user":user.Username,
			"pwd":user.Passwd,
			"age":user.Age,
		})
	})
	router.Run()

	/*
	postman POST请求->http://localhost:8080/login
	Body->raw->json
	{
		"username": "kaixinyufeng",
		"passwd": "123",
		"age": 31
    }
	返回响应
	{
		"age": 31,
		"pwd": "123",
		"user": "kaixinyufeng"
     }

	注意：使用json还需要注意一点，json是有数据类型的，因此对于 {"passwd": "123"} 和 {"passwd": 123}是不同的数据类型，解析需要符合对应的数据类型,否则会出错
	*/
}

/*
gin提供Bind，更高级用法，根据content-type自动推断是bind表单还是JSON参数
*/
func JsonParamDemo2(){
	router:=gin.Default()
	router.POST("/login", func(context *gin.Context) {
		var user User
		err:=context.Bind(&user)
		if err!=nil{
			fmt.Println(err)
			log.Fatal(err)
		}
		context.JSON(http.StatusOK,gin.H{
			"username":user.Username,
			"pwd":user.Passwd,
			"age":user.Age,
		})
	})
	router.Run()

	/*
	postman POST请求-》http://localhost:8080/login
    	Body->raw->json格式
	    参数：
		{
			"username": "yufeng",
			"passwd": "123",
			"age": 31
		}
       响应结果:
		{
			"age": 31,
			"pwd": "123",
			"username": "yufeng"
		}
	*/
}


/*
多格式渲染

请求可以使用不同的content-type，响应也可以有html,text,plain,json,xml等
*/
func ReturnJsonAndXml(){
	router:=gin.Default()
	router.GET("/xml", func(context *gin.Context) {
		contentType:=context.DefaultQuery("content_type","json")
		if contentType=="json"{
			context.JSON(http.StatusOK,gin.H{
				"user":"kaixin",
				"pwd":123,
			})
		}else if contentType=="xml"{
			context.XML(http.StatusOK,gin.H{
				"user":"yufeng",
				"pwd":456,
			})
		}
	})
	router.Run()

	/*
	postman GET请求->http://localhost:8080/xml?content_type=xml
	返回响应:
	<map>
		<user>yufeng</user>
		<pwd>456</pwd>
	</map>

	postman GET请求->http://localhost:8080/xml?content_type=json
	返回响应:
	{
		"pwd": 123,
		"user": "kaixin"
	}
	*/
}

/*
#########重定向
*/
func RedirectDemo(){
	router:=gin.Default()
	router.GET("/redirect/google", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently,"https://google.com")
	})
	router.Run()
}
