package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int64
	Username string
	Password string
}

func (User) TableName() string {
	return "userr"
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	//创建表
	db.CreateTable(&User{})

	/*user := User{Username:"root",Password:"root"}
	result := db.Create(&user)

	if result.Error != nil {
		fmt.Printf("insert row err %v",result.Error)
		return
	}
	fmt.Println(user.ID)*/

	/*getUser := User{}
	db.Select([]string{"id","username","password"}).First(&getUser,1)
	fmt.Println(getUser)*/

	/*user := User{}
	user.Username = "update username"
	user.Password = "update password"
	db.Save(&user)*/

	users := []User{}

	db.Find(&users)
	fmt.Println(&users)

}
