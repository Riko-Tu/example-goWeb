package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

var tefsDb *gorm.DB
func SetUp() error {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	port := viper.GetString("database.port")
	ip := viper.GetString("database.ip")
	dbName := viper.GetString("database.dbName")
	dbStr :=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		username,password,ip,port,dbName)
	var  err error
	db, err = gorm.Open("mysql", dbStr)
	db.SingularTable(true)
	return err
}

func GetTdb()  *gorm.DB{
	return tefsDb
}


func GetDB() *gorm.DB {
	return db
}

func SetUpT() error {
	username := "root"
	password := "12345"
	port := "4406"
	ip := "81.71.9.72"
	dbName := "tefs-admin"
	dbStr :=fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		username,password,ip,port,dbName)
	var  err error
	tefsDb, err = gorm.Open("mysql", dbStr)
	return err
}





