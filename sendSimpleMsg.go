package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
)

func main() {
	rabbitMq := rabbitMQ.NewRabbitmqSimple("test")
	rabbitMq.PublishSimple("你好啊！")
	fmt.Println("发送完消息")
}
