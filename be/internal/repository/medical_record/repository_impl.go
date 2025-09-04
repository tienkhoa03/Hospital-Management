package medicalrecord

import (
	"BE_Hospital_Management/internal/domain/entity"
	"BE_Hospital_Management/internal/domain/filter"

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

func (r *PostgreSQLMedicalRecordRepository) GetAllMedicalRecords() ([]*entity.MedicalRecord, error) {
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

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordsByPatientId(patientId int64) ([]*entity.MedicalRecord, error) {
	var medicalRecord []*entity.MedicalRecord
	result := r.db.Model(&entity.MedicalRecord{}).Where("patient_id = ?", patientId).Find(&medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecord, nil
}

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordsByDoctorId(doctorId int64) ([]*entity.MedicalRecord, error) {
	var medicalRecord []*entity.MedicalRecord
	result := r.db.Model(&entity.MedicalRecord{}).Where("doctor_id = ?", doctorId).Find(&medicalRecord)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecord, nil
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

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordsByPatientIdWithFilter(patientId int64, medicalRecordFilter *filter.MedicalRecordFilter) ([]*entity.MedicalRecord, error) {
	var medicalRecords []*entity.MedicalRecord
	db := r.db.Model(&entity.MedicalRecord{}).Where("patient_id = ?", patientId)
	db = medicalRecordFilter.ApplyFilter(db)
	result := db.Find(&medicalRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecords, nil
}

func (r *PostgreSQLMedicalRecordRepository) GetMedicalRecordsByDoctorIdWithFilter(doctorId int64, medicalRecordFilter *filter.MedicalRecordFilter) ([]*entity.MedicalRecord, error) {
	var medicalRecords []*entity.MedicalRecord
	db := r.db.Model(&entity.MedicalRecord{}).Where("doctor_id = ?", doctorId)
	db = medicalRecordFilter.ApplyFilter(db)
	result := db.Find(&medicalRecords)
	if result.Error != nil {
		return nil, result.Error
	}
	return medicalRecords, nil
}
