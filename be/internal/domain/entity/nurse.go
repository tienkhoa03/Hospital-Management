package entity

import "time"

type Nurse struct {
	Id                   int64 `gorm:"primaryKey;autoIncrement"`
	StaffId              int64 `gorm:"uniqueIndex"`
	NursingLicenseNumber string
	CreatedAt            time.Time
	UpdatedAt            time.Time

	Staff *Staff `gorm:"foreignKey:StaffId;references:Id"`
}
