package demo2

import (
	"fmt"
	"github.com/streadway/amqp"
	"sync"
)

//定义q全局变量，指针类型
var mqConn *amqp.Connection
var mqChan *amqp.Channel

//定义生产者j接口
type Producer interface {
	MsgContent() string
}

//定义接收者接口
type Receiver interface {
	Consumer([]byte) error
}

//定义RabbitMQ对象
type RabbitMQ struct {
	connection   *amqp.Connection
	channel      *amqp.Channel
	queueName    string //队列名称
	routingKey   string //key名称
	exchangeName string //交换机名称
	exchangeType string //交换机类型
	producerList []Producer
	receiver     []Receiver
	mu           sync.RWMutex
}

//定义队列交换机对象
type QueueExchange struct {
	QuName string //队列名称
	RtKey  string //key值
	ExName string //交换机名称
	ExType string //交换机类型
}

//连接rabbitmq
func (r *RabbitMQ) mqConnect() {
	var err error
	RabbitUrl := fmt.Sprintf("amqp://%s:%s@%s:%d/", "guest", "guest", "localhost", 5672)
	mqConn, err = amqp.Dial(RabbitUrl)
	if err != nil {
		fmt.Printf("MQ打开链接失败:%s \n", err)
	}
	mqChan, err = mqConn.Channel()
	if err != nil {
		fmt.Printf("MQ打开管道失败:%s \n", err)
	}
	r.connection = mqConn
	r.channel = mqChan
}
