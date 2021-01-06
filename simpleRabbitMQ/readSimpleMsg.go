package main

import (
	"rabbitMQ_demo/rabbitMQ"
)

func main() {
	rabbitMq := rabbitMQ.NewSimpleRabbitmq("simple_rabbitmq")
	rabbitMq.ConsumeSimpleMsg()
}
