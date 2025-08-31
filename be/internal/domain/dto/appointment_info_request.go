package dto

import (
	"time"
)

type AppointmentInfoRequest struct {
	PatientUID *int64    `json:"patient_uid"`
	DoctorUID  *int64    `json:"doctor_uid"`
	BeginTime  time.Time `json:"begin_time"`
	FinishTime time.Time `json:"finish_time"`
	Status     string    `json:"status" default:"scheduled"`
}

type UpdateAppointmentRequest struct {
	DoctorUID  *int64    `json:"doctor_uid"`
	BeginTime  time.Time `json:"begin_time"`
	FinishTime time.Time `json:"finish_time"`
	Status     string    `json:"status"`
}
