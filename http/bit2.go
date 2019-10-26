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

	url := "https://dev.bjt.beijing.gov.cn/ebus/gssdahan/jimp/qixiangju/interfaces/forecast.do?area=101010100"

	paasid := "gssdahan"
	token := "M6tQkaMGqxlUUZZkN72Q99ToW0THwicU"
	nonce := "123456789abcdefg"

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	rawStr := timestamp + token + nonce + timestamp
	signature := strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256([]byte(rawStr))))

	request, _ := http.NewRequest("POST", url, nil)

	//request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("x-tif-paasid", paasid)
	request.Header.Set("x-tif-timestamp", timestamp)
	request.Header.Set("x-tif-signature", signature)
	request.Header.Set("x-tif-nonce", nonce)

	client := &http.Client{}
	resp, _ := client.Do(request)
	data,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
}
