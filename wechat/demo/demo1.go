package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	grant_type  = ""
	appid       = "wx85005ccfa27148c2"
	secret      = "55a5dc60f6f1295b514ee88cb8596bc3"
	accessToken = "25_KXwDBGVOkFKWvBv_ZVeuMKSiuGDeION8eTDQR4J58nrTXETCYXqMAe1XLBpPXEpIzsaIH_KRr0HMzr7NcPNgoIfjd5H5VAHK88Sk3I5Tb6PktytMiNOYBJAyn16Sl5bxiCCTPZe1-7G_yC7FEZQbAIAUPG"
)

type RespData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func main() {

	var url = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=` +
		appid + `&secret=` + secret
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	var respData = &RespData{}

	byte, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(byte, respData)
	fmt.Println(string(byte))
	fmt.Println(respData.AccessToken, respData.ExpiresIn)

}
