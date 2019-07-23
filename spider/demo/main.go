package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://book.zongheng.com/book/785016.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(string(body))

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".btn-group").Each(func(i int, s *goquery.Selection) {
		bookname := s.Find("a").Eq(1).Text()
		href, ok := s.Find("a").Eq(1).Attr("href")
		if !ok {
			log.Fatal("failed")
		}
		fmt.Println("bookname:", bookname)
		fmt.Println("href:", href)
	})
	//start := doc.Find(".btn-group").Find("a").Text()
	//fmt.Println("start:", start)
}
