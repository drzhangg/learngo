package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response,err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
