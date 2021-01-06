package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式：amqp://账号:密码@服务器地址:端口号/vhost
const mqUrl = "amqp://imooc_user:imooc_user@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string //交换机
	Key       string //key
	MqUrl     string //连接信息
}

//创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	rabbitMq :=  &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: mqUrl}
	var err error
	//创建rabbitmq连接
	rabbitMq.conn, err = amqp.Dial(rabbitMq.MqUrl)
	rabbitMq.failOnErr(err, "创建连接错误")
	rabbitMq.channel, err = rabbitMq.conn.Channel()
	rabbitMq.failOnErr(err, "获取channel失败")
	return rabbitMq
}

//断开channel和connection
func (r *RabbitMQ) Disconnect() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,msg string) {
	if err != nil {
		log.Fatalf("%s%s",err,msg)
		panic(fmt.Sprintf("%s%s",err,msg))
	}
}