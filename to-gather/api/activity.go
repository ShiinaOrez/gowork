package api

import (
	_ "fmt"
	"github.com/ShiinaOrez/gowork/to-gather/models"
	"github.com/kataras/iris"
	_ "strconv"
	"strings"
	"time"
)

var REQUIRED = []string{"year", "month", "day", "location", "event", "question", "tel", "qq", "time"}

type Token string

func ActivityPost(ctx iris.Context) {
	var data ActivityPostData
	var act Act
	uid, _ := ctx.Values().GetInt("uid")

	err := ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	currentTime := myTime(time.Now())
	if !currentTime.Before(data.Year, data.Month, data.Day) {
		ctx.StatusCode(405)
		ctx.JSON(map[string]string{
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
	uid, _ := ctx.Values().GetInt("uid")
	aid, _ := ctx.Params().GetInt("aid")

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
		ctx.JSON(map[string]string{
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
	aid, _ := ctx.Params().GetInt("aid")
	uid, _ := ctx.Values().GetInt("uid")
	var act models.Activity
	var usr models.User
	var record models.Picker2Activity

	token := Token(ctx.GetHeader("token"))
	statu, code := token.LoginRequired()
	if statu {
		uid = code
	} else {
		ctx.StatusCode(code)
		ctx.JSON(map[string]string{
			"msg": "authenticate fail",
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

	if act.PosterID == usr.ID {
		ctx.StatusCode(407)
		ctx.JSON(map[string]string{
			"msg": "Can't pick yourself.",
		})
	}

	if !DB.Where("aid=? AND picker_id=?", act.ID, usr.ID).First(&record).RecordNotFound() {
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
			"msg": "Already picked it before!",
		})
		return
	}

	currentTime := myTime(time.Now())
	if !currentTime.Before(act.Date.Year(), int(act.Date.Month()), act.Date.Day()) {
		act.Close = true
		act.Pickable = false
		DB.Save(&act)
		ctx.StatusCode(403)
		ctx.JSON(map[string]string{
			"msg": "activity out of time range",
		})
		return
	}

	//Ready to pick
	record.PickerID = usr.ID
	record.AID = act.ID
	record.Answer = data.Answer
	record.Fail = false
	record.Waiting = true
	DB.Create(&record)
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(map[string]string{
		"msg": "pick the activity successful!",
	})
}

func ActivityPut(ctx iris.Context) {
	var data ActivityPutData
	var act models.Activity
	var usr models.User
	var record models.Picker2Activity
	aid, err := ctx.Params().GetInt("aid")
	uid, _ := ctx.Values().GetInt("uid")

	token := Token(ctx.GetHeader("token"))
	statu, code := token.LoginRequired()
	if statu {
		uid = code
	} else {
		ctx.StatusCode(code)
		ctx.JSON(map[string]string{
			"msg": "authenticate fail",
		})
		return
	}

	err = ctx.ReadJSON(&data)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}

	currentTime := myTime(time.Now())
	if !currentTime.Before(act.Date.Year(), int(act.Date.Month()), act.Date.Day()) {
		act.Close = true
		act.Pickable = false
		DB.Save(&act)
		ctx.StatusCode(403)
		ctx.JSON(map[string]string{
			"msg": "activity out of time range",
		})
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

	if data.PickerID == uid {
		ctx.StatusCode(407)
		ctx.JSON(map[string]string{
			"msg": "Can't pick yourself.",
		})
	}

	if !DB.Where("aid=? AND picker_id=?", act.ID, usr.ID).First(&record).RecordNotFound() {
		ctx.StatusCode(402)
		ctx.JSON(map[string]string{
			"msg": "Already picked it before!",
		})
		return
	}

	if !DB.Where("aid=? AND picker_id=?", act.ID, data.PickerID).First(&record).RecordNotFound() {
		if data.Atti {
			record.Waiting = false
			record.Fail = false
			ctx.StatusCode(iris.StatusOK)
			ctx.JSON(map[string]string{
				"msg": "makesure successful!",
			})
		} else {
			record.Waiting = false
			record.Fail = true
			ctx.StatusCode(201)
			ctx.JSON(map[string]string{
				"msg": "refuse succeddful!",
			})
		}
		return
	} else {
		ctx.StatusCode(406)
		ctx.JSON(map[string]string{
			"msg": "User does not pick this activity!",
		})
		return
	}
}

/*func ActivityPickableList(ctx iris.Context) {
	var activities []DatabaseRecord
	page_string := ctx.Params().Get("")
    page, _ := strconv.Atoi(page_string)
	DB.Where("pickable=?",true).Find(&activities)
	getPage, statu := GetPage(activities, 10, page)
	if !statu {
		ctx.StatusCode(201)
		ctx.JSON(map[string]string{
			"msg": "Page out of Range",
		})
		return
	} else {
		if getPage.RowsNum == 0 {
			ctx.StatusCode(201)
			ctx.JSON(map[string]string{
				"msg": "None pickable activity",
			})
			return
		} else {
			ctx.StatusCode(iris.StatusOK)
			var finalData []ActivityInfo
			for data, _ := range(getPage.Data) {

			}
			ctx.JSON(map[string]ActivityPickableListReturnData{
				"List": ActivityPickableListReturnData{
					getPage.Data,
					getPage.PageNow,
					getPage.PageMax,
					getPage.HasNext,
					getPage.RowsNum
				},
			})
		}
	}
}
*/
