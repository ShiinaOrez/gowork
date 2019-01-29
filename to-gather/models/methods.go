package models

import (
	"github.com/ShiinaOrez/gowork/to-gather/api"
)

func (act *Activity) Init(data api.ActivityPostData) bool {
	act.Date = api.GetDate(data.Year, data.Month, data.Day)
	act.Time = data.Time
	act.Event = data.Event
	act.Location = data.Location
	act.QQ = data.QQ
	act.Question = data.Question
	act.Tel = data.Tel
	act.Close = false
	act.Pickable = true
	api.DB.create(act)
	return true
}
