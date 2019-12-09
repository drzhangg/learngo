package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	s := `2019-12-06 15:57:52`
	s1 := `2019-12-09 15:57:50`
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", s1, time.Local)
	now := t2.Sub(t1).String()

	//now := time.Now().Sub(t1).String()
	ss := strings.Split(now, ".")

	hs := strings.Split(ss[0], "h")
	if len(hs) == 2 {
		ms := strings.Split(hs[1], "m")
		second := strings.Split(ms[1],"s")
		hint, _ := strconv.Atoi(hs[0])
		if hint%24 == 0 {
			d := hint / 24
			//ms := hs[1]
			fmt.Println(strconv.Itoa(d) + "天" + ms[0] + "分" + second[0] + "秒")
		} else {
			d := hint % 24
			fmt.Println(strconv.Itoa(hint/24) + "天" + strconv.Itoa(d) + "小时" + ms[0] + "分" + second[0] + "秒")
		}
	} else {
		ms := strings.Split(hs[0], "m")
		second := strings.Split(ms[1],"s")
		if len(ms) == 2 {
			fmt.Println(ms[0] + "分" + second[0] + "秒")
		} else {
			ss := strings.Split(ms[0], "s")
			fmt.Println(ss[0] + "秒")
		}
	}

}
