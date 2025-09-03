package utils

import (
	"BE_Hospital_Management/constant"
	"time"
)

func CheckTimeInWorkingHours(beginTime, endTime time.Time) bool {
	if beginTime.Hour() >= constant.BeginWorkingHourMorning && endTime.Hour() <= constant.EndWorkingHourMorning {
		return true
	}
	if beginTime.Hour() >= constant.BeginWorkingHourAfternoon && endTime.Hour() <= constant.EndWorkingHourAfternoon {
		return true
	}
	return false
}
