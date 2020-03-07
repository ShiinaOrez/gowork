package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/url"
	"strings"
	"net/http"
	"net/http/cookiejar"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Info struct {
	Data  struct{
		UserInfoVO  struct{
			ID  string `json:"id"`
		} `json:"userInfoVO"`
	} `json:"data"`
}

type Response struct {
	Code      uint8     `json:"code"`
	Msg       string    `json:"msg"`
	Data      struct{
		UserInfoVO    struct{
			ID        string    `json:"id"`
			Archived  bool      `json:"archived"`
			Username  string    `json:"username"`
			Phone     string    `json:"phone"`
			Email     string    `json:"email"`
			UserInfo    struct{
				ID          string    `json:"id"`
				RealName    string    `json:"realname"`
				Sex         string    `json:"sex"`
				Age         uint      `json:"age"`
				Phone       string    `json:"phone"`
				Email       string    `json:"email"`
				QQ          string    `json:"qq"`
				WeChat      string    `json:"wechat"`
				DepartmentCode  string  `json:"departmentCode"`
				University  string    `json:"university"`
				LoginName   string    `json:"loginName"`
				HeadImageURL    string  `json:"headImageUrl"`
				Sign        string    `json:"sign"`
				AddTime     int       `json:"addtime"`
				UpdateTime  int       `json:"updatetime"`
			}    `json:"userInfo"`
		}    `json:"userInfoVO"`
		RoleDepartment    struct{
			ID         string    `json:"id"`
			LoginName  string    `json:"loginName"`
			DomainCode string    `json:"domainCode"`
			RoleCode   string    `json:"roleCode"`
			DepartmentCode  string  `json:"departmentCode"`
			AddTime    int       `json:"addtime"`
			UpdateTime int       `json:"updatetime"`
			DomainName string    `json:"domainName"`
			DepartmentName  string  `json:"departmentName"`
			RoleName   string    `json:"roleName"`
			RealName   string    `json:"realname"`
		}    `json:"roleDepartment"`
	}    `json:"data"`
}

func LoginSPOC(sno, password string, client *http.Client) (*Response, error) {
	v := url.Values{}
	v.Set("loginName", sno)
	v.Set("password", password)

	request, err := http.NewRequest("GET", "http://spoc.ccnu.edu.cn/userLoginController/getVerifCode", nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	_, err = client.Do(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	request, err = http.NewRequest("POST", "http://spoc.ccnu.edu.cn/userLoginController/getUserProfile", strings.NewReader(v.Encode()))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Host", "spoc.ccnu.edu.cn")
    request.Header.Set("Origin", "http://spoc.ccnu.edu.cn")
	request.Header.Set("Referer", "http://spoc.ccnu.edu.cn/studentHomepage")
	_, err = client.Do(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	request, err = http.NewRequest("POST", "http://spoc.ccnu.edu.cn/userInfo/getUserInfo", strings.NewReader(v.Encode()))
	if err != nil {
		log.Print(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Host", "spoc.ccnu.edu.cn")
    request.Header.Set("Origin", "http://spoc.ccnu.edu.cn")
	request.Header.Set("Referer", "http://spoc.ccnu.edu.cn/studentHomepage")
	resp, respErr := client.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if respErr != nil {
		log.Print(respErr)
		return nil, respErr
	}
	if _, ok := resp.Header["Content-Language"]; !ok {
		info := Info{}
		err = json.Unmarshal(body, &info)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		// uid := info.Data.UserInfoVO.ID

		request, err = http.NewRequest("GET", "http://spoc.ccnu.edu.cn/userInfo/getUserInfo", nil)
		resp, err = client.Do(request)
		body, err = ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()

		response := Response{}
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		return &response, nil
	}
	return nil, nil
}

func main() {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	
	response, err := LoginSPOC("2017211712", "hjl20030119", &client)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(response.Data.RoleDepartment)
}