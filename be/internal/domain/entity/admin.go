package entity

import (
	"time"
)

type Admin struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`
	Name     string
	Email    string `gorm:"unique"`
	Password string `gorm:"not null" json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool
}
