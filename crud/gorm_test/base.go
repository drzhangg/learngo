package gorm_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name string
	Age  int
	Addr string
}

func DBConn() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()

	hasTab := db.HasTable("user")
	if !hasTab {
		panic("dont has table")
	}

	return db
}
