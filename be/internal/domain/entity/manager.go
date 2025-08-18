package entity

import "time"

type Manager struct {
	Id         int64 `gorm:"primaryKey"`
	UserId     int64 `json:"id"`
	Department string
	Status     string `gorm:"staff_status_slug"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User *User `gorm:"foreignKey:UserId;references:Id"`
}
