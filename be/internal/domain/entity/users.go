package entity

import (
	"time"
)

type User struct {
	Id          int64     `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"type:varchar(256);not null"`
	Email       string    `gorm:"type:varchar(256);not null;unique"`
	Password    string    `gorm:"type:varchar(256);not null" json:"-"`
	CitizenId   string    `gorm:"type:varchar(256);not null"`
	DateOfBirth time.Time `json:"date_of_birth"`
	PhoneNumber string    `gorm:"type:varchar(256);not null"`
	Address     string    `gorm:"type:varchar(256);not null"`

	Gender string `gorm:"type:gender_slug"`

	RoleId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
	Role      *Role `gorm:"foreignKey:RoleId;references:Id"`
}
