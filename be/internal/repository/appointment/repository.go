package appointment

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
	"time"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_appointment_repository.go

type AppointmentRepository interface {
	GetDB() *gorm.DB
	CreateAppointment(tx *gorm.DB, appointment *entity.Appointment) (*entity.Appointment, error)
	GetAllAppointment() ([]*entity.Appointment, error)
	GetAppointmentById(appointmentId int64) (*entity.Appointment, error)
	GetAppointmentsByPatientId(patientId int64) ([]*entity.Appointment, error)
	GetAppointmentsByDoctorId(doctorId int64) ([]*entity.Appointment, error)
	GetAppointmentByPatientIdAndDoctorId(patientId, doctorId int64) (*entity.Appointment, error)
	GetPatientIdsByDoctorId(staffId int64) ([]int64, error)
	GetAppointmentsFromIds(appointmentIds []int64) ([]*entity.Appointment, error)
	UpdateAppointment(tx *gorm.DB, appointment *entity.Appointment) (*entity.Appointment, error)
	DeleteAppointmentById(tx *gorm.DB, appointmentId int64) error
	ExistsOverlapAppointmentOfDoctor(doctorId int64, beginTime, endTime time.Time) (bool, error)
}
