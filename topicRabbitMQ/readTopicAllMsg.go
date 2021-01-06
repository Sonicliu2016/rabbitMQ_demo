package main

import "rabbitMQ_demo/rabbitMQ"

func main() {
	rabbitMqAll := rabbitMQ.NewTopicRabbitmq("TopicRabbitmq","#")
	rabbitMqAll.ConsumeTopicMsg()
}
