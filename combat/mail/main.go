package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func main() {

	auth := smtp.PlainAuth("","654014730@qq.com","xxxx","smtp.qq.com")
	to := []string{"949104693@qq.com"}
	nickname := "test"
	user := "654014730@qq.com"
	subject := "testtestestestestestest email"
	content_type := "Content-Type:text/plain;charset=UTF-8"
	body := "This is the email body"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25",auth,user,to,msg)
	if err != nil {
		fmt.Printf("send email error : %v",err)
	}
}
