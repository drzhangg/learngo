package main

import "fmt"

type user struct {
	Name string
}


func main() {

	userChan := make(chan interface{},1)

	u:= user{"nick"}
	userChan <- &u
	close(userChan)

	var u1 interface{}
	u1 = <- userChan

	var u2 *user
	u2,ok := u1.(*user)
	if !ok{
		fmt.Println("cant not convert.")
		return
	}
	fmt.Println(u2)
}
