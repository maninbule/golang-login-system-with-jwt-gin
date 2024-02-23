package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maninbule/golang-login-system-with-jwt-gin/models"
	"log"
	"net/http"
	"os"
	"time"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	//tokenString, err := c.Cookie("Authorization")
	if len(tokenString) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "Authorization 字段为空",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token错误或者过期")
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "过期的token")
			return
		}
		email, _ := claims.GetSubject()
		fmt.Println("email ", email)
		user_db := models.GetuserByEmail(email)
		if user_db.ID == 0 {
			fmt.Println("user_db.ID == 0")
			c.AbortWithStatusJSON(http.StatusUnauthorized, "过期的cookie")
			return
		}
		c.Set("user", email)
	} else {
		log.Fatal(err)
	}
	c.Next()
}
