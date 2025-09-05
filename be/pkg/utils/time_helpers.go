package utils

import (
	"BE_Hospital_Management/constant"
	"time"
)

func CheckTimeInWorkingHours(beginTime, endTime time.Time) bool {
	if beginTime.Hour() >= constant.BeginWorkingHourMorning && endTime.Hour() < constant.EndWorkingHourMorning {
		return true
	}
	if beginTime.Hour() >= constant.BeginWorkingHourMorning && endTime.Hour() == constant.EndWorkingHourMorning && endTime.Minute() == 0 && endTime.Second() == 0 {
		return true
	}
	if beginTime.Hour() >= constant.BeginWorkingHourAfternoon && endTime.Hour() < constant.EndWorkingHourAfternoon {
		return true
	}
	if beginTime.Hour() >= constant.BeginWorkingHourAfternoon && endTime.Hour() == constant.EndWorkingHourAfternoon && endTime.Minute() == 0 && endTime.Second() == 0 {
		return true
	}
	return false
}
