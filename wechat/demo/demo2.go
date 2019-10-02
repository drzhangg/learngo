package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

type XmlText struct {
	//Xml          xml.Name `xml:"xml"`
	ToUserName   string `xml:"ToUserName"`   //开发者微信号
	FromUserName string `xml:"FromUserName"` //发送方账号(一个OpenId)
	CreateTime   string `xml:"CreateTime"`   //消息创建时间
	MsgType      string `xml:"MsgType"`      //消息类型
	Content      string `xml:"Content"`      //文本消息内容
	MsgId        string `xml:"MsgId"`        //消息id
}

func main() {
	var xmlText = `<xml>
					  <ToUserName>654014730@qq.com</ToUserName>
					  <FromUserName>654014730</FromUserName>
					  <CreateTime>%d</CreateTime>
					  <MsgType>%s</MsgType>
					  <Content>%s</Content>
					  <MsgId>1234567890123456</MsgId>
					</xml>`

	xmlText = fmt.Sprintf(xmlText, time.Now().Unix(), "text", "测试内容")
	var xmlFile = &XmlText{}
	err := xml.Unmarshal([]byte(xmlText), xmlFile)
	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(xmlFile)
}
