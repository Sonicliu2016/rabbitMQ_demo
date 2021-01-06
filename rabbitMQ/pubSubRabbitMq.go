package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
)

//创建订阅模式下的rabbitMq实例
func NewPubSubRabbitmq(exchangeName string) *RabbitMQ {
	return NewRabbitMQ("", exchangeName, "")
}

//生产消息
func (r *RabbitMQ) PublishPubSubMsg(msg string) {
	//1.创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",//交换机类型：广播类型
		true, //是否持久化
		false, //是否自动删除
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		nil,
	)
	if err != nil {
		fmt.Println("ExchangeDeclare err:", err)
	}
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
}

//消費消息
func (r *RabbitMQ) ConsumePubSubMsg() {
	//1.創建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//交换机类型
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println("ExchangeDeclare err:", err)
	}
	//2.创建队列，这里的队列名称不要写
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,//是否排他
		false,
		nil,
	)
	if err != nil {
		fmt.Println("QueueDeclare err:", err)
	}
	//3.绑定队列到exchange
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下，这里的key要为空
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
