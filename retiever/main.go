package main

import (
	"fmt"
	"learngo/retiever/mock"
	real2 "learngo/retiever/real"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url string = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":"ccmouse",
			"course":"golang",
		})
}

type RetriverPoster interface {
	Retriever
	Poster
}

func session(s RetriverPoster) string {
	s.Post(url, map[string]string{
		"contents":"another faked imooc.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fask imooc.com"}
	fmt.Printf("%T %v\n",r,r)
	r = real2.Retriever{}
	fmt.Printf("%T %v\n",r,r)
}
