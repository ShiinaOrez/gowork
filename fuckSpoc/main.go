package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/url"
	"regexp"
	"strings"
)

type Session map[string]string

const (
	temp1 = "https://account.ccnu.edu.cn/cas/login"
	temp2 = "?service=http%3A%2F%2Fspoc.ccnu.edu.cn%2FuserLoginController%2FuserCasLogin"
)

func main() {
	var session Session = make(map[string]string)
	client := http.Client{}

	resp, err := client.Get(temp1 + temp2)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	for _, cookie := range resp.Cookies() {
		cookiePairs := strings.Split(cookie.String(), "; ")
		for _, pair := range cookiePairs {
			s := strings.Split(pair, "=")
			if len(s) == 1 {
				continue
			}
			session[s[0]] = s[1]
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	eventR, _ := regexp.Compile("execution\" value=\"(.*)\" />")
	r, _ := regexp.Compile("LT-(.*)-account.ccnu.edu.cn")
	lt := r.FindString(string(body))
	eventId := eventR.FindString(string(body))

	resp.Body.Close()

	formData := fmt.Sprintf(
		"username=%s&password=%s&execution="+eventId[18:22]+"&_eventId=submit&lt=%s",
		"xxx",
		"xxx",
		lt,
	) + "&submit=%E7%99%BB%E5%BD%95"

	request, err := http.NewRequest(
		"POST",
		temp1+temp2,
		strings.NewReader(formData),
	)
	request.AddCookie(&http.Cookie{
		Name:  "CASPRIVACY",
		Value: "",
		Path:  "/cas/",
	})
	request.AddCookie(&http.Cookie{
		Name:  "CASTGC",
		Value: "",
		Path:  "/cas/",
	})
	request.AddCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: session["JSESSIONID"],
		Path:  "/cas/",
	})
	for _, c := range request.Cookies() {
		fmt.Println(c.String())
	}
	if err != nil {
		log.Println("Error:", err)
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	request.Header.Set("Cache-Control", "no-cache")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Host", "account.ccnu.edu.cn")
	request.Header.Set("Origin", "https://account.ccnu.edu.cn")
	request.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fspoc.ccnu.edu.cn%2FuserLoginController%2FuserCasLogin&t=0.6981381970870573")
	request.Header.Set("Sec-Fetch-Mode", "navigate")
	request.Header.Set("Sec-Fetch-Site", "same-origin")
	request.Header.Set("Sec-Fetch-User", "?1")
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err = client.Do(request)

	fmt.Println(resp.Status, resp.Header.Get("Location"))
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
	return
}
