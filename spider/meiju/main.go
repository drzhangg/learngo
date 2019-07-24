package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	url := "https://91mjw.com/"
	doc := GetUrlDoc(url)
	//fmt.Println(doc.Text())
	doc.Find("div.container.header ul.nav").Each(func(i int, s *goquery.Selection) {
		pai, _ := s.Find("#menu-item-1436").Find("a").Attr("href")
		fmt.Println(pai)

		sixpai := GetUrlDoc(pai)
		sixpai.Find(".m-movies").Each(func(i int, s *goquery.Selection) {
			s.Find(".u-movie").Each(func(i int, s1 *goquery.Selection) {
				href,_ := s1.Find("a").Attr("href")
				name := s1.Find("a").Text()
				fmt.Printf("name:%v,href:%v\n",name,href)
			})
		})
	})
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
