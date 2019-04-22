package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CityList struct {
	City []City	`json:"cityList"`
}

type City struct {
	Url  string `json:"linkURL"`
	Name string `json:"linkContent"`
}

func main() {

	//MapToJson()

	JsonToCity()
}

func JsonToCity(){
	file,err := ioutil.ReadFile("json/unmarshal/text.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

	var cityList []CityList

	json.Unmarshal(file,&cityList)
	fmt.Println(cityList)
}

func MapToJson() {
	file,err := ioutil.ReadFile("json/unmarshal/text.json")
	if err != nil {
		fmt.Println(err)
	}

	var dat map[string]interface{}

	json.Unmarshal(file,&dat)
	if v,ok := dat["cityList"];ok{
		cl := v.([]interface{})
		for _,city := range cl{
			clInter := city.(map[string]interface{})
			city = clInter["linkContent"]
			url:= clInter["linkURL"]
			fmt.Println("城市名字：",city)
			fmt.Println("城市URL：",url)
		}
	}

	//fmt.Println(dat)
}
