package patient

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_patient_repository.go

type PatientRepository interface {
	GetDB() *gorm.DB
	CreatePatient(tx *gorm.DB, patient *entity.Patient) (*entity.Patient, error)
	GetAllPatient() ([]*entity.Patient, error)
	GetPatientById(patientId int64) (*entity.Patient, error)
	GetPatientByIdWithUserInfo(patientId int64) (*entity.Patient, error)
	GetPatientByUserIdWithUserInfo(patientId int64) (*entity.Patient, error)
	GetPatientByUserId(userId int64) (*entity.Patient, error)
	GetPatientsFromIds(patientIds []int64) ([]*entity.Patient, error)
	GetPatientsFromIdsWithUserInfo(patientIds []int64) ([]*entity.Patient, error)
	UpdatePatient(tx *gorm.DB, patient *entity.Patient) (*entity.Patient, error)
	DeletePatientByUserId(tx *gorm.DB, patientId int64) error
}
