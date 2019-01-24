package api

import (
	"time"
)

type LoginPostData struct {
	Username      string      `json:"username"`
	StdNum        string      `json:"std_num"`
}

type LoginReturnData struct {
	Token         string      `json:"token"`
	UID           int         `json:"uid"`
	StdNum        string      `json:"std_num"`
}

type ActivityPostData struct {
	Year          int         `json:"year"`
	Month         int         `json:"month"`
	Day           int         `json:"day"`
	Time          time.Time   `json:"time"`
	Location      string      `json:"location"`
	Event         string      `json:"event"`
	Question      string      `json:"question"`
	Tel           string      `json:"tel"`
	QQ            string      `json:"qq"`
}

type ActivityPostReturnData struct {
	ActivityID    int         `json:"activity_id"`
}