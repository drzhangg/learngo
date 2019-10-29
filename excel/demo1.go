package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	var fileDatas []*File
	var fileData *File
	filepath.Walk("./excel/file",
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				//fmt.Println("dir:", path)
				return nil
			}
			//fmt.Println("file:", path)
			bytes, _ := ioutil.ReadFile(path)
			json.Unmarshal(bytes, &fileData)
			//fmt.Println(string(bytes))
			fileDatas = append(fileDatas, fileData)
			fmt.Println(fileData)
			fmt.Println(len(fileDatas))

			//for k,v := range fileData.Item{
			//	fmt.Println(k,v)
			//}

			return nil
		})

}

