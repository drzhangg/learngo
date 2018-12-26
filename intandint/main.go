package main

import "io"

//声明一个设备结构
type device struct {
}

//实现io.Writer的Writer()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0,nil
}

//实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}

func main() {

	var wc io.WriteCloser = new(device)

	wc.Write(nil)

	wc.Close()

	var writeOnly io.Writer = new(device)

	writeOnly.Write(nil)
}

