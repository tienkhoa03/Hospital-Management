package appointment

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/entity"
	"time"

	"gorm.io/gorm"
)

type PostgreSQLAppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
	return &PostgreSQLAppointmentRepository{db: db}
}

func (r *PostgreSQLAppointmentRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLAppointmentRepository) GetAllAppointment() ([]*entity.Appointment, error) {
	var appointments = []*entity.Appointment{}
	result := r.db.Model(&entity.Appointment{}).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (r *PostgreSQLAppointmentRepository) GetAppointmentById(appointmentId int64) (*entity.Appointment, error) {
	var appointment = entity.Appointment{}
	result := r.db.Model(&entity.Appointment{}).Where("id = ?", appointmentId).First(&appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &appointment, nil
}

func (r *PostgreSQLAppointmentRepository) GetAppointmentsByDoctorId(doctorId int64) ([]*entity.Appointment, error) {
	var appointments []*entity.Appointment
	result := r.db.Model(&entity.Appointment{}).Where("doctor_id = ?", doctorId).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (r *PostgreSQLAppointmentRepository) GetAppointmentsByPatientId(patientId int64) ([]*entity.Appointment, error) {
	var appointments []*entity.Appointment
	result := r.db.Model(&entity.Appointment{}).Where("patient_id = ?", patientId).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (r *PostgreSQLAppointmentRepository) GetPatientIdsByDoctorId(doctorId int64) ([]int64, error) {
	var appointment []int64
	result := r.db.Model(&entity.Appointment{}).Where("doctor_id = ?", doctorId).Pluck("patient_id", &appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointment, nil
}

func (r *PostgreSQLAppointmentRepository) GetAppointmentByPatientIdAndDoctorId(patientId, doctorId int64) (*entity.Appointment, error) {
	var appointment = entity.Appointment{}
	result := r.db.Model(&entity.Appointment{}).Where("patient_id = ? AND doctor_id = ?", patientId, doctorId).First(&appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	return &appointment, nil
}

func (r *PostgreSQLAppointmentRepository) GetAppointmentsFromIds(appointmentIds []int64) ([]*entity.Appointment, error) {
	var appointments []*entity.Appointment
	result := r.db.Model(&entity.Appointment{}).Where("id IN ?", appointmentIds).Find(&appointments)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointments, nil
}

func (r *PostgreSQLAppointmentRepository) CreateAppointment(tx *gorm.DB, appointment *entity.Appointment) (*entity.Appointment, error) {
	result := tx.Create(appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	return appointment, result.Error
}

func (r *PostgreSQLAppointmentRepository) DeleteAppointmentById(tx *gorm.DB, appointmentId int64) error {
	result := tx.Model(&entity.Appointment{}).Where("id = ?", appointmentId).Update("status", constant.AppointmentStatusCanceled)
	return result.Error
}

func (r *PostgreSQLAppointmentRepository) UpdateAppointment(tx *gorm.DB, appointment *entity.Appointment) (*entity.Appointment, error) {
	result := tx.Model(&entity.Appointment{}).Where("id = ?", appointment.Id).Updates(appointment)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedAppointment = entity.Appointment{}
	result = r.db.First(&updatedAppointment, appointment.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedAppointment, nil
}

func (r *PostgreSQLAppointmentRepository) ExistsOverlapAppointmentOfDoctor(doctorId int64, beginTime, endTime time.Time) (bool, error) {
	var appointment entity.Appointment
	err := r.db.Where("doctor_id = ?", doctorId).Where("finish_time > ? AND begin_time < ? AND status = ?", beginTime, endTime, constant.AppointmentStatusScheduled).Limit(1).Take(&appointment).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
