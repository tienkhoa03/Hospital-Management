package utils

import "time"

const (
	beginWorkingHourMorning   = 7
	endWorkingHourMorning     = 11
	beginWorkingHourAfternoon = 13
	endWorkingHourAfternoon   = 17
)

func CheckTimeInWorkingHours(beginTime, endTime time.Time) bool {
	if beginTime.Hour() >= beginWorkingHourMorning && endTime.Hour() <= endWorkingHourMorning {
		return true
	}
	if beginTime.Hour() >= beginWorkingHourAfternoon && endTime.Hour() <= endWorkingHourAfternoon {
		return true
	}
	return false
}
