package entity

import (
	"time"
)

type User struct {
	Id          int64 `gorm:"primaryKey;autoIncrement"`
	Name        string
	Email       string `gorm:"unique"`
	Password    string `json:"-"`
	CitizenId   string `gorm:"unique"`
	DateOfBirth time.Time
	PhoneNumber string `gorm:"unique"`
	Address     string
	Gender      string `gorm:"type:gender_slug;default:'unknown'"`
	RoleId      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool `gorm:"default:false"`

	Role *UserRole `gorm:"foreignKey:RoleId;references:Id"`
}
