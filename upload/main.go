package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", uploadFileHandle())
	http.ListenAndServe(":8080", nil)
}

func uploadFileHandle() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, fh1, err := r.FormFile("file")
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(fh1.Filename)
	})
}
