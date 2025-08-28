package staff

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLStaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &PostgreSQLStaffRepository{db: db}
}

func (r *PostgreSQLStaffRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLStaffRepository) GetAllStaff() ([]*entity.Staff, error) {
	var staffs = []*entity.Staff{}
	result := r.db.Model(&entity.Staff{}).Find(&staffs)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffs, nil
}

func (r *PostgreSQLStaffRepository) GetStaffById(staffId int64) (*entity.Staff, error) {
	var staff = entity.Staff{}
	result := r.db.Model(&entity.Staff{}).Preload("Role").Where("id = ?", staffId).First(&staff)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staff, nil
}

func (r *PostgreSQLStaffRepository) GetStaffsFromIds(staffIds []int64) ([]*entity.Staff, error) {
	var staffs []*entity.Staff
	result := r.db.Model(&entity.Staff{}).Preload("Role").Where("id IN ?", staffIds).Find(&staffs)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffs, nil
}

func (r *PostgreSQLStaffRepository) GetStaffByUserId(userId int64) (*entity.Staff, error) {
	var staff = entity.Staff{}
	result := r.db.Model(&entity.Staff{}).Preload("Role").Where("user_id = ?", userId).First(&staff)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staff, nil
}

func (r *PostgreSQLStaffRepository) GetStaffsByManagerIdWithInformation(managerId int64) ([]*entity.Staff, error) {
	var staffs []*entity.Staff
	result := r.db.Model(&entity.Staff{}).Joins("User").Joins("Nurse").Joins("Doctor").Preload("Role").Where("manage_by = ?", managerId).Find(&staffs)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffs, nil
}
func (r *PostgreSQLStaffRepository) GetStaffsByUserIdWithInformation(userId int64) (*entity.Staff, error) {
	var staff entity.Staff
	result := r.db.Model(&entity.Staff{}).Joins("User").Joins("Nurse").Joins("Doctor").Preload("Role").Where("user_id = ?", userId).First(&staff)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staff, nil
}

func (r *PostgreSQLStaffRepository) CreateStaff(tx *gorm.DB, staff *entity.Staff) (*entity.Staff, error) {
	result := tx.Create(staff)
	if result.Error != nil {
		return nil, result.Error
	}
	return staff, result.Error
}

func (r *PostgreSQLStaffRepository) DeleteStaffByUserId(tx *gorm.DB, staffUID int64) error {
	result := tx.Model(&entity.Staff{}).Where("user_id = ?", staffUID).Update("status", constant.StaffStatusInactive)
	return result.Error
}

func (r *PostgreSQLStaffRepository) UpdateStaff(tx *gorm.DB, staff *entity.Staff) (*entity.Staff, error) {
	result := tx.Model(&entity.Staff{}).Where("id = ?", staff.Id).Updates(staff)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedStaff = entity.Staff{}
	result = tx.First(&updatedStaff, staff.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedStaff, nil
}
