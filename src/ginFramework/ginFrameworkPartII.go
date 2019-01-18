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
		fmt.Println("before middleware")
	}
}
func MiddleWareDemo1(){
	router:=gin.Default()
	router.Use(MiddleWare())
	{
		//router.
	}
}