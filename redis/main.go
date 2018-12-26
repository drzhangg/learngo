package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	c,err := redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,err:",err)
		return
	}
	defer c.Close()

	//_,err = c.Do("Set","name","nick")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//r,err := redis.String(c.Do("Get","name"))
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(r)

	//_,err = c.Do("MSet","name","nick","age","12")
	//if err != nil {
	//	fmt.Println("MSet error:",err)
	//	return
	//}
	//
	//r2,err := redis.String(c.Do("MGet","name","age"))
	//if err != nil {
	//	fmt.Println("MGet error:",err)
	//	return
	//}
	//fmt.Println(r2)

	_,err = c.Do("HSet","names","nick","suoning")
	if err != nil {
		fmt.Println("hset error:",err)
		return
	}

	r,err := redis.String(c.Do("HGet","names","nick"))
	if err != nil {
		fmt.Println("hget errror:",err)
		return
	}
	fmt.Println(r)

	//设置过期时间
	_,err =  c.Do("expire","names",5)
	if err != nil {
		fmt.Println("expire error:",err)
		return
	}

	//队列
	_,err = c.Do("lpush","Queue","nick","dawn",9)
	if err != nil {
		fmt.Println("lpush error:",err)
		return
	}

	for{
		r,err = redis.String(c.Do("lpop","Queue"))
		if err != nil {
			fmt.Println("lpop error:",err)
			break
		}
		fmt.Println(r)
	}
	r3,err:=redis.Int(c.Do("llen","Queue"))
	if err != nil {
		fmt.Println("llen error:",err)
		return
	}
}
