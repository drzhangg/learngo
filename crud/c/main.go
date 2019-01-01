package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var(
	id int
	name string
)

func main() {

	db ,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	//insert(db)
	//query(db)
	//update(db)
	query(db)
	//delete(db)
	
}

func insert(db *sql.DB)  {
	result,err := db.Prepare("insert into test (id,name) values (?,?)")
	if err != nil {
		fmt.Println("数据库插入出错",err)
		return
	}
	defer result.Close()

	result.Exec(1,"tom")
	result.Exec(2,"jerry")
}

func query(db *sql.DB) {
	result,err := db.Query("select * from test")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer result.Close()

	for result.Next() {

		//按表名的顺序读取参数
		err := result.Scan(&id,&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id,name)
	}

}

func update(db *sql.DB) {
	result,err := db.Prepare("update test set name = ? where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer result.Close()

	result.Exec("bob",1)
}

func delete(db *sql.DB) {
	result,err := db.Prepare("delete from test where id = ?")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer result.Close()

	count,_ := result.Exec(2)

	num,_ := count.RowsAffected()
	fmt.Printf("删除了%v个\n",num)
}
