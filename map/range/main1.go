package main

import "fmt"

func main() {
	var typeMap = make(map[string]string, 6)

	var abnMaps = []struct {
		Key string `json:"key"`
		Val string `json:"val"`
	}{}
	var abnMap = struct {
		Key string `json:"key"`
		Val string `json:"val"`
	}{}

	/**
	1.受理时限异常
	2.审批时限异常
	3.办结时限异常
	4.咨询受理时限异常
	5.回复咨询时限异常
	6.其他时限异常
	*/
	typeMap["1"] = "受理时限异常"
	typeMap["2"] = "审批时限异常"
	typeMap["3"] = "办结时限异常"
	typeMap["4"] = "咨询受理时限异常"
	typeMap["5"] = "回复咨询时限异常"
	typeMap["6"] = "其他时限异常"

	for k, v := range typeMap {
		fmt.Println(k, v)
		abnMap.Key = string(k)
		abnMap.Val = v
		abnMaps = append(abnMaps, abnMap)
	}

	fmt.Println(abnMaps)
}
