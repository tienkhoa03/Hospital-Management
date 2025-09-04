package dto

import (
	"time"
)

type TaskInfoRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BeginTime   time.Time `json:"begin_time"`
	FinishTime  time.Time `json:"finish_time"`
	Status      string    `json:"status" default:"scheduled"`
}

type TaskInfoResponse struct {
	TaskID      int64     `json:"task_id"`
	StaffUID    int64     `json:"staff_uid"`
	AssignerUID int64     `json:"assigner_uid"`
	Title       string    `json:"title"`
	Description *string   `json:"description,omitempty"`
	BeginTime   time.Time `json:"begin_time"`
	FinishTime  time.Time `json:"finish_time"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTaskInfoRequest struct {
	StaffUID    *int64     `json:"staff_uid"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	BeginTime   *time.Time `json:"begin_time"`
	FinishTime  *time.Time `json:"finish_time"`
	Status      *string    `json:"status" default:"scheduled"`
}
