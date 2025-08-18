package entity

import "time"

type Appointment struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	PatientId  int64
	DoctorId   int64
	BeginTime  time.Time
	FinishTime time.Time
	Status     string `gorm:"type:appointment_status_slug;default:'scheduled'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	Patient *Patient `gorm:"foreignKey:PatientId;references:Id"`
	Doctor  *Doctor  `gorm:"foreignKey:DoctorId;references:Id"`
}
