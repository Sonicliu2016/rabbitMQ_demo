package main

import "rabbitMQ_demo/rabbitMQ"

func main() {
	rabbitMq := rabbitMQ.NewRouteRabbitmq("RouteRabbitmq","one")
	rabbitMq.ConsumeRouteMsg()
}

