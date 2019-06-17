package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
	UserId   int    `gorm:"primary_key;column:user_id"`
	UserName string `gorm:"column:username"`
	Sex      string `gorm:"column:sex"`
	Email    string `gorm:"column:email"`
}

func main() {
	var (
		db     *gorm.DB
		err    error
		person Person
	)

	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("open failed:", err)
		return
	}
	db.AutoMigrate(&person)
	fmt.Println("open success:", db)

	//if !db.HasTable(person) {
	//db.Create(&person)
	//}

	//InsertData(db)

	person = Person{UserName: "stu001"}

	db = db.Find(&person)

	fmt.Println(person.Email)
}
