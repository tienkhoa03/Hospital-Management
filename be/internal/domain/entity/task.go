package entity

import "time"

type Task struct {
	Id          int64 `gorm:"primaryKey;autoIncrement"`
	StaffId     int64 `json:"doctor_id"`
	AssignerId  int64 `json:"assigner_id"`
	Title       string
	Description *string
	BeginTime   time.Time
	FinishTime  time.Time
	Status      string `gorm:"type:task_status_slug;default:'scheduled'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Staff    *Staff   `gorm:"foreignKey:StaffId;references:Id"`
	Assigner *Manager `gorm:"foreignKey:AssignerId;references:Id"`
}
