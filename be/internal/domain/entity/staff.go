package entity

import "time"

type Staff struct {
	Id         int64 `gorm:"primaryKey;autoIncrement"`
	UserId     int64 `gorm:"uniqueIndex"`
	ManageBy   int64
	Department string
	RoleId     int64
	Status     string `gorm:"type:staff_status_slug;default:'inactive'"`
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User    *User      `gorm:"foreignKey:UserId;references:Id"`
	Manager *Manager   `gorm:"foreignKey:ManageBy;references:Id"`
	Role    *StaffRole `gorm:"foreignKey:RoleId;references:Id"`
}
