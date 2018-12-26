package main

import (
	"fmt"
	"sort"
)

func main() {

	scene := make(map[string]int)

	//准备map数据
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	//声明一个切片保存map数据
	var scenList []string

	//将map数据遍历复制到切片中
	for K := range scene{
		scenList = append(scenList,K)
	}

	//对切片进行排序
	sort.Strings(scenList)

	fmt.Println(scenList)
}
