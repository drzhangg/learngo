package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect redis failed:",err)
		return
	}
	defer c.Close()

	_,err = c.Do("SET","money","1w","EX",3)
	if err != nil {
		fmt.Println("redis set failed:",err)
	}

	val,err := redis.String(c.Do("GET","money"))
	if err != nil {
		fmt.Println("redis get failed:",err)
		return
	}
	fmt.Printf("Get redis value:%v\n",val)

	time.Sleep(5 * time.Second)

	val,err = redis.String(c.Do("GET","money"))
	if err != nil {
		fmt.Println("redis get failed:",err)
	}
}
