package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis dial failed,err = ",err)
	}

	defer conn.Close()
	//fmt.Println("conn success，conn = ",conn)

	res,err := redis.Strings(conn.Do("smembers","emails"))
	fmt.Println(res)

}
