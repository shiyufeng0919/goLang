package ginFramework

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"net/http"
)

/*
########异步协程
golang的高并发一大利器就是协程。gin里可以借助协程实现异步任务。因为涉及异步过程，
请求的上下文需要copy到异步的上下文，并且这个上下文是只读的。
*/
func AsyncDemo1(){
	router:=gin.Default()
	router.GET("/sync", func(context *gin.Context) {
		time.Sleep(5*time.Second)
		fmt.Println("done!in path"+context.Request.URL.Path)
	})

	router.GET("/async", func(context *gin.Context) {
		cCp:=context.Copy()
		go func() {
			time.Sleep(5*time.Second)
			fmt.Println("Done!in path",cCp.Request.URL.Path)
		}()
	})
	router.Run(":8181")
	/*
	postman测试GET请求->http://localhost:8181/sync & http://localhost:8181/async
	*/
}

/*
自定义router
gin不仅可以使用框架本身的router进行Run，也可以配合使用net/http本身的功能：
*/
func RouterDemo(){
	router:=gin.Default()
	http.ListenAndServe(":8181",router)
}

func RouterDemos(){
	router:=gin.Default()
	s:=&http.Server{
		Addr:":8181",
		Handler:router,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	s.ListenAndServe()
}
