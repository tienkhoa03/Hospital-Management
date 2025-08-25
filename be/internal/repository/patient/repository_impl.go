package patient

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLPatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &PostgreSQLPatientRepository{db: db}
}

func (r *PostgreSQLPatientRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLPatientRepository) GetAllPatient() ([]*entity.Patient, error) {
	var patients = []*entity.Patient{}
	result := r.db.Model(&entity.Patient{}).Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}

func (r *PostgreSQLPatientRepository) GetPatientById(patientId int64) (*entity.Patient, error) {
	var patient = entity.Patient{}
	result := r.db.Model(&entity.Patient{}).Where("id = ?", patientId).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

func (r *PostgreSQLPatientRepository) GetPatientByIdWithUserInfo(patientId int64) (*entity.Patient, error) {
	var patient = entity.Patient{}
	result := r.db.Model(&entity.Patient{}).Preload("User").Where("id = ?", patientId).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

func (r *PostgreSQLPatientRepository) GetPatientByUserId(userId int64) (*entity.Patient, error) {
	var patient = entity.Patient{}
	result := r.db.Model(&entity.Patient{}).Where("user_id = ?", userId).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

func (r *PostgreSQLPatientRepository) GetPatientByUserIdWithUserInfo(userId int64) (*entity.Patient, error) {
	var patient = entity.Patient{}
	result := r.db.Model(&entity.Patient{}).Preload("User").Where("user_id = ?", userId).First(&patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return &patient, nil
}

func (r *PostgreSQLPatientRepository) GetPatientsFromIds(patientIds []int64) ([]*entity.Patient, error) {
	var patients []*entity.Patient
	result := r.db.Model(&entity.Patient{}).Where("id IN ?", patientIds).Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}

func (r *PostgreSQLPatientRepository) GetPatientsFromIdsWithUserInfo(patientIds []int64) ([]*entity.Patient, error) {
	var patients []*entity.Patient
	result := r.db.Model(&entity.Patient{}).Preload("User").Where("id IN ?", patientIds).Find(&patients)
	if result.Error != nil {
		return nil, result.Error
	}
	return patients, nil
}

func (r *PostgreSQLPatientRepository) CreatePatient(tx *gorm.DB, patient *entity.Patient) (*entity.Patient, error) {
	result := tx.Create(patient)
	if result.Error != nil {
		return nil, result.Error
	}
	return patient, result.Error
}

func (r *PostgreSQLPatientRepository) DeletePatientById(tx *gorm.DB, patientId int64) error {
	result := tx.Model(&entity.Patient{}).Where("id = ?", patientId).Delete(entity.Patient{})
	return result.Error
}

func (r *PostgreSQLPatientRepository) UpdatePatient(tx *gorm.DB, patient *entity.Patient) (*entity.Patient, error) {
	result := tx.Model(&entity.Patient{}).Where("id = ?", patient.Id).Updates(patient)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedPatient = entity.Patient{}
	result = tx.First(&updatedPatient, patient.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedPatient, nil
}
