package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maninbule/golang-login-system-with-jwt-gin/loadinit"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
}

func InitDB() {
	db = loadinit.GetDB()
	db.AutoMigrate(&User{})
}

func GetuserByEmail(email string) User {
	var user User
	db.Debug().Model(&User{}).Where("email = ?", email).First(&user)
	return user
}

func (u *User) CreateUser() bool {
	create := db.Create(u)
	if create.Error != nil {
		return false
	}
	return true
}
