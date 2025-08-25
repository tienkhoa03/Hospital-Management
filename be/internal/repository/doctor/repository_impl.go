package doctor

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLDoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) DoctorRepository {
	return &PostgreSQLDoctorRepository{db: db}
}

func (r *PostgreSQLDoctorRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLDoctorRepository) GetAllDoctor() ([]*entity.Doctor, error) {
	var doctors = []*entity.Doctor{}
	result := r.db.Model(&entity.Doctor{}).Find(&doctors)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctors, nil
}

func (r *PostgreSQLDoctorRepository) GetDoctorById(doctorId int64) (*entity.Doctor, error) {
	var doctor = entity.Doctor{}
	result := r.db.Model(&entity.Doctor{}).Where("id = ?", doctorId).First(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}

func (r *PostgreSQLDoctorRepository) GetDoctorByStaffId(staffId int64) (*entity.Doctor, error) {
	var doctor = entity.Doctor{}
	result := r.db.Model(&entity.Doctor{}).Where("staff_id = ?", staffId).First(&doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &doctor, nil
}

func (r *PostgreSQLDoctorRepository) GetDoctorsFromIds(doctorIds []int64) ([]*entity.Doctor, error) {
	var doctors []*entity.Doctor
	result := r.db.Model(&entity.Doctor{}).Where("id IN ?", doctorIds).Find(&doctors)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctors, nil
}

func (r *PostgreSQLDoctorRepository) CreateDoctor(tx *gorm.DB, doctor *entity.Doctor) (*entity.Doctor, error) {
	result := tx.Create(doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	return doctor, result.Error
}

func (r *PostgreSQLDoctorRepository) DeleteDoctorById(tx *gorm.DB, doctorId int64) error {
	result := tx.Model(&entity.Doctor{}).Where("id = ?", doctorId).Delete(entity.Doctor{})
	return result.Error
}

func (r *PostgreSQLDoctorRepository) UpdateDoctor(tx *gorm.DB, doctor *entity.Doctor) (*entity.Doctor, error) {
	result := tx.Model(&entity.Doctor{}).Where("id = ?", doctor.Id).Updates(doctor)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedDoctor = entity.Doctor{}
	result = r.db.First(&updatedDoctor, doctor.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedDoctor, nil
}
