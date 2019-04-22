package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n",all)

	//将读取到的html文件写入本地
	/*var d1 = []byte(all)

	err = ioutil.WriteFile("crawler/crawler/demo.html",d1,0666)
	if err != nil {
		panic(err)
	}*/

	/*re := regexp.MustCompile(`window.__INITIAL_STATE__ = {(\w+)};`)
	fmt.Println("re:",re)
	match := re.FindAll(all,-1)
	fmt.Println(match)*/

	printCityList(all)

}

//自动获取页面编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" [^>]*>([^<]+)</a>`)
	matchs := re.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		fmt.Printf("City:%s,URL:%s\n", m[2],m[1])
	}
	fmt.Printf("Matches found :%d\n",len(matchs))

}
