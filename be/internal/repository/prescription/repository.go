package prescription

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_prescription_repository.go

type PrescriptionRepository interface {
	GetDB() *gorm.DB
	CreatePrescription(tx *gorm.DB, prescription *entity.Prescription) (*entity.Prescription, error)
	GetAllPrescription() ([]*entity.Prescription, error)
	GetPrescriptionById(prescriptionId int64) (*entity.Prescription, error)
	GetPrescriptionsByMedicalRecordId(medicalRecordId int64) ([]*entity.Prescription, error)
	GetPrescriptionsFromIds(prescriptionIds []int64) ([]*entity.Prescription, error)
	UpdatePrescription(tx *gorm.DB, prescription *entity.Prescription) (*entity.Prescription, error)
}
