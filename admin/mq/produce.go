package mq

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"turan/example-goWeb/admin/utils"
)

// NSQ Producer Demo

var producer *nsq.Producer

// 初始化生产者

func GetProduce() *nsq.Producer {
	return producer
}
func EmailQueue(email utils.Email) error {
	emailJson, _ := json.Marshal(email)

	err := producer.Publish("My_NSQ_Topic",emailJson)
	return err
}
func InitProducer() error {

	nsqdAddr := viper.GetString("mq.d")
	config := nsq.NewConfig()
	var  err error
	producer, err = nsq.NewProducer(nsqdAddr, config)
	return err
}

