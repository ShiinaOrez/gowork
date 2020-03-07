package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ShiinaOrez/LoginCCNU/spoc"
)

type Request struct {
    Sno      string `json:"sno"`
    Password string `json:"password"`
}

type Response struct {
    Code     int    `json:"code"`
    Data     struct{
        Sno    string  `json:"sno"`
        Name   string  `json:"name"`
        Gender int     `json:"gender"`
        // School string  `json:"school"`
        Major  string  `json:"major"`
	}    `json:"data"`
	Msg      string `json:"msg"`
}

func main() {
    router := gin.Default()
    router.POST("/api/signin", func(c *gin.Context) {
		request := Request{}
		response := Response{}
		if err := c.BindJSON(&request); err != nil {
			response.Code = 1
		} else {
			resp, err := spoc.Login(request.Sno, request.Password)
			if resp == nil || err != nil {
				response.Code = 1
			} else {
				response.Code = 0
				response.Data.Sno = request.Sno
				response.Data.Name = resp.Data.RoleDepartment.RealName
				response.Data.Gender = func(s string) int {
					if s == "男" { return 0 }
					return 1
				}(resp.Data.UserInfoVO.UserInfo.Sex)
				response.Data.Major = resp.Data.RoleDepartment.DepartmentName
				response.Msg = "登陆成功"
				c.JSON(200, response)
				return
			}
		}
		response.Msg = "登陆失败"
    	c.JSON(401, response)
    })
    router.Run(":2020")
}
