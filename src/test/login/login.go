package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func  main(){
	gin.SetMode(gin.DebugMode)
	router:=gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("login", func(context *gin.Context) {
		context.HTML(http.StatusOK,"login.html",nil)
	})
	router.POST("login", func(context *gin.Context) {
		user:=context.PostForm("username")
		pass:=context.PostForm("password")
		if(user=="aaa"){
			fmt.Println("a")
			if(pass=="bbb"){
				fmt.Println("b")
				context.String(http.StatusOK,"login ok!!")
				fmt.Println("c")
			}
			fmt.Println("d")
		}else{
			context.String(http.StatusOK,"login err!")
			fmt.Println("e")
		}
	})
	
	
	router.Run(":9090")

}
