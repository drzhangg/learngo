package main

import (
	//_ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)


func main() {

	var a = make([]int,5)
	fmt.Println(len(a),cap(a))

	b := []int{1,2,3,4,5}
	a = append(a,b...)
	fmt.Println(len(a),cap(a))

}
