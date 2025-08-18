package entity

import "time"

type Task struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	DoctorId    int64  `json:"doctor_id"`
	AssignerId  int64  `json:"assigner_id"`
	Description string `gorm:"type:varchar(256)"`
	BeginTime   time.Time
	FinishTime  time.Time
	Status      string `gorm:"type:task_status_slug"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Doctor   *Doctor  `gorm:"foreignKey:DoctorId;references:Id"`
	Assigner *Manager `gorm:"foreignKey:AssignerId;references:Id"`
}
