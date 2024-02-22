package loadinit

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var db *gorm.DB

func InitDb() {
	user := os.Getenv("user")
	password := os.Getenv("password")
	address := os.Getenv("address")
	port := os.Getenv("port")
	databasename := os.Getenv("database")
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, address, port, databasename)
	var err error
	db, err = gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}
