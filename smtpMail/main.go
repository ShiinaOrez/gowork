package main

import (
    "fmt"
    "net/smtp"
    "strings"
)

func main() {
    auth := smtp.PlainAuth("", "shiina_orez@qq.com", "", "smtp.qq.com")
    to := []string{""}
    nickname := "Test"
    user := "shiina_orez@qq.com"
    subject := "Test Mail"
    contentType := "Content-Type: text/plain; charset=UTF-8"
    body := "This is a test email~"
    msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
    err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
    if err != nil {
        fmt.Println("Send mail error:", err.Error())
    }
}
