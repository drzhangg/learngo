package main

import (
	"fmt"
	"learngo/crud/gorm_test"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := gorm_test.DBConn()

	user := db.First(&gorm_test.User{})
	fmt.Println(user)
}
