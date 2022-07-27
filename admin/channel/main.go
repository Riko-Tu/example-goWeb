package main

import (
	"fmt"
	"time"
)
type Server interface {
	send() <-chan string
}

type numbServer struct {
	name string
	
}

//数字服务
func(n *numbServer) send() <-chan string{
	s := make(chan string)
	go func() {
		i := 0
		for  {
			time.Sleep(time.Second*2)
			s <- fmt.Sprintf("%s 发送消息：%d",n.name,i)
			i++
		}
	}()
	return s
}


type timeServer struct {
	name string
}


//服务二
func(t *timeServer) send() <-chan string{
	s := make(chan string)
	go func() {
		for  {
			time.Sleep(time.Second*1)
			s <- fmt.Sprintf("%s 当前时间：%s",t.name,time.Now().String())

		}
	}()
	return s
}


//多服务接收 :该写法适用于不知道业务有多少个服务
func serverReceiveFor(serverChan ...<-chan string) <-chan string {
	serverMsg := make(chan string,2)
	for _, v := range serverChan {
		go func(v <-chan string) {   //for加匿名函数时，必须传参
			for {
				serverMsg <- <-v
			}
		}(v)
	}
	return serverMsg
}

//多服务接收 ：该写法适用于明确知道服务的个数
func serverReceiveSelect(num,time <-chan string) <-chan string {
	serverMsg := make(chan string,2)
	go func() {
		for  {
			select {
			case numMsg:=<-num:
				serverMsg<-numMsg
			case timeMsg:=<-time:
				serverMsg<-timeMsg
			}
		}
	}()
	return serverMsg
}





//服务处理
func serverConsumer(serverReceive <-chan string)  {
	for  {
		fmt.Printf(<-serverReceive+"\n")
	}
}



func main() {
	/*  应用场景：将多个服务的消息发送到一个channel中统一处理
	*/

	//数字服务
	numbChan :=(&numbServer{name: "数字服务"}).send()
	//时间服务
	timeChane := (&timeServer{name: "时间服务"}).send()

	//多服务消息接收
	receive := serverReceiveSelect(timeChane,numbChan)
	//消息统一处理
	serverConsumer(receive)


}