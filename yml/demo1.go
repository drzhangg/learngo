package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Test struct {
	Test1 int    `yaml:"Test1"`
	Test2 string `yaml:"Test2"`
}

func main() {
	var yamls Test
	byes, err := ioutil.ReadFile("server.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byes))
	err = yaml.Unmarshal(byes, &yamls)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(yamls)
}
