package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maninbule/golang-login-system-with-jwt-gin/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signup(c *gin.Context) {
	type user struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var u user
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, "不完整的账号和密码")
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	u.Password = string(password)

	dbuser := models.User{
		Email:    u.Email,
		Password: u.Password,
	}
	if !dbuser.CreateUser() {
		c.JSON(http.StatusBadRequest, "注册失败")
		return
	} else {
		c.JSON(http.StatusBadRequest, "注册成功")
		return
	}

}
