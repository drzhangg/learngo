package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fileinfos,err := ioutil.ReadDir("../path")
	if err == nil {
		for _,fileinfo := range fileinfos{
			fmt.Println(fileinfo.Name(),fileinfo.IsDir(),fileinfo.Size(),fileinfo.ModTime())
		}
	}
}
