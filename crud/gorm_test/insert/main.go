package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string `gorm:"primary_key""`
	Age  int
	Addr string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("open db failed,", err)
		return
	}

	if !db.HasTable("user"){
		db.CreateTable(&User{})
	}else {
		user := &User{
			Name:"drzhangg",
			Age:23,
			Addr:"hangzhou",
		}

		db.Create(user)
	}



	//fmt.Println("user name:",user.Name)
}
