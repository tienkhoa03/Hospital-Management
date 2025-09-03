package appointment

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	"errors"
	"time"
)

var (
	ErrDoctorNotWorking         = errors.New("doctor is not working")
	ErrUserNotFound             = errors.New("user not found")
	ErrAppointmentNotFound      = errors.New("appointment not found")
	ErrNotPermitted             = errors.New("you are not permitted to perform this action")
	ErrInvalidTimeRange         = errors.New("invalid time range")
	ErrExistsOverlapTask        = errors.New("there is an overlapping task for the staff in the given time range")
	ErrExistsOverlapAppointment = errors.New("there is an overlapping appointment for the doctor in the given time range")
	ErrMissingDoctorId          = errors.New("missing doctor id")
	ErrMissingPatientId         = errors.New("missing patient id")
	ErrOutOfWorkingHours        = errors.New("time is out of working hours")
)

//go:generate mockgen -source=service.go -destination=../mock/mock_appointment_service.go

type AppointmentService interface {
	CreateAppointment(authUserId int64, authUserRole string, appointmentRequest *dto.AppointmentInfoRequest) (*entity.Appointment, error)
	UpdateAppointment(patientUID, appointmentId int64, request *dto.UpdateAppointmentRequest) (*entity.Appointment, error)
	DeleteAppointment(requestorUID int64, requestorRole string, appointmentId int64) error
	GetAvailableSlots(doctorUID int64, date time.Time) ([]*dto.AppointmentSlot, error)
	CheckAvailableSlot(doctorUID int64, beginTime, finishTime time.Time) (bool, error)
	GetAllAppointments(authUserId int64, authUserRole string) ([]*entity.Appointment, error)
	GetAppointmentById(authUserId int64, authUserRole string, appointmentId int64) (*entity.Appointment, error)
}
