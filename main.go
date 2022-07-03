package main

import (
	_ "turan/example-goWeb/admin/dataload"
	"turan/example-goWeb/admin/db"
	"turan/example-goWeb/admin/route"
)

func main() {
	err := db.SetUp()
	if err!=nil{
		panic(err.Error())
	}
	err = route.SetUp()
	if err!=nil{
		panic(err.Error())
	}
}
