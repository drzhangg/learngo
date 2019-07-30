package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"log"
)

func main() {
	url := "https://www.zhipin.com/c101210100/e_104/?query=java&ka=sel-exp-104"
	doc := GetUrlDoc(url)
	text := doc.Find(".job-list").Find("ul").Each(func(i int, s *goquery.Selection) {
		li := s.Find("li").Find(".job-primary").Find(".info-primary").Find(".name").Find("a").Text()
		fmt.Println(li)
	})
	fmt.Println(text)
}

func GetUrlDoc(href string) *goquery.Document {
	resp, err := http.Get(href)
	if err != nil {
		log.Fatal("err:", err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}
