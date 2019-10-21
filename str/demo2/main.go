package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s := "30m10s"
	ss := strings.Replace(s, "h", "小时", -1)
	sss := strings.Replace(ss,"m","分钟",-1)

	t := time.Unix(1571495639,0)
	tt := t.Format("2006-01-02 15:04:05")
	fmt.Println(tt)


	fmt.Println(sss)
}
