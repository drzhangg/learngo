package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		url       string
		paasid    string
		token     string
		timestamp string
		nonce     string
		signature string
	)
	//https://dev.bjt.beijing.gov.cn/ebus/gssdahan/jimp/qixiangju/interfaces/ + forecast.do?area=101010100
	url = "https://dev.bjt.beijing.gov.cn/ebus/gssdahan/jimp/qixiangju/interfaces/quality.do?area=101010100"
	paasid = "gssdahan"
	token = "M6tQkaMGqxlUUZZkN72Q99ToW0THwicU"
	nonce = "123456789abcdefg"

	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("url-----------", url)
	rawStr := timestamp + token + nonce + timestamp
	signature = strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(rawStr))))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err----", err)
	}
	fmt.Println("request---",request)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("paasid", paasid)
	request.Header.Set("timestamp", timestamp)
	request.Header.Set("signature", signature)
	request.Header.Set("nonce", nonce)

	//did := "9v3ee3iZnx"
	//sid := "eUH9zxGnHI/jkdKVDUHs+A=="
	//request.Header.Set("x-tif-did",did)
	//request.Header.Set("x-tif-sid",sid)

	//fmt.Println("paasid", paasid)
	//fmt.Println("timestamp", timestamp)
	//fmt.Println("signature", signature)
	//fmt.Println("nonce", nonce)


	client := &http.Client{}
	resp,err := client.Do(request)
	if err != nil {
		fmt.Println("err2----", err)
	}
	defer resp.Body.Close()
	fmt.Println("resp",resp)

	//s,err := httputil.DumpResponse(resp,true)
	//if err != nil {
	//	panic(err)
	//}

	body,err :=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err3----", err)
	}
	fmt.Println(string(body))

}
