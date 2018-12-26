package main

import "fmt"

//车轮
type Wheel struct {
	Size int
}

//引擎
type Engine struct {
	Power int //功率
	Type string	//类型
}

//car
type Car struct {
	Wheel
	Engine
}

func main() {

	c := Car{
		//初始化轮子
		Wheel{
			Size:18,
		},
		Engine{
			Type:"1.4T",
			Power:143,
		},
	}

	fmt.Printf("%+v",c)
}

