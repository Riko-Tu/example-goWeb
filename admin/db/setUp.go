package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


func SetUp() error {
	db, err := gorm.Open("mysql", "root:123456@(172.20.112.1:3308)/exercise.sql?charset=utf8mb4&parseTime=True")
	defer db.Close()
	return err
}
