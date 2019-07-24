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
		chapterList string
	)

	url := "http://book.zongheng.com/store/c0/c0/b0/u0/p1/v9/s1/t0/u0/i1/ALL.html"
	//url := "http://book.zongheng.com/book/860929.html"
	urlResp := GetRespBody(url)
	defer urlResp.Close()

	doc, err := goquery.NewDocumentFromReader(urlResp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Html())

	doc.Find(".bookbox .bookname").Each(func(i int, s *goquery.Selection) {
		bookname := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		fmt.Println(href)
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
			//start := s.Find("a").Eq(1).Text()
			href, _ := s.Find("a").Eq(1).Attr("href")
			//fmt.Println("start:", d.BookName)
			//fmt.Println("href:", href)

			resb := GetRespBody(href)
			defer resb.Close()
			doc1,err := goquery.NewDocumentFromReader(resb)
			if err != nil {
				log.Fatal("err:", err)
			}
			doc1.Find(".chapter-list").Each(func(i int, s *goquery.Selection) {
				chapterList ,_ = s.Find("li").Find("a").Attr("href")
				//fmt.Println("chapterList:",chapterList)

				contentHref := GetRespBody(chapterList)
				defer contentHref.Close()
				dd,err := goquery.NewDocumentFromReader(contentHref)
				if err != nil {
					log.Fatal("err:", err)
				}
				dd.Find("content").Each(func(i int, s *goquery.Selection) {
					ss := s.Text()
					fmt.Println("ss:",ss)
				})

			})
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
