package main

import (
	"encoding/xml"
	"fmt"
)

type Config struct {
	Config       xml.Name `xml:"config"`
	SmtpServer   string   `xml:"smtpServer"`
	SmtpPort     string   `xml:"smtpPort"`
	Sender       string   `xml:"sender"`
	SenderPasswd int64    `xml:"senderPasswd"`
	Receivers    Receiver `xml:"receivers"`
}

type Receiver struct {
	Flag  string   `xml:"flag,attr"`
	Age   int64    `xml:"age"`
	Users []string `xml:"user"`
}

func main() {

	xmlfile := `
			<config>
			   <smtpServer>smtp.163.com</smtpServer>
			   <smtpPort>25</smtpPort>
			   <sender>user@163.com</sender>
			  <senderPasswd>123456</senderPasswd>
			   <receivers flag="true">
				<age>16</age>
				 <user>Mike_Zhang@live.com</user>
				 <user>test1@qq.com</user>
			  </receivers>
			 </config>
		`

	//byte, err := ioutil.ReadFile(xmlfile)
	var config = &Config{}
	err := xml.Unmarshal([]byte(xmlfile), config)
	if err != nil {
		fmt.Println("err-----", err)
		return
	}
	fmt.Println(config)
	//fmt.Println("解析后的：", string(byte))
}
