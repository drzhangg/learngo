package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file,err := os.Open("./file/dir/demo4.go")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer file.Close()
	//获取文件大小
	fileinfo,err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	filesize := fileinfo.Size()
	//创建和文件大小一样的缓冲
	buffer := make([]byte,filesize)
	reader := bufio.NewReader(file)
	bytesread,err := reader.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("bytes read :",bytesread)
	fmt.Println("bytesstream to string:",string(buffer))
}
