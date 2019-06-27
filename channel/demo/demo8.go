package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i,v := range values{
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf,"%d",v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	var runes []rune
	for _,v := range "hello, 世界！"{
		runes = append(runes,v)
	}
	fmt.Printf("%q ",runes)
	fmt.Println(intsToString([]int{2,3,4,5}))
}
