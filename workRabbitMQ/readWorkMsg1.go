package main

import "rabbitMQ_demo/rabbitMQ"

func main(){
	rabbitMq := rabbitMQ.NewSimpleRabbitmq("work_rabbitmq")
	rabbitMq.ConsumeSimpleMsg()
}

