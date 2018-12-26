package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func err_handler(err error) {
	//fmt.Printf("err_handler,error:%s\n",err.Error())
	if err != nil {
		panic(err)
	}
}

func test_delete_db() {
	fmt.Println("test delete db")

	sql_uri := "root:root@(127.0.0.1:3306)/test?charset=utf8"
	db,err := sql.Open("mysql",sql_uri)

	err_handler(err)
	defer db.Close()

	var delete_sql string = "delete from test_userinfo"
	_,err = db.Exec(delete_sql)
	if err != nil {
		fmt.Printf("try do sql[%s] error[%s]\n",delete_sql)
		err_handler(err)
	}
}

func insert() {
	db,err := sql.Open("mysql","root:root@/test?charset=utf8")
	err_handler(err)

	stmt,err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	err_handler(err)
	res,err := stmt.Exec("Tom",22,"男")
	err_handler(err)
	id,err := res.LastInsertId()
	err_handler(err)
	fmt.Println(id)
}

func query() {
	db,err := sql.Open("mysql","root:root@(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	rows,err := db.Query("SELECT * from user ")
	if err != nil {
		panic(err)
	}

	//字典类型
	//构造scanArgs、values两个数组scanArgs的每个值指向values相应值得地址
	columns,_ := rows.Columns()
	scanArgs := make([]interface{},len(columns))
	values := make([]interface{},len(columns))
	for i := range values{
		scanArgs[i] = &values[i]
	}

	for rows.Next(){
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i,col := range values{
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

func update() {
	db,err := sql.Open("mysql","root:root@(127.0.0.1:3306)/test?charset=utf8")
	err_handler(err)

	stmt,err := db.Prepare("update user set user_name = ?,user_sex=? where user_id=?")
	err_handler(err)
	res,err := stmt.Exec("张三","女",1)
	err_handler(err)
	num,err := res.RowsAffected()
	err_handler(err)
	fmt.Println(num)
}

func main() {
	//insert()
	//update()
	query()
}
