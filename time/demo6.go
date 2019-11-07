package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	next := now.Add(time.Hour * 24)
	year, month, day, location := next.Year(), next.Month(), next.Day(), next.Location()
	next = time.Date(year, month, day, 8, 0, 0, 0, location)

	if now.String() > next.String() {

	}
	hh,_ := time.ParseDuration("1h")
	hhhh := now.Add(hh*8)
	fmt.Println(hhhh)
	fmt.Println(next)
	fmt.Println(next.Sub(hhhh))
	fmt.Println(now.String() < next.String())



}
