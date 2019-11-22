package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var (
	conn    *amqp.Connection
	channel *amqp.Channel
)

var count = 0

const (
	queue    = "push.msg.q"
	exchange = "t.msg.x"
	mqurl    = "amqp://guest:guest@localhost:5672/"
)

func main() {
	go func() {
		for {
			push()
			time.Sleep(1 * time.Second)
		}
	}()
	receive()
	fmt.Printf("end")
	close()
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func mqConnect() {
	var err error
	conn, err := amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func close() {
	channel.Close()
	conn.Close()
}

func push() {
	if channel == nil {
		mqConnect()
	}

	msgContent := "hello world!"

	channel.Publish(exchange, queue, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})
}

func receive() {
	if channel == nil {
		mqConnect()
	}

	msgs, err := channel.Consume(queue, "", true, false, false, false, nil)
	failOnErr(err, "")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			s := BytesToString(&(d.Body))
			count++
			fmt.Printf("receve msg is :%s -- %d\n", *s, count)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
