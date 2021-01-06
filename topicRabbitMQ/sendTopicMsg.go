package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMqOne := rabbitMQ.NewTopicRabbitmq("TopicRabbitmq","test.topic.one")
	rabbitMqTwo := rabbitMQ.NewTopicRabbitmq("TopicRabbitmq","test.topic.two")
	for i := 0; i <= 10; i++ {
		rabbitMqOne.PublishTopicMsg("Hello one!" + strconv.Itoa(i))
		rabbitMqTwo.PublishTopicMsg("Hello Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("我是:",i)
	}
}
