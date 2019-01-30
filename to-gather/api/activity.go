package api

import (
	_ "fmt"
	"github.com/ShiinaOrez/gowork/to-gather/models"
	"github.com/kataras/iris"
	"strings"
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

func ActivityGet(ctx iris.Context) {
	var returnData ActivityGetReturnData
	var act models.Activity
	var usr, poster models.User
	var uid int

	aid := ctx.Params().Get("aid")

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

	DB.Where("id=?", uid).First(&usr)
	if DB.Where("id=?", aid).First(&act).RecordNotFound() {
		ctx.StatusCode(406)
		ctx.JSON(map[string]string{
			"msg": "activity not found",
		})
	}
	DB.Where("id=?", act.PosterID).First(&poster)

	currentTime := myTime(time.Now())
	if !currentTime.Before(act.Date.Year(), int(act.Date.Month()), act.Date.Day()) {
		act.Close = true
		act.Pickable = false
		DB.Save(&act)
		ctx.StatusCode(403)
		ctx.JSON(map[string]string {
			"msg": "activity out of time range",
		})
		return
	}

	returnData = ActivityGetReturnData{
		strings.Split(act.Date.String(), " ")[0],
		act.Time,
		act.Location,
		act.Event,
		act.QQ,
		act.Tel,
		act.Question,
		poster.Name,
		poster.StdNum,
		ActivityStatu{
			act.Pickable,
			act.Close,
		},
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]ActivityGetReturnData{
		"ActivityInfo": returnData,
	})
}

func ActivityPick(ctx iris.Context) {
	var data ActivityPickData
	aid := ctx.Params().Get("aid")
    var uid int
	var act models.Activity
	var usr models.User

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

	DB.Where("id=?", uid).First(&usr)
	if DB.Where("id=?", aid).First(&act).RecordNotFound() {
		ctx.StatusCode(406)
		ctx.JSON(map[string]string{
			"msg": "activity not found",
		})
		return
	}

	if act.Close {
		ctx.StatusCode(405)
		ctx.JSON(map[string]string{
			"msg": "activity is over",
		})
		return
	}



	currentTime := myTime(time.Now())
	if !currentTime.Before(act.Date.Year(), int(act.Date.Month()), act.Date.Day()) {
		act.Close = true
		act.Pickable = false
		DB.Save(&act)
		ctx.StatusCode(403)
		ctx.JSON(map[string]string {
			"msg": "activity out of time range",
		})
		return
	}

}