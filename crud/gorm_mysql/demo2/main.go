package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	ID        int    `gorm:"primary_key"`
	Code      string `gorm:"type:varchar(20);"`
	Price     int    `gorm:"type:int;"`
	Name      string `gorm:"type:varchar(64);"`
	Mail      string `gorm:"type:varchar(256);"`
	CreatedAt time.Time
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}

	if !db.HasTable(&Product{}) {
		db.CreateTable(&Product{})
	}

	var product Product
	//db.SingularTable(true)
	//db.Create(&Product{Code: "ik01001", Price: 1000}).SingularTable(true)

	//db.Where("id=?",1).Model(&product)
	//db.First(&product,1)
	db.Find(&product)
	fmt.Println(product)
}
