package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maninbule/golang-login-system-with-jwt-gin/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Signin(c *gin.Context) {
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
	user_db := models.GetuserByEmail(u.Email)
	if user_db.ID == 0 {
		c.JSON(http.StatusBadRequest, "邮箱不存在")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(u.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, "密码错误")
		return
	} else {
		c.JSON(http.StatusBadRequest, "登陆成功")
		return
	}

}
