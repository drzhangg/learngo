package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	//设置处理函数
	http.HandleFunc("/",TestAction)
	//开启监听
	log.Fatal(http.ListenAndServe(":8084",nil))
}

func TestAction(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t,_ := template.ParseFiles("template/index.html")
	//设置模板数据
	data := map[string]interface{}{
		"User":"drzhang",
		"List":[]string{"GO","JAVA","PHP","Python"},
	}

	t.Execute(w,data)
}
