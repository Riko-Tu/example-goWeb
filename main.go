package main

import (
	_ "turan/example-goWeb/admin/dataload"
	"turan/example-goWeb/admin/db"
	"turan/example-goWeb/admin/route"
)

func setUp() {
	//err := cache.SetUp()
	//if err!=nil {
	//	panic(err.Error())
	//}
	err := db.SetUp()
	if err != nil {
		panic(err.Error())
	}
	err = db.SetUpT()
	if err !=nil{
		panic(err.Error())
	}
	err = route.SetUp()
	if err != nil {
		panic(err.Error())
	}
}
func main() {
setUp()
}
