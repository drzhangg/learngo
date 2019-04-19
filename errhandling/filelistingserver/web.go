package main

import (
	"log"
	"learngo/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter,*http.Request)  {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			log.Printf("Panic:%v",r)
			http.Error(writer,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
		}()

		err := handler(writer,request)
		if err != nil {
			log.Printf("Error handling request: %s",err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				//没有权限
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888",nil)
	if err != nil {
		panic(err)
	}
}
