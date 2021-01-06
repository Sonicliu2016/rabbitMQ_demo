package main

import "rabbitMQ_demo/rabbitMQ"

func main(){
	rabbitmq := rabbitMQ.NewPubSubRabbitmq("PubSubRabbitmq")
	rabbitmq.ConsumePubSubMsg()
}
