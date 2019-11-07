package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	s := [...][2]string{{"a", "b"}, {"c", "d"}} //二维数组的赋值初始化
	b := s                                      //直接给数组赋值
	for _, v := range b {
		for _, t := range v {
			fmt.Print(t)
		}
	}

	//flag.String("crul","curl ifconfig.me","get ip")
	var stdout bytes.Buffer
	cmd := exec.Command("curl ifconfig.me")
	cmd.Stdout = &stdout
	cmd.Start()
	fmt.Println(stdout.String())
}
