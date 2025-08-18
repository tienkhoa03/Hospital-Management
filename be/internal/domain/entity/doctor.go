package entity

import "time"

type Doctor struct {
	Id             int64 `gorm:"primaryKey"`
	UserId         int64 `json:"id"`
	ManageBy       *int64
	Specialization string
	Department     string
	Status         string `gorm:"staff_status_slug"`
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User     *User    `gorm:"foreignKey:UserId;references:Id"`
	Assigner *Manager `gorm:"foreignKey:ManageBy;references:Id"`
}
