package main

import (
	"turan/Example-goWeb/admin/route"
	"turan/example-goWeb/admin/db"
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
