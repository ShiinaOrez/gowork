package api

import (
	_ "fmt"
	"github.com/ShiinaOrez/gowork/to-gather/models"
	"github.com/kataras/iris"
	"time"
)

var  REQUIRED = []string{"year", "month", "day", "location", "event", "question", "tel", "qq", "time"}
type Token string

func ActivityPost(ctx iris.Context) {
	var data ActivityPostData
	var act models.Activity

	token := Token(ctx.GetHeader("token"))
	statu, code := token.LoginRequired()
	if statu {
		uid := code
	} else {
		ctx.StatusCode(code)
		ctx.JSON(map[string]string{
			"msg": "token invalid",
		})
		return
	}

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	currentTime = time.Now()
	if currentTime.Before(data.year, data.month, data.day) {
		ctx.StatusCode(405)
		ctx.JSON(map[string]string {
			"msg": "date invalid",
		})
	} 

	act.PosterID = uid
	act.Date
	act.Init(data)
	/*var data LoginPostData
	var usr models.User
	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	if DB.Where("stdnum=?", data.Username).First(&usr).RecordNotFound() {
		//add a new record into User
		usr.Name = data.Username
		usr.StdNum = data.StdNum
		DB.Create(&usr)

	}
	ReturnData := data_structure2.LoginReturnData {
		strconv.Itoa(usr.ID) + usr.Name,
		usr.ID,
		usr.StdNum,
	}
	ctx.StatusCode(200)
	ctx.JSON(map[string]LoginReturnData {
		"Data": ReturnData,
	})
	return*/
}



