package main

import (
	"fmt"
	"time"
)

func main() {

	str := "2019-01-02 15:04:05"
	str2 := "2019-01-02 15:04:06"
	ii := TimeToUnix(str)
	ii2 := TimeToUnix(str2)
	fmt.Println(ii)
	fmt.Println(ii2)
}

func TimeToUnix(timeStr string) int64 {
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, loc)
	sr := theTime.Unix()
	return sr
}
