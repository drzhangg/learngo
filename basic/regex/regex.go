package main

import (
	"fmt"
	"regexp"
)

const text = `
	my email is drzhangg@google.com
	my1 email is drzhangg3@qq.com
	my2 email is drzhangg4@def.com
	my2 email is drzhangg4@def.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	//fmt.Println(match)

	var name []string
	for _,m := range match{
		name = append(name,m[1])
		fmt.Println(m)
	}
	fmt.Println(name)

}