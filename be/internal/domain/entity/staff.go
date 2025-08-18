package entity

import "time"

type Staff struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	UserId     int64
	ManageBy   int64
	Department string
	Status     string `gorm:"type:staff_status_slug;default:'inactive'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User    *User    `gorm:"foreignKey:UserId;references:Id"`
	Manager *Manager `gorm:"foreignKey:ManageBy;references:Id"`
}
