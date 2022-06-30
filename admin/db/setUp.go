package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func SetUp() error {
	db, err := gorm.Open("mysql", "root:123456@(localhost:80)/exercise?charset=utf8mb4&parseTime=True&loc=Local")
	defer db.Close()
	return err
}