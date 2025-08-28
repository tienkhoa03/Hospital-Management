package entity

import "time"

type Manager struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	UserId     int64 `gorm:"uniqueIndex"`
	Department string
	Status     string `gorm:"type:manager_status_slug;default:'inactive'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User *User `gorm:"foreignKey:UserId;references:Id"`
}
