package main

import (
	"fmt"
	"rabbitMQ_demo/rabbitMQ"
	"strconv"
	"time"
)

func main() {
	rabbitMqOne := rabbitMQ.NewRouteRabbitmq("RouteRabbitmq","one")
	rabbitMqTwo := rabbitMQ.NewRouteRabbitmq("RouteRabbitmq","two")
	for i := 0; i <= 10; i++ {
		rabbitMqOne.PublishRouteMsg("Hello one!" + strconv.Itoa(i))
		rabbitMqTwo.PublishRouteMsg("Hello Two!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("我是:",i)
	}
}
