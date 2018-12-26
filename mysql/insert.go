package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId int	`db:"user_id"`
	Username string	`db:"username"`
	Sex string	`db:"sex"`
	Email string	`db:email`
}

var Db *sqlx.DB

func init() {
	database,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,",err)
		return
	}
	Db = database
}

func main() {
	r,err := Db.Exec("insert into person(username,sex,email) values (?,?,?)","suoning","man","suoning@qq.com")
	if err != nil {
		fmt.Println("exec failed,",err)
		return
	}
	id,err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed,",err)
		return
	}
	fmt.Println("insert succ:",id)

	update()
}

func update() {
	_,err := Db.Exec("update person set user_id=? where username=?",111,"suoning")
	if err != nil {
		fmt.Println(err)
		return
	}
}


