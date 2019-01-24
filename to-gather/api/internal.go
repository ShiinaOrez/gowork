package api

import (
	"time"
)

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