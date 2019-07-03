package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial failed ,err = ", err)
	}
	defer conn.Close()

	_, err = conn.Do("hmset", "user1", "name", "drzhang", "age", 20)
	if err != nil {
		fmt.Println("redis do failed,err = ", err)
	}

	res, err := redis.Strings(conn.Do("hmget", "user1", "name", "age"))
	if err != nil {
		fmt.Println("get data, err = ", err)
	}
	fmt.Println(res)

	for i, v := range res {
		fmt.Println(i, v)
	}

}
