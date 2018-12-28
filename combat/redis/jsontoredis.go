package main

import (
	"encoding/json"
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

	key := "profile"
	imap := map[string]string{"username":"jerry","address":"hangzhou "}
	value,_ := json.Marshal(imap)

	n,err := c.Do("SETNX",key,value)
	if err != nil {
		fmt.Println("redis set failed:",err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var impGet map[string]string

	valueGet,err := redis.Bytes(c.Do("get",key))
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(valueGet,&impGet)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(impGet["username"])
	fmt.Println(impGet["address"])
}
