package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	//使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	//构建要发送的消息
	msg := &sarama.ProducerMessage{
		//Topic:     "",
		Key:       sarama.StringEncoder("key"),
		Partition: int32(10),
	}

	var value string
	var msgType string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			break
		}
		fmt.Scanf("%s", &msgType)
		fmt.Println("msgType = ", msgType, ",value = ", value)
		msg.Topic = msgType
		//将字符串转换为字节数组
		msg.Value = sarama.ByteEncoder(value)
		partition, offset, err := producer.SendMessage(msg)

		if err != nil {
			fmt.Println("Send message Fail")
		}
		fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}
