package main

import "io/ioutil"

func check(err error) {
	if err!=nil {
		panic(err)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err:=ioutil.WriteFile("text.txt",d1,0665)
	check(err)
}
