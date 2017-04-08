package common

import (
	"strconv"
	"time"
)

type MyTime struct {
	Year   string
	Month  string
	Day    string
	Hour   string
	Minute string
	Sec    string
}

func NewMyTime(time time.Time) *MyTime {
	Year := strconv.Itoa(time.Year())
	Month := getDoubleNum(int(time.Month()))
	Day := getDoubleNum(time.Day())
	Hour := getDoubleNum(time.Hour())
	Minute := getDoubleNum(time.Minute())
	Sec := getDoubleNum(time.Second())

	return &MyTime{Year, Month, Day, Hour, Minute, Sec}
}

func getDoubleNum(num int) string {
	if num < 10 {
		return "0" + strconv.Itoa(num)
	} else {
		return strconv.Itoa(num)
	}
}
