package entity

import "time"

type Patient struct {
	Id                int64 `gorm:"primaryKey"`
	UserId            int64
	InsuranceNumber   string
	BloodType         string
	Allergies         string
	ChronicConditions string
	Status            string `gorm:"patient_status_slug"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	User *User `gorm:"foreignKey:UserId;references:Id"`
}
