package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Content struct {
	NickName    string
	User        string
	Subject     string
	Body        string
	ContentType string
}

func SendMail(from, authCode string, to []string, content Content) error {
	auth := smtp.PlainAuth("", from, authCode, "smtp.qq.com")
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + content.NickName + "<" + content.User + ">\r\nSubject: " + content.Subject + "\r\n" + content.ContentType + "\r\n\r\n" + content.Body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err != nil {
		return err
	}
	return nil
}
