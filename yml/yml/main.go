package main

import (
	"fmt"
	"github.com/ghodss/yaml"
)

func main() {

	j := []byte(`{"name":"Jerry","age":"20"}`)
	y,err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(y))

	j2,err := yaml.YAMLToJSON(y)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(j2))
}
