package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/maninbule/golang-login-system-with-jwt-gin/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
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
	}
	// 登陆成功，发放token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user_db.Email,
		"exp": time.Now().Add(time.Second * 60).Unix(),
	})
	fmt.Println("secret", os.Getenv("secret"))
	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		fmt.Println("发放token失败 ", err)
		return
	}
	//c.SetSameSite(http.SameSiteLaxMode)
	//c.SetCookie("Authorization", tokenString, 20, "", "", false, true)
	c.Header("Authorization", "Bearer "+tokenString)
	c.JSON(http.StatusOK, "登陆成功")
}

func Validata(c *gin.Context) {
	email, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "验证失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "验证成功",
		"用户":   email,
	})
}
