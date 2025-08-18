package entity

import "time"

type Patient struct {
	Id                int64 `gorm:"primaryKey;autoIncrement"`
	UserId            int64
	InsuranceNumber   *string `gorm:"unique"`
	BloodType         *string `gorm:"type:blood_type_slug"`
	Allergies         *string
	ChronicConditions *string
	Status            string `gorm:"type:patient_status_slug;default:inactive"`
	CreatedAt         time.Time
	UpdatedAt         time.Time

	User *User `gorm:"foreignKey:UserId;references:Id"`
}
