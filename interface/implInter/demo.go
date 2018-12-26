package main

import "fmt"

//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}
//定义文件结构,用于实现DataWrite
type file struct {
}

//实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {

	//模拟写入数据
	fmt.Println("WriteData:",data)
	return nil
}

func main() {

	f := new(file)

	var writer DataWriter

	writer = f
	writer.WriteData("data")
}
