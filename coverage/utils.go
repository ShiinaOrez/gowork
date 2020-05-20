package utils

import (
	"encoding/json"
	"time"
	"fmt"
)

func GetTimeNow() string {
	now := time.Now()
	return fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
}

func GetTimeJSON() string {
	timeNow := struct {
		Time string `json:"time"`
	}{GetTimeNow()}
	bytes, _ := json.Marshal(timeNow)
	return string(bytes)
}