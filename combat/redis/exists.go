package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis failed",err)
		return
	}
	defer c.Close()

	_,err = c.Do("set","age",14)
	if err != nil {
		fmt.Println("redis set failed:",err)
	}

	age,err := redis.Bool(c.Do("exists","age1"))
	if err != nil {
		fmt.Println("error",err)
		return
	}

	fmt.Printf("exists or not :%v\n",age)
}
