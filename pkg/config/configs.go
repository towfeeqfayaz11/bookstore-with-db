package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	// here mysql is used in case we are using mysql db,
	// godb is the name of database
	// the name of our table inside "godb" database of "mysql" is "books"
	d, err := gorm.Open("mysql", "root:Root@123@/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("db connection error")
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
