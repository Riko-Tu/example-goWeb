package main

import (
	"turan/example-goWeb/admin/cache"
	_ "turan/example-goWeb/admin/dataload"
	"turan/example-goWeb/admin/db"
	"turan/example-goWeb/admin/mq"
	"turan/example-goWeb/admin/route"

)

func setUp() {
	err := cache.SetUp()
	if err!=nil {
		panic(err.Error())
	}
	go mq.InitConsumer()
	err = mq.InitProducer()
	if err != nil {
		panic(err.Error())
	}
	defer mq.GetProduce().Stop()
	err = db.SetUp()
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
