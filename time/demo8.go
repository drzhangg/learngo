package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var subTime time.Duration
	now := time.Now()
	one := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	two := time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, now.Location())
	three := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	four := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, now.Location())
	last := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	if now.String() > one.String() && now.String() < two.String() {
		subTime = two.Sub(now)
		fmt.Println(1)
		fmt.Println(subTime)

	} else if now.String() > two.String() && now.String() < three.String() {
		//subTime = three.Sub(now)

		h, _ := time.ParseDuration("1h")
		h1 := now.Add(h * 8)
		subTime = last.Sub(h1)
		fmt.Println(2)
		fmt.Println(subTime)
		fmt.Println(reflect.TypeOf(subTime))
	} else if now.String() > three.String() && now.String() < four.String() {
		subTime = four.Sub(now)
		fmt.Println(3)
		fmt.Println(subTime)

	} else if now.String() > four.String() && now.String() < last.String() {
		h, _ := time.ParseDuration("1h")
		h1 := now.Add(h * 8)
		subTime = last.Sub(h1)
		fmt.Println(4)
		fmt.Println(subTime)

	}
}
