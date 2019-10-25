package main

import (
	"fmt"
	"time"
)

func main() {
	//ss := "1970-01-01 08:00:00"
	ss := "2019-10-23 21:53:28"
	tt,_ := time.ParseInLocation("2006-01-02 15:04:05",ss,time.Local)
	fmt.Println(tt.UnixNano())
	fmt.Println(tt)

	//28800000000000
	//1568969447000
}
