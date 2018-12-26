package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",set)
	http.HandleFunc("/read",read)
	http.ListenAndServe(":8080",nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{
		Name:"my-cookie",
		Value:"some cookie value",
	})
	fmt.Fprintln(w,"Cookie written - check you browser")
}

func read(w http.ResponseWriter, r *http.Request) {
	c,err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w,http.StatusText(400),http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w,"you cookie:",c)
}
