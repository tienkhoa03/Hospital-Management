package medicalrecord

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLMedicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) MedicalRecordRepository {
	return &PostgreSQLMedicalRecordRepository{db: db}
}

func (r *PostgreSQLMedicalRecordRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLMedicalRecordRepository) GetAllMedicalRecord() ([]*entity.MedicalRecord, error) {
	var medicalRecords = []*entity.MedicalRecord{}
	result := r.db.Model(&entity.MedicalRecord{}).Find(&medicalRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecords, nil
}

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordById(medicalRecordId int64) (*entity.MedicalRecord, error) {
	var medicalRecord = entity.MedicalRecord{}
	result := r.db.Model(&entity.MedicalRecord{}).Where("id = ?", medicalRecordId).First(&medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	return &medicalRecord, nil
}

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordsFromIds(medicalRecordIds []int64) ([]*entity.MedicalRecord, error) {
	var medicalRecords []*entity.MedicalRecord
	result := r.db.Model(&entity.MedicalRecord{}).Where("id IN ?", medicalRecordIds).Find(&medicalRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecords, nil
}

func (r *PostgreSQLMedicalRecordRepository) CreateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error) {
	result := tx.Create(medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecord, result.Error
}

func (r *PostgreSQLMedicalRecordRepository) UpdateMedicalRecord(tx *gorm.DB, medicalRecord *entity.MedicalRecord) (*entity.MedicalRecord, error) {
	result := tx.Model(&entity.MedicalRecord{}).Where("id = ?", medicalRecord.Id).Updates(medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedMedicalRecord = entity.MedicalRecord{}
	result = tx.First(&updatedMedicalRecord, medicalRecord.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedMedicalRecord, nil
}
