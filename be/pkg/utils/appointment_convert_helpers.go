package utils

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
)

func MapToAppointmentResponse(appointment *entity.Appointment, patientUID, doctorUID int64) *dto.AppointmentInfoResponse {
	return &dto.AppointmentInfoResponse{
		AppointmentId: appointment.Id,
		PatientUID:    patientUID,
		DoctorUID:     doctorUID,
		BeginTime:     appointment.BeginTime,
		FinishTime:    appointment.FinishTime,
		Status:        appointment.Status,
		CreatedAt:     appointment.CreatedAt,
		UpdatedAt:     appointment.UpdatedAt,
	}
}
