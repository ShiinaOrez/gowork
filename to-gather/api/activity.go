package api

import (
	_ "fmt"
	"github.com/kataras/iris"
	"time"
)

var  REQUIRED = []string{"year", "month", "day", "location", "event", "question", "tel", "qq", "time"}
type Token string

func ActivityPost(ctx iris.Context) {
	var data ActivityPostData
	var act Act
	var uid int

	token := Token(ctx.GetHeader("token"))
	statu, code := token.LoginRequired()
	if statu {
		uid = code
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

	currentTime := myTime(time.Now())
	if !currentTime.Before(data.Year, data.Month, data.Day) {
		ctx.StatusCode(405)
		ctx.JSON(map[string]string {
			"msg": "date invalid",
		})
		return
	} 

	act.PosterID = uid
	actID := (&act).Init(data)

	ctx.StatusCode(200)
	ctx.JSON(map[string]ActivityPostReturnData{
		"activityID": ActivityPostReturnData{actID},
	})
	return
}



