package main

import (
	"fmt"
	"time"
)

func main() {

	data := time.Now().Format("2006-01-02 15:04:05")
	data1 := time.Now().Format("2006-01-02 ")
	data1 = data1 + "12:00:00"
	fmt.Println(data)
	//fmt.Println(data1)

	fmt.Println(data > data1)

	now := time.Now()
	next := now.Add(time.Hour * 24)
	nn := next.Format("2006-01-02")
	nn1 := nn + " 08:00:00"
	fmt.Println(nn1)
	//fmt.Println(data > nn1)

	time1,_ := time.Parse("2006-01-02 15:04:05",nn1)
	fmt.Println(time1.Sub(now))


	year,month,day,location := next.Year(),next.Month(),next.Day(),next.Location()
	next = time.Date(year,month,day,0,0,0,0,location)
	t := time.NewTimer(next.Sub(now))
	fmt.Println(&t)

}
