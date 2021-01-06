package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

//创建route模式下的rabbitMq实例
func NewRouteRabbitmq(exchangeName, routeKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routeKey)
}

//发送消息
func (r *RabbitMQ) PublishRouteMsg(msg string) {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("ExchangeDeclare err:", err)
	}
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

//消费消息
func (r *RabbitMQ) ConsumeRouteMsg() {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("ExchangeDeclare err:", err)
	}
	//2.创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("QueueDeclare err:", err)
	}
	//3.绑定队列到exchange
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//4.消费消息(一直从消息里面获取)
	msgs, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//实现逻辑判断
			fmt.Println("msg:", string(d.Body))
		}
	}()
	fmt.Println("wait for msg!")
	<-forever
}
