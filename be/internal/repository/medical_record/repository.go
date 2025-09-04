package medicalrecord

import (
	"BE_Hospital_Management/internal/domain/entity"
	"BE_Hospital_Management/internal/domain/filter"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_medical_record_repository.go

type MedicalRecordRepository interface {
	GetDB() *gorm.DB
	CreateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error)
	GetAllMedicalRecords() ([]*entity.MedicalRecord, error)
	GetMedicalRecordById(medicalRecordId int64) (*entity.MedicalRecord, error)
	GetMedicalRecordsByPatientId(patientId int64) ([]*entity.MedicalRecord, error)
	GetMedicalRecordsByDoctorId(doctorId int64) ([]*entity.MedicalRecord, error)
	GetMedicalRecordsFromIds(medicalRecordIds []int64) ([]*entity.MedicalRecord, error)
	UpdateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error)
	GetMedicalRecordsByPatientIdWithFilter(patientId int64, medicalRecordFilter *filter.MedicalRecordFilter) ([]*entity.MedicalRecord, error)
	GetMedicalRecordsByDoctorIdWithFilter(doctorId int64, medicalRecordFilter *filter.MedicalRecordFilter) ([]*entity.MedicalRecord, error)
}
