package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitmq := rabbitMQ.NewPubSubRabbitmq("PubSubRabbitmq")
	for i:=1;i<=100;i++{
		fmt.Println("生产第" + strconv.Itoa(i) + "条消息")
		rabbitmq.PublishPubSubMsg("我是第" + strconv.Itoa(i) + "条消息")
		time.Sleep(1 * time.Second)
	}
}
