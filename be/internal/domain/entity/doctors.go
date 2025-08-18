package entity

import "time"

type Doctor struct {
	Id                   int64 `gorm:"primaryKey;autoIncrement"`
	StaffId              int64
	Specialization       string
	MedicalLicenseNumber string
	CreatedAt            time.Time
	UpdatedAt            time.Time

	Staff *Staff `gorm:"foreignKey:StaffId;references:Id"`
}
