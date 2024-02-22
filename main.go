package main

import (
	"fmt"
	"github.com/maninbule/golang-login-system-with-jwt-gin/loadinit"
)

func main() {
	loadinit.LoadEnvVaria()
	loadinit.InitDb()
	fmt.Println("加载数据库完成")
}
