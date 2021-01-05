package main

import (
	"rabbitMQ_demo/rabbitMQ"
)

func main() {
	rabbitMq := rabbitMQ.NewRabbitmqSimple("test")
	rabbitMq.ConsumeSimple()
}
