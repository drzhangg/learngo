package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//type User struct {
//	ID int	`gorm:"primary_key"`
//	Name string `gorm:"column:name;not_null"`
//}

type Products struct {
	gorm.Model
	Code        string
	Price       int
	ProductName string
}

var (
	db  *gorm.DB
	err error
)

func main() {
	db, err = gorm.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("connection success")
	}
	defer db.Close()

	//db.Create(Products{Code:"codeddd",Price:5,ProductName:"product_name"})

	//db.SingularTable(true)
	products := []Products{}
	db.Find(&products)
	//db.Select([]string{"id","code","product_name"}).First(&products,1)
	////db.Model(&products).Where("id = ?",1)
	//fmt.Println("products:",products)
	//
	fmt.Println(products)
	/*for _,p := range products{
		fmt.Println(p.Code)
	}*/

}




