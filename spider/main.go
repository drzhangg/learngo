package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
)

type Doc struct {
	BookName string
	BookHref string
}

func main() {

	var (
		docs []Doc
	)

	resp, err := http.Get("http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s1/t0/u0/i1/ALL.html")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".bookbox .bookname").Each(func(i int, s *goquery.Selection) {
		bookname := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		docs1 := Doc{
			BookName: bookname,
			BookHref: href,
		}
		docs = append(docs, docs1)
	})

	fmt.Println("doc:", docs)

	for _, d := range docs {
		respBody := GetRespBody(d.BookHref)
		defer respBody.Close()
		doc, err := goquery.NewDocumentFromReader(respBody)
		if err != nil {
			log.Fatal("err:", err)
		}
		doc.Find(".btn-group").Each(func(i int, s *goquery.Selection) {
			start := s.Find("a").Eq(1).Text()
			href, _ := s.Find("a").Eq(1).Attr("href")
			fmt.Println("start:", start)
			fmt.Println("href:", href)
		})

	}

}

func GetRespBody(url string) io.ReadCloser {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body
}
