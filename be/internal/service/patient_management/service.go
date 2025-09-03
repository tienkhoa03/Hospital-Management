package patientmanagement

import (
	"BE_Hospital_Management/internal/domain/dto"
	"errors"
)

var (
	ErrDoctorNotWorking    = errors.New("doctor is not working")
	ErrUserNotFound        = errors.New("user not found")
	ErrAppointmentNotFound = errors.New("appointment not found")
	ErrNotPermitted        = errors.New("you are not permitted to perform this action")
	ErrInvalidTimeRange    = errors.New("invalid time range")
	ErrMissingDoctorId     = errors.New("missing doctor id")
	ErrMissingPatientId    = errors.New("missing patient id")
	ErrOutOfWorkingHours   = errors.New("time is out of working hours")
)

//go:generate mockgen -source=service.go -destination=../mock/mock_patient_management_service.go

type PatientManagementService interface {
	CreateTreatmentPlan(doctorUID int64, treatmentPlan dto.TreatmentPlanRequest) (*dto.TreatmentPlanResponse, error)
}
