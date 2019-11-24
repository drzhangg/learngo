package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

const (
	queueName = "hello"
	exchange  = "exchange1"
	mqurl     = "amqp://guest:guest@localhost:5672/"
)

func main() {
	push()
	receive()
	fmt.Println("end")
	close()
}

//错误处理
func failOnErr(err error, msg string) {
	if err != nil {
		fmt.Println("msg", err, msg)
		panic("error")
	}
}

//建立连接
func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "connect")

	channel, err = conn.Channel()
	failOnErr(err, "channel")
}

//断开连接
func close() {
	conn.Close()
	channel.Close()
}

//声明交换器、消息队列并发送消息
func push() {
	if channel == nil {
		mqConnect()
	}

	msgConnect := "hello world"
	err := channel.ExchangeDeclare(exchange, "direct", false, false, false, false, nil)
	failOnErr(err, "ExchangeDeclare")

	_, err = channel.QueueDeclare(queueName, false, false, false, false, nil)
	failOnErr(err, "QueueDeclare")

	err = channel.QueueBind(queueName, "info", exchange, false, nil)
	failOnErr(err, "QueueBind")

	err = channel.Publish(exchange, "info", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgConnect),
	})
	failOnErr(err, "Publish")

	fmt.Println("push ok")
}

//消费者消息
func receive() {
	if channel == nil {
		mqConnect()
	}

	msg, ok, err := channel.Get(queueName, false)
	failOnErr(err, "")
	if !ok {
		fmt.Println("do not get msg")
		return
	}

	err = channel.Ack(msg.DeliveryTag, false)
	failOnErr(err, "")

	s := bytes.NewBuffer(msg.Body).String()

	//s := BytesToString(&(msg.Body))
	fmt.Printf("receve msg is :%s\n", s)
}
func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
