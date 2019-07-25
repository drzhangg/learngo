package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	url := "http://www.nmc.cn/publish/forecast/ABJ/beijing.html"
	doc := GetUrlDoc(url)
	doc.Find(".forecast").Each(func(i int, s *goquery.Selection) {

		for j := 0; j < 7; j++ {
			//今，明，后
			if j < 3 {
				s.Find(".detail").Find(".today").Eq(j).Find("tbody").Each(func(x int, s1 *goquery.Selection) {
					week := s1.Find("tr").Find("p").Eq(1).Text()
					date := s1.Find("tr").Find("td").Eq(1).Text()
					weather := s1.Eq(i).Find("td").Text()
					fmt.Println("week:", week)
					fmt.Println("date:", date)
					fmt.Println("weather:", weather)
				})
			} else {
				s.Find(".detail").Find(".today").Eq(j).Find("tbody").Each(func(x int, s1 *goquery.Selection) {
					date := s1.Find("tr").Find("td").Eq(0).Text()
					week := s1.Find("tr").Find("td").Eq(1).Text()
					weather := s1.Eq(i).Find("td").Eq(1).Text()
					st := s1.Eq(2).Text()
					fmt.Println("week:", week)
					fmt.Println("date:", date)
					fmt.Println("weather:", weather)
					fmt.Println("st:",st)
				})
			}
		}

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
