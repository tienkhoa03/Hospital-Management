package entity

import (
	"time"
)

type UserRole struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`
	RoleSlug string

	CreatedAt time.Time
	UpdatedAt time.Time
}

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleStaff   = "staff"
	RolePatient = "patient"
)
