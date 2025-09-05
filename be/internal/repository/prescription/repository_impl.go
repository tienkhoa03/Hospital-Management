package prescription

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLPrescriptionRepository struct {
	db *gorm.DB
}

func NewPrescriptionRepository(db *gorm.DB) PrescriptionRepository {
	return &PostgreSQLPrescriptionRepository{db: db}
}

func (r *PostgreSQLPrescriptionRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLPrescriptionRepository) GetAllPrescription() ([]*entity.Prescription, error) {
	var prescriptions = []*entity.Prescription{}
	result := r.db.Model(&entity.Prescription{}).Preload("Medicine").Find(&prescriptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return prescriptions, nil
}

func (r *PostgreSQLPrescriptionRepository) GetPrescriptionById(prescriptionId int64) (*entity.Prescription, error) {
	var prescription = entity.Prescription{}
	result := r.db.Model(&entity.Prescription{}).Preload("Medicine").Where("id = ?", prescriptionId).First(&prescription)
	if result.Error != nil {
		return nil, result.Error
	}
	return &prescription, nil
}

func (r *PostgreSQLPrescriptionRepository) GetPrescriptionsByMedicalRecordId(medicalRecordId int64) ([]*entity.Prescription, error) {
	var prescriptions []*entity.Prescription
	result := r.db.Model(&entity.Prescription{}).Preload("Medicine").Where("medical_record_id = ?", medicalRecordId).Find(&prescriptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return prescriptions, nil
}

func (r *PostgreSQLPrescriptionRepository) GetPrescriptionsFromIds(prescriptionIds []int64) ([]*entity.Prescription, error) {
	var prescriptions []*entity.Prescription
	result := r.db.Model(&entity.Prescription{}).Preload("Medicine").Where("id IN ?", prescriptionIds).Find(&prescriptions)
	if result.Error != nil {
		return nil, result.Error
	}
	return prescriptions, nil
}

func (r *PostgreSQLPrescriptionRepository) CreatePrescription(tx *gorm.DB, prescription *entity.Prescription) (*entity.Prescription, error) {
	result := tx.Create(prescription)
	if result.Error != nil {
		return nil, result.Error
	}
	var newPrescription = &entity.Prescription{}
	result = tx.Model(&entity.Prescription{}).Preload("Medicine").Where("id = ?", prescription.Id).First(newPrescription)
	if result.Error != nil {
		return nil, result.Error
	}
	return newPrescription, nil
}

func (r *PostgreSQLPrescriptionRepository) UpdatePrescription(tx *gorm.DB, prescription *entity.Prescription) (*entity.Prescription, error) {
	result := tx.Model(&entity.Prescription{}).Preload("Medicine").Where("id = ?", prescription.Id).Updates(prescription)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedPrescription = entity.Prescription{}
	result = tx.First(&updatedPrescription, prescription.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedPrescription, nil
}
