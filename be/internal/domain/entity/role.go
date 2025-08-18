package entity

import (
	"time"
)

type Role struct {
	Id       int64  `gorm:"primaryKey;autoIncrement"`
	RoleSlug string `gorm:"type:role_slug"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
