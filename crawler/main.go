package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code",resp.StatusCode)
		return
	}

	//当网页编码为gbk时需要转换编码
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all)
	printCityAll(all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}

func printCityAll(contents []byte) {
	re := regexp.MustCompile(`<a target="_blank" href="http://www.zhenai.com/zhenghun/[0-9a-z] + "[^>]*[^<]+</a>`)
	fmt.Println(re)
	matches := re.FindAll(contents,-1)
	fmt.Println(matches)
	for _,m := range matches{
		fmt.Printf("%s\n",m)
	}
}
