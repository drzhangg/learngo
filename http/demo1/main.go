package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	url := "https://cn.bing.com/"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	req := resp.Body
	fmt.Println(req)
	node,_:=html.Parse(req)
	fmt.Println(node.Attr)
	doc, err := goquery.NewDocumentFromReader(req)
	if err != nil {
		panic(err)
	}
	//fmt.Println(doc.Html())

	ss := doc.Find("#bgImgProgLoad").AddBack()
	fmt.Println(ss)

	doc.Find("#bgImgProgLoad").Each(func(i int, s *goquery.Selection) {
		t := s.Find("data-ultra-definition-src").Text()
		fmt.Println(t)
	})

}
