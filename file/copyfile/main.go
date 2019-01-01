package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	list := os.Args
	if len(list) != 3 {
		fmt.Println("usage: xxx")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件和目标文件名字不能相同")
		return
	}

	//只读方式打开源文件
	sf,err1 := os.Open(srcFileName)
	if err1 != nil{
		fmt.Println("err1 = ",err1)
		return
	}

	//新建目标文件
	df,err2 := os.Create(dstFileName)
	if err2 != nil{
		fmt.Println("err2 = ",err2)
		return
	}

	defer sf.Close()
	defer df.Close()

	//核心处理，从源文件读取内容，往目标文件写，读多少写多少
	buf := make([]byte,1024 * 4)	//4k大小临时缓冲区
	for{
		n,err := sf.Read(buf)		//从源文件读取内容
		if err != nil {
			fmt.Println("err = ",err)
			if err == io.EOF {
				break
			}
		}

		df.Write(buf[:n])
	}

}
