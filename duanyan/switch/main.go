package main

import "fmt"

type Student struct {
	name string
	id int
}


func main() {
	i := make([]interface{},3)
	i[0] = 213
	i[1] = "hehehe"
	i[2] = Student{"jerry",3}

	for index,value := range i{
		switch data := value.(type) {
		case int:
			fmt.Printf("x[%d] 类型为int，内容为[%d]\n",index,data)
		case string:
			fmt.Printf("x[%d]类型为string,内容为[%s]\n",index,data)
		case Student:
			fmt.Printf("x[%d]类型为Student，内容为name = %s,id = %d\n",index,data.name,data.id)
		}
	}
}
