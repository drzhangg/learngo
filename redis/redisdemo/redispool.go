package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool


func main() {
	pool = &redis.Pool{
		MaxIdle:8,//最大空闲连接数
		MaxActive:0,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:100,//最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","localhost:6379")
		},
	}

	c := pool.Get()
	defer c.Close()

	_,err := c.Do("set","sex","man")
	if err != nil {
		fmt.Println("redis do failed,err = ",err)
		return
	}

	r,err := redis.String(c.Do("get","sex"))
	if err != nil {
		fmt.Println("err = ",err)
		return
	}
	fmt.Println(r)


}
