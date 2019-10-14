package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//5.15  -> 5*60 + 15
	var x  = 5.15
	ss := strconv.FormatFloat(x,'f',-1,64)
	fmt.Println(ss)

	s1:= strings.Split(ss,".")
	sint,_ := strconv.Atoi(s1[0])
	sint2,_:=strconv.Atoi(s1[1])
	sum := sint * 60 + sint2
	fmt.Println("sum---",sum)

	min := time.Now().Format("04")
	second := time.Now().Format("05")
	fmt.Println("当前分钟：",min)
	fmt.Println("当前秒钟：",second)
	fmt.Println(min+"."+second)
}
