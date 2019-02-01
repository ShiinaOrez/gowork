package api

import (
	"github.com/ShiinaOrez/gowork/to-gather/models"
	_ "github.com/jinzhu/gorm"
	"strconv"
	"strings"
	"time"
)

type myTime time.Time

type DatabaseRecord interface {
	TableName()
}

type GetPageData struct {
	Data    []DatabaseRecord
	PageNow  int
	PageMax  int
	HasNext  bool
	RowsNum  int
}

func (token Token) LoginRequired() (bool, int) {
	var usr models.User
	sli := strings.Split(string(token), "?")
	uid, err := strconv.Atoi(sli[0])
	if err != nil {
		return false, 401
	}
	if DB.Where("id=?", uid).First(&usr).RecordNotFound() {
		return false, 401
	} else {
		if usr.Name != sli[1] {
			return false, 401
		} else {
			return true, uid
		}
	}
}

func GetDate(year, month, day int) time.Time {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	return time.Date(year, time.Month(month), day, 00, 0,0,0, beijing)
}

func (now myTime) Before(year, month, day int) bool {
	currentYear := time.Time(now).Year()
	currentMonth := int(time.Time(now).Month())
	currentDay := time.Time(now).Day()
	if (currentYear > year) ||
	   ((currentYear == year) && (currentMonth > month)) ||
	   ((currentYear == year) && (currentMonth == month) && (currentDay > day)) {
		return false
	} else {
		return true
	}
}

func GetPage(OriginData []DatabaseRecord, PageSize int, PageNow int) (GetPageData, bool) {
	var has_next bool
	rows_num := len(OriginData)
	page_now := PageNow
	page_max := (rows_num/PageSize) + 1
	if page_now<page_max {
		has_next = true
	} else {
		if page_now > page_max {
			return GetPageData{}, false
		}
		has_next = false
	}
	var data []DatabaseRecord
	for i := PageSize*(page_now - 1); i < PageSize*page_now; i++ {
		if i == rows_num {
			break
		} else {
			data = append(data, OriginData[i])
		}
	}
	return GetPageData{
		data,
		page_now,
		page_max,
		has_next,
		rows_num}, true
}