package nurse

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLNurseRepository struct {
	db *gorm.DB
}

func NewNurseRepository(db *gorm.DB) NurseRepository {
	return &PostgreSQLNurseRepository{db: db}
}

func (r *PostgreSQLNurseRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLNurseRepository) GetAllNurse() ([]*entity.Nurse, error) {
	var nurses = []*entity.Nurse{}
	result := r.db.Model(&entity.Nurse{}).Find(&nurses)
	if result.Error != nil {
		return nil, result.Error
	}
	return nurses, nil
}

func (r *PostgreSQLNurseRepository) GetNurseById(nurseId int64) (*entity.Nurse, error) {
	var nurse = entity.Nurse{}
	result := r.db.Model(&entity.Nurse{}).Where("id = ?", nurseId).First(&nurse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &nurse, nil
}

func (r *PostgreSQLNurseRepository) GetNurseByStaffId(staffId int64) (*entity.Nurse, error) {
	var nurse = entity.Nurse{}
	result := r.db.Model(&entity.Nurse{}).Where("staff_id = ?", staffId).First(&nurse)
	if result.Error != nil {
		return nil, result.Error
	}
	return &nurse, nil
}

func (r *PostgreSQLNurseRepository) GetNursesFromIds(nurseIds []int64) ([]*entity.Nurse, error) {
	var nurses []*entity.Nurse
	result := r.db.Model(&entity.Nurse{}).Where("id IN ?", nurseIds).Find(&nurses)
	if result.Error != nil {
		return nil, result.Error
	}
	return nurses, nil
}

func (r *PostgreSQLNurseRepository) CreateNurse(tx *gorm.DB, nurse *entity.Nurse) (*entity.Nurse, error) {
	result := tx.Create(nurse)
	if result.Error != nil {
		return nil, result.Error
	}
	return nurse, result.Error
}

func (r *PostgreSQLNurseRepository) DeleteNurseById(tx *gorm.DB, nurseId int64) error {
	result := tx.Model(&entity.Nurse{}).Where("id = ?", nurseId).Delete(entity.Nurse{})
	return result.Error
}

func (r *PostgreSQLNurseRepository) UpdateNurse(tx *gorm.DB, nurse *entity.Nurse) (*entity.Nurse, error) {
	result := tx.Model(&entity.Nurse{}).Where("id = ?", nurse.Id).Updates(nurse)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedNurse = entity.Nurse{}
	result = tx.First(&updatedNurse, nurse.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedNurse, nil
}
