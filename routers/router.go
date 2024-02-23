package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maninbule/golang-login-system-with-jwt-gin/controllers"
	"github.com/maninbule/golang-login-system-with-jwt-gin/middlewares"
)

func InitGin() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/signin", controllers.Signin)
	r.GET("/validate", middlewares.RequireAuth, controllers.Validata)
	fmt.Println("服务器开始运行，端口：8080")
	r.Run(":8080")
}
