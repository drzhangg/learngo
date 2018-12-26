package main

import (
	"bytes"
	"fmt"
)

//定义一个函数，参数数量为0~n,类型为字符串
func joinStrings(slist ...string) string {

	//定义一个字节缓冲，快速的连接字符串
	var b bytes.Buffer

	//遍历可变参数列表slist,类型为[]string
	for _,s := range slist{
		b.WriteString(s)
	}
	return b.String()
}

func main() {
	fmt.Println(joinStrings("pig","and","rat"))
	fmt.Println(joinStrings("hammer", " mom", " and", " hawk"))
}
