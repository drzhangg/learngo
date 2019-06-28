package main

import "fmt"

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)

	m1 := make(map[string]string)
	m1["drzhang"] = "失业"
	m1["golang"] = "半年经验"

	mapChan <- m1

	m2 := make(map[string]string,10)
	m2["address"] = "即将要去广州"
	m2["salary"] = "能有公司收留就行"

	mapChan <- m2

	for v := range mapChan{
		fmt.Println(v)
		if len(mapChan) == 0 {
			break
		}
	}
	//fmt.Println(mapChan)
}
