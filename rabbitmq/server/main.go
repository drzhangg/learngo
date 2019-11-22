package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	uri = "amqp://guest:guest@localhost:5672/"

	exchangeName = ""

	queueName = "test-queue"
)

//如果存在错误，则输出
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	bodyMsg := bodyForm(os.Args)
	//调用发布消息函数
	publish(uri,exchangeName,queueName,bodyMsg)
	log.Printf("published %dB OK", len(bodyMsg))
}

func bodyForm(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello idoll.org"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

//发布者的方法
//@amqpURI,amqp的地址
//@exchange,exchange的名称
//@queue,queue的名称
//@body，主体内容
func publish(amqpURI, exchange, queue, body string) {
	var (
		err        error
		connection *amqp.Connection
		channel    *amqp.Channel
		que        amqp.Queue
		tick       *time.Ticker
	)

	log.Printf("dialing %q", amqpURI)
	connection, err = amqp.Dial(amqpURI)
	failOnError(err, "Failed to connect rabbitmq")
	defer connection.Close()

	//创建一个channel
	log.Printf("got connection,getting channel")
	channel, err = connection.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	log.Printf("got queue, declaring %q", queue)

	//创建一个queue
	que, err = channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("declared queue, publishing %dB body (%q)", len(body), body)

	tick = time.NewTicker(time.Millisecond * time.Duration(rand.Intn(1000)))
	for {
		<-tick.C
		err = channel.Publish(
			exchange,
			que.Name,
			false,
			false,
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "",
				Body:            []byte(body),
			})
	}
	failOnError(err, "Failed to publish a message")
}
