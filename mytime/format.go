package mytime

import (
	"strings"
	"time"
)

const (
	format = "20060102150405"
	numStr = "0123456789"
)

func StrToTime(s string) int64 {
	listStr := strings.Split(s, "")
	tmpArr := []string{}
	for _, i := range listStr {
		index := strings.Index(numStr, i)
		if index != -1 {
			tmpArr = append(tmpArr, i)
		}
	}
	tmpArr = tmpArr[:14]
	formatTime := strings.Join(tmpArr, "")
	loc, _ := time.LoadLocation("Local")
	the_time, _ := time.ParseInLocation(format, formatTime, loc)

	unix_time := the_time.Unix()
	return unix_time
}
