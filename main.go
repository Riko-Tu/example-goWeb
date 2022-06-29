package main

import "turan.com/Example-goWeb/web/route"

func main() {
	err := route.SetUp()
	if err!=nil{
		panic(err.Error())
	}
}
