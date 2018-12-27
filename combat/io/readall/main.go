package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	link := "http://xkcd.com/55"

	tr := &http.Transport{
		TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	}
	client := &http.Client{Transport:tr}
	response,err:=client.Get(link)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content,_:=ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
