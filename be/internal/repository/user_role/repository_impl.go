package userrole

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

type PostgreSQLUserRoleRepository struct {
	db *gorm.DB
}

func NewUserRoleRepository(db *gorm.DB) UserRoleRepository {
	return &PostgreSQLUserRoleRepository{db: db}
}

func (r *PostgreSQLUserRoleRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLUserRoleRepository) GetAllUserRole() ([]*entity.UserRole, error) {
	var userRoles = []*entity.UserRole{}
	result := r.db.Model(&entity.UserRole{}).Find(&userRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return userRoles, nil
}

func (r *PostgreSQLUserRoleRepository) GetUserRoleById(userRoleId int64) (*entity.UserRole, error) {
	var userRole = entity.UserRole{}
	result := r.db.Model(&entity.UserRole{}).Where("id = ?", userRoleId).First(&userRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userRole, nil
}

func (r *PostgreSQLUserRoleRepository) GetUserRolesFromIds(userRoleIds []int64) ([]*entity.UserRole, error) {
	var userRoles []*entity.UserRole
	result := r.db.Model(&entity.UserRole{}).Where("id IN ?", userRoleIds).Find(&userRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return userRoles, nil
}

func (r *PostgreSQLUserRoleRepository) GetUserRoleBySlug(userRoleSlug string) (*entity.UserRole, error) {
	var userRole = entity.UserRole{}
	result := r.db.Model(&entity.UserRole{}).Where("role_slug = ?", userRoleSlug).First(&userRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return &userRole, nil
}

func (r *PostgreSQLUserRoleRepository) GetUserRolesFromSlugs(userRoleSlugs []string) ([]*entity.UserRole, error) {
	var userRoles []*entity.UserRole
	result := r.db.Model(&entity.UserRole{}).Where("role_slug IN ?", userRoleSlugs).Find(&userRoles)
	if result.Error != nil {
		return nil, result.Error
	}
	return userRoles, nil
}

func (r *PostgreSQLUserRoleRepository) CreateUserRole(tx *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error) {
	result := tx.Create(userRole)
	if result.Error != nil {
		return nil, result.Error
	}
	return userRole, result.Error
}

func (r *PostgreSQLUserRoleRepository) DeleteUserRoleById(tx *gorm.DB, userRoleId int64) error {
	result := tx.Model(&entity.UserRole{}).Where("id = ?", userRoleId).Delete(entity.UserRole{})
	return result.Error
}

func (r *PostgreSQLUserRoleRepository) UpdateUserRole(tx *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error) {
	result := tx.Model(&entity.UserRole{}).Where("id = ?", userRole.Id).Updates(userRole)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedUserRole = entity.UserRole{}
	result = tx.First(&updatedUserRole, userRole.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedUserRole, nil
}
