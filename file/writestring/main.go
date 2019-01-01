package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteString(path string) {
	//1.开启连接，创建文件
	f,err := os.Create(path)
	if err != nil {
		fmt.Println("err : ",err)
		return
	}

	//2.关闭连接
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {

		buf = fmt.Sprintf("i = %d\n",i)
		fmt.Println("buf : ",buf)

		n,err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err:",err)
		}
		fmt.Println("n:",n)
	}
}

func ReadFile(path string) {
	f,err := os.Open(path)
	if err != nil{
		fmt.Println("err=",err)
		return
	}

	defer f.Close()

	buf := make([]byte,1024*2)

	//n代表从文件读取内容的长度
	n,err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF {	//文件出错，同时没有到结尾
		fmt.Println("err1=",err1)
		return
	}
	fmt.Println("buf = ",string(buf[:n]))
}

func ReadFileLine(path string) {
	f,err := os.Open(path)
	if err != nil {
		fmt.Println("err = ",err)
		return
	}

	defer f.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(f)

	for{
		//遇到'\n'结束读取，但是'\n'也读取进入
		buf,err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {	//文件结束
				break
			}
			fmt.Println("err:",err)
		}
		fmt.Printf("buf = #%s#\n",string(buf))
	}
}

func main() {
	path := "./demo.txt"

	WriteString(path)
	//ReadFile(path)
	//ReadFileLine(path)
}
