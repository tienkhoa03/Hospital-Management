package dto

import (
	"time"
)

type TaskInfoRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BeginTime   time.Time `json:"begin_time"`
	FinishTime  time.Time `json:"finish_time"`
	Status      string    `json:"status"`
}
