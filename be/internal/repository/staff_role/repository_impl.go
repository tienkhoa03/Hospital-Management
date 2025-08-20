package staffrole

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

type PostgreSQLStaffRoleRepository struct {
	db *gorm.DB
}

func NewStaffRoleRepository(db *gorm.DB) StaffRoleRepository {
	return &PostgreSQLStaffRoleRepository{db: db}
}

func (r *PostgreSQLStaffRoleRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLStaffRoleRepository) GetAllStaffRole() ([]*entity.StaffRole, error) {
	var staffRoles = []*entity.StaffRole{}
	result := r.db.Model(&entity.StaffRole{}).Find(&staffRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffRoles, nil
}

func (r *PostgreSQLStaffRoleRepository) GetStaffRoleById(staffRoleId int64) (*entity.StaffRole, error) {
	var staffRole = entity.StaffRole{}
	result := r.db.Model(&entity.StaffRole{}).Where("id = ?", staffRoleId).First(&staffRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staffRole, nil
}

func (r *PostgreSQLStaffRoleRepository) GetStaffRolesFromIds(staffRoleIds []int64) ([]*entity.StaffRole, error) {
	var staffRoles []*entity.StaffRole
	result := r.db.Model(&entity.StaffRole{}).Where("id IN ?", staffRoleIds).Find(&staffRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffRoles, nil
}

func (r *PostgreSQLStaffRoleRepository) GetStaffRoleBySlug(staffRoleSlug string) (*entity.StaffRole, error) {
	var staffRole = entity.StaffRole{}
	result := r.db.Model(&entity.StaffRole{}).Where("role_slug = ?", staffRoleSlug).First(&staffRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return &staffRole, nil
}

func (r *PostgreSQLStaffRoleRepository) GetStaffRolesFromSlugs(staffRoleSlugs []string) ([]*entity.StaffRole, error) {
	var staffRoles []*entity.StaffRole
	result := r.db.Model(&entity.StaffRole{}).Where("role_slug IN ?", staffRoleSlugs).Find(&staffRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffRoles, nil
}

func (r *PostgreSQLStaffRoleRepository) CreateStaffRole(tx *gorm.DB, staffRole *entity.StaffRole) (*entity.StaffRole, error) {
	result := tx.Create(staffRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return staffRole, result.Error
}

func (r *PostgreSQLStaffRoleRepository) DeleteStaffRoleById(tx *gorm.DB, staffRoleId int64) error {
	result := tx.Model(&entity.StaffRole{}).Where("id = ?", staffRoleId).Delete(entity.StaffRole{})
	return result.Error
}

func (r *PostgreSQLStaffRoleRepository) UpdateStaffRole(tx *gorm.DB, staffRole *entity.StaffRole) (*entity.StaffRole, error) {
	result := tx.Model(&entity.StaffRole{}).Where("id = ?", staffRole.Id).Updates(staffRole)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedStaffRole = entity.StaffRole{}
	result = tx.First(&updatedStaffRole, staffRole.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedStaffRole, nil
}
