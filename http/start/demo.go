package main

import (
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/hello",IndexActive)
	log.Fatal(http.ListenAndServe(":8888",nil))
}
func IndexActive(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1 align="center">hello go</h1>`))
	w.Write([]byte(`<h1 align="center">来自小韩说课的问候</h1>`))
}

func test(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Cookies())
	c,_:= r.Cookie("User")

	log.Println(c.Name)

	//设置响应头
	w.Header().Set("Content-Type","application/json")
}
