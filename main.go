package main

import (
	"fmt"
	"github.com/maninbule/golang-login-system-with-jwt-gin/loadinit"
	"github.com/maninbule/golang-login-system-with-jwt-gin/models"
	"github.com/maninbule/golang-login-system-with-jwt-gin/routers"
)

func main() {
	loadinit.LoadEnvVaria()
	loadinit.InitDB()
	fmt.Println("加载数据库完成")
	models.InitDB()
	routers.InitGin()
}
