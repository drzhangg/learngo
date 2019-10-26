package main

import (
	"fmt"
	"time"
)

func main() {
	s := "2019-10-23 21:53:28"

	times,_ :=time.Parse("2006-01-02 15:04:05",s)
	times1:=times.Format(s)
	time.ParseInLocation("2006-01-02 15:04:05",s,time.Local)

	dd,_:=time.ParseDuration("48h")
	d := times.Add(dd)
	fmt.Println(times)
	fmt.Println(times1)
	fmt.Println(d)
}
