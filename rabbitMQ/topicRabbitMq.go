package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

//创建topic模式下的rabbitMq实例
func NewTopicRabbitmq(exchangeName, routeKey string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, routeKey)
}

//发送消息
func (r *RabbitMQ) PublishTopicMsg(msg string) {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
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

//话题模式接受消息
//要注意key,规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 imooc.* 表示匹配 imooc.hello, 但是imooc.hello.one需要用imooc.#才能匹配到
func (r *RabbitMQ) ConsumeTopicMsg() {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic",
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
		//在pub/sub模式下，这里的key要为空
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
