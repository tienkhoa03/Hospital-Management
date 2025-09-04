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

type AppointmentInfoResponse struct {
	AppointmentId int64     `json:"appointment_id"`
	PatientUID    int64     `json:"patient_uid"`
	DoctorUID     int64     `json:"doctor_uid"`
	BeginTime     time.Time `json:"begin_time"`
	FinishTime    time.Time `json:"finish_time"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UpdateAppointmentRequest struct {
	DoctorUID  *int64     `json:"doctor_uid"`
	BeginTime  *time.Time `json:"begin_time"`
	FinishTime *time.Time `json:"finish_time"`
	Status     *string    `json:"status"`
}

type AppointmentSlot struct {
	BeginTime  time.Time `json:"begin_time"`
	FinishTime time.Time `json:"finish_time"`
}

type IsAvailableSlotResponse struct {
	IsAvailable bool `json:"is_available"`
}
