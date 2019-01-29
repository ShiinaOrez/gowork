package api

import (
	"time"
	"strings"
	"strconv"
	"github.com/ShiinaOrez/gowork/to-gather/models"
)


/*
func LoginRequired:
    return false, statusCode
        or true, userID
 */
func (token Token) LoginRequired() (bool, int) {
	var usr models.User
	sli := strings.Split(token, "?")
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

func (now time.Time) Before(year, month, day int) bool {
	currentYear := now.Year()
	currentMonth := int(now.Month())
	currentDay := now.Day()
	if (currentYear > year) ||
	   ((currentYear == year) && (currentMonth > month)) ||
	   ((currentYear == year) && (currentMonth == month) && (currentDay > day)) {
		return false
	} else {
		return true
	}
}