package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	//send向连接的输出缓冲中写入命令
	c.Send("SET","name1","sss1")
	c.Send("SET","name2","sss2")

	//flush将连接的输出缓冲清空并写入服务器端
	c.Flush()

	//receive按照FIFO（先进先出）顺行依次读取服务器的响应
	v,err:= c.Receive()
	fmt.Printf("v:%v,err:%v\n",v,err)
	v,err=c.Receive()
	fmt.Printf("v:%v,err:%v\n",v,err)

	v,err = c.Receive()
	fmt.Printf("v:%v,err:%v\n",v,err)
}
