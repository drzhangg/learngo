package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	//a := make([]string,100)
	//b := a[:]
	//d := [...]string{}
	c := []string{}

	filepath.Walk("D:\\code\\java\\demo\\go\\src\\learngo", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {

			c = append(c,path)
			//fmt.Println(c)
			//fmt.Println(len(c))
		}
		return err
	})
	fmt.Println(c)

	for _,value := range c{
		fmt.Println(value)
	}

	fmt.Println(len(c))
}
