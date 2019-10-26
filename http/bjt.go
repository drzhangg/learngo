package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
	"time"
)

func GetGSSWeatherData(c *fasthttp.RequestCtx) {
	url := ""

	//request := c.Request
	request := &fasthttp.Request{}
	request.SetRequestURI(url)

	paasid := "gssdahan"
	token := "M6tQkaMGqxlUUZZkN72Q99ToW0THwicU"
	nonce := "123456789abcdefg"

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("url-----------", url)
	rawStr := timestamp + token + nonce + timestamp
	signature := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(rawStr))))


	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("paasid", paasid)
	request.Header.Set("timestamp", timestamp)
	request.Header.Set("signature", signature)
	request.Header.Set("nonce", nonce)

	did := "9v3ee3iZnx"
	sid := "jeyGw9Gv/Y6hMX86woscnQ=="
	request.Header.Set("x-tif-did",did)
	request.Header.Set("x-tif-sid",sid)

	resp := &fasthttp.Response{}
	client := &fasthttp.Client{}
	if err := client.Do(request, resp);err != nil {
		fmt.Println("请求失败:", err.Error())
		return
	}

	b := resp.Body()
	fmt.Println("result:\r\n", string(b))

}

func main() {
	c := &fasthttp.RequestCtx{}
	GetGSSWeatherData(c)
}
