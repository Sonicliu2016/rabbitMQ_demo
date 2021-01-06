package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMq := rabbitMQ.NewSimpleRabbitmq("work_rabbitmq")
	for i:=0;i<=100;i++{
		fmt.Println("发送第:",i)
		rabbitMq.PublishSimpleMsg("我是第" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
