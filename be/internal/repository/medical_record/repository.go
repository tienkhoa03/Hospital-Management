package medicalrecord

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_medical_record_repository.go

type MedicalRecordRepository interface {
	GetDB() *gorm.DB
	CreateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error)
	GetAllMedicalRecord() ([]*entity.MedicalRecord, error)
	GetMedicalRecordById(medicalRecordId int64) (*entity.MedicalRecord, error)
	GetMedicalRecordByPatientId(patientId int64) ([]*entity.MedicalRecord, error)
	GetMedicalRecordByDoctorId(doctorId int64) ([]*entity.MedicalRecord, error)
	GetMedicalRecordsFromIds(medicalRecordIds []int64) ([]*entity.MedicalRecord, error)
	UpdateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error)
}
