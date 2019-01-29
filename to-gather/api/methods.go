package api

import (
	"github.com/ShiinaOrez/gowork/to-gather/models"
)

type Act models.Activity

func (act *Act) Init(data ActivityPostData) bool {
	act.Date = GetDate(data.Year, data.Month, data.Day)
	act.Time = data.Time
	act.Event = data.Event
	act.Location = data.Location
	act.QQ = data.QQ
	act.Question = data.Question
	act.Tel = data.Tel
	act.Close = false
	act.Pickable = true
	DB.Create(act)
	return true
}