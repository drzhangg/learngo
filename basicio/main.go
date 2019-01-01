package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte,num)

	n,err := reader.Read(p)
	if n > 0 {
		return p[:n],nil
	}
	return p,err
}

func sampleReadFromString() {
	data,_ := ReadFrom(strings.NewReader("from string"),12)

	fmt.Println(data)
}

func main() {
	///sampleReadFromString()

	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)

	data,_ := bufReader.Peek(7)

	fmt.Println(string(data),bufReader.Buffered())

	str,_ := bufReader.ReadString(' ')
	fmt.Println(str,bufReader.Buffered())
}
