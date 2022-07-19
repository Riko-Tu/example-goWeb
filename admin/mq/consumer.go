package mq

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"turan/example-goWeb/admin/utils"

	"github.com/nsqio/go-nsq"
)

type myMessageHandler struct {}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	var emailAndCode utils.Email
	err := json.Unmarshal(m.Body, &emailAndCode)
	if err !=nil {
		return err
	}
	err = utils.SendEmail(emailAndCode)
	if err != nil {
		return err
	}
	log.Println(fmt.Sprintf("email : %s send success",emailAndCode.Email))
	return nil
}


func InitConsumer() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	decodeConfig := nsq.NewConfig()
	c, err := nsq.NewConsumer("email", "Channel", decodeConfig)
	if err != nil {
		log.Panic("Could not create consumer")
	}
	//c.MaxInFlight defaults to 1

	c.AddHandler(&myMessageHandler{})

	err = c.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		log.Panic("Could not connect")
	}
	log.Println("Awaiting messages from NSQ topic \"My NSQ Topic\"...")
	wg.Wait()
}
