package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

//创建简单模式下rabbitMq实例
func NewSimpleRabbitmq(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

//发送消息代码
func (r *RabbitMQ) PublishSimpleMsg(msg string) {
	//1.申请队列，如果队列不存在，则会自动创建，如果存在，则跳过创建，保证队列存在，消息能发到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//消息是否持久化(重启后消息没了)
		false,
		//是否自动删除
		false,
		//是否具有排他性（自己可见，其他用户看不到）
		false,
		//是否阻塞（发送消息后，是否等待服务器响应）
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println("QueueDeclare err:", err)
	}
	//2.发送消息到队列
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，会根据Exchange类型和routkey规则，如果无法找到符合条件的队列，则会把消息返还给发送者
		false,
		//如果为true，当Exchange发送消息到队列后发现队列上没有绑定消费者，那么会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

//消费消息
func (r *RabbitMQ) ConsumeSimpleMsg() {
	//1.申请队列，如果队列不存在，则会自动创建，如果存在，则跳过创建，保证队列存在，消息能发到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//消息是否持久化(重启后消息没了)
		false,
		//是否自动删除
		false,
		//是否具有排他性（自己可见，其他用户看不到）
		false,
		//是否阻塞（发送消息后，是否等待服务器响应）
		false,
		//额外属性
		nil,
	)
	if err != nil {
		fmt.Println("err:", err)
	}
	//接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否告诉服务器已经消费了消息
		true,
		//是否具有排他性（自己可见，其他用户看不到）
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//是否阻塞（发送消息后，是否等待服务器响应）
		false,
		nil,
	)
	if err != nil {
		fmt.Println("err:", err)
	}
	//消费消息(一直从消息里面获取)
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
