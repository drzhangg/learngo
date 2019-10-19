package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "30m10s"
	ss := strings.Replace(s, "h", "小时", -1)
	sss := strings.Replace(ss,"m","分钟",-1)
	fmt.Println(sss)
}
