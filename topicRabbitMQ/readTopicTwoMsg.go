package main

import "rabbitMQ_demo/rabbitMQ"

func main() {
	rabbitMqTwo := rabbitMQ.NewTopicRabbitmq("TopicRabbitmq","test.*.two")
	rabbitMqTwo.ConsumeTopicMsg()
}
