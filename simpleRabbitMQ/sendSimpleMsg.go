package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
)

func main() {
	rabbitMq := rabbitMQ.NewSimpleRabbitmq("simple_rabbitmq")
	rabbitMq.PublishSimpleMsg("你好啊！")
	fmt.Println("发送完消息")
}
