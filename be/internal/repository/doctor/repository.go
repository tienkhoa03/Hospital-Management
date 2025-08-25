package doctor

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_doctor_repository.go

type DoctorRepository interface {
	GetDB() *gorm.DB
	CreateDoctor(tx *gorm.DB, doctor *entity.Doctor) (*entity.Doctor, error)
	GetAllDoctor() ([]*entity.Doctor, error)
	GetDoctorById(doctorId int64) (*entity.Doctor, error)
	GetDoctorByStaffId(staffId int64) (*entity.Doctor, error)
	GetDoctorsFromIds(doctorIds []int64) ([]*entity.Doctor, error)
	UpdateDoctor(tx *gorm.DB, doctor *entity.Doctor) (*entity.Doctor, error)
	DeleteDoctorById(tx *gorm.DB, doctorId int64) error
}
