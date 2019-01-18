package ginFramework

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

/*
分组路由

熟悉Flask的同学应该很了解蓝图分组。Flask提供了蓝图用于管理组织分组api。
gin也提供了这样的功能，让你的代码逻辑更加模块化，同时分组也易于定义中间件的使用范围。
*/
func RouterGroupDemo1(){
	router:=gin.Default()
	v1:=router.Group("/v1")
	v1.GET("/login", func(context *gin.Context) {
		context.String(http.StatusOK,"v1 login")
	})

	v2:=router.Group("/v2")
	v2.GET("/login", func(context *gin.Context) {
		context.String(http.StatusOK,"v2 login")
	})
	router.Run(":8181")

	/*
	postman GET请求 ->http://localhost:8181/v1/login
	响应结果: v1 login

	GET请求 ->http://localhost:8181/v2/login
	响应结果：v2 login
	*/
}

/*
middleware中间件

golang的net/http设计的一大特点就是特别容易构建中间件。gin也提供了类似的中间件。
需要注意的是中间件只对注册过的路由函数起作用。
对于分组路由，嵌套使用中间件，可以限定中间件的作用范围。
中间件分为全局中间件，单个路由中间件和群组中间件。
*/
/*
全局中间件
*/
func MiddleWare() gin.HandlerFunc{
	return func(context *gin.Context) {
		fmt.Println("before middleware")
		context.Set("request","client request") //添加一个属性
		context.Next()
		fmt.Println("after middleware")
	}
}
func MiddleWareDemo1(){
	router:=gin.Default()
	router.Use(MiddleWare()) //装饰中间件,在之上的路由函数，将不会有被中间件装饰效果
	{
		router.GET("/middleware", func(context *gin.Context) {
			/*
			MustGet:该方法必须先装饰再使用，否则报错。可以使用GET。效果一样
			*/
			request:=context.MustGet("request").(string) //直接读取request值
			req,_:=context.Get("request")
			context.JSON(http.StatusOK,gin.H{
				"middle_request":request,
				"request":req,
			})
		})
	}
	router.Run(":8181")

	/*
	测试: postman发送GET请求-》http://localhost:8181/middleware
		{
		"middle_request": "client request",
		"request": "client request"
		}
	*/
}

/*
单个路由中间件 --> /before被装饰了中间件
*/
func MiddleWareDemo2(){
	router:=gin.Default()
	router.GET("/before",MiddleWare(), func(context *gin.Context) {
		request:=context.MustGet("request").(string)
		context.JSON(http.StatusOK,gin.H{
			"middle_request":request,
		})
	})
	router.Run(":8181")

	/*
	postman GET请求 -> http://localhost:8181/before
	*/
}

/*
群组中间件
*/
func MiddleWareDemo3(){
	router:=gin.Default()
	authorized:=router.Group("/v1",MiddleWare())
	authorized.POST("/login", func(context *gin.Context) {
		request:=context.MustGet("request").(string)
		context.JSON(http.StatusOK,gin.H{
			"middle_request":request,
		})
	})
	router.Run(":8181")
	/*
	postman POST请求 -》http://localhost:8181/v1/login
	*/
}

/*
中间件实践：
中间件最大的作用，莫过于用于一些记录log，错误handler，还有就是对部分接口的鉴权
*/
//简易鉴权中间件
func AuthMiddleWare() gin.HandlerFunc{
	/*
	从上下文请求中读取cookie,然后校对cookie，若有问题，则终止请求，直接返回
	*/
	return func(context *gin.Context) {
		if cookie,err:=context.Request.Cookie("session_id");err == nil{
			value:=cookie.Value
			fmt.Println(value)
			if value=="123"{
				context.Next()
				return
			}
		}
		context.JSON(http.StatusUnauthorized,gin.H{
			"error":"unauthorized",
		})
		context.Abort() //终止请求
		return
	}
}
func MiddleWareDemo4(){
	router:=gin.Default()
	//注册。设置cookie
	router.GET("/auth/sign", func(context *gin.Context) {
		cookie:=&http.Cookie{
			Name:"session_id",
			Value:"123",
			Path:"/",
			HttpOnly:true,
		}
		http.SetCookie(context.Writer,cookie)
		context.String(http.StatusOK,"login success")
	})
	//验证cookie
	router.GET("/home",AuthMiddleWare(), func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{"data":"home"})
	})
	router.Run(":8181")

	/*
	Postman GET请求 ->http://localhost:8181/auth/sign
	响应结果:login success

	Postman GET请求 -> http://localhost:8181/home
	响应结果：
	{
    	"data": "home"
	}
	*/
}