package api

import (
	"github.com/ShiinaOrez/gowork/to-gather/models"
)

type Act models.Activity

func (act *Act) Init(data ActivityPostData) int {
	var activity models.Activity
	activity.Date = GetDate(data.Year, data.Month, data.Day)
	activity.Time = data.Time
	activity.Event = data.Event
	activity.Location = data.Location
	activity.QQ = data.QQ
	activity.Question = data.Question
	activity.Tel = data.Tel
	activity.Close = false
	activity.Pickable = true
	activity.PosterID = act.PosterID
	DB.Create(&activity)
	return activity.ID
}