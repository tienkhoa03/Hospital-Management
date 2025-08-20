package userrole

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_userRole_repository.go

type UserRoleRepository interface {
	GetDB() *gorm.DB
	CreateUserRole(tx *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error)
	GetAllUserRole() ([]*entity.UserRole, error)
	GetUserRoleById(userRoleId int64) (*entity.UserRole, error)
	GetUserRolesFromIds(userRoleIds []int64) ([]*entity.UserRole, error)
	GetUserRoleBySlug(userRoleSlug string) (*entity.UserRole, error)
	GetUserRolesFromSlugs(userRoleSlugs []string) ([]*entity.UserRole, error)
	UpdateUserRole(tx *gorm.DB, userRole *entity.UserRole) (*entity.UserRole, error)
	DeleteUserRoleById(tx *gorm.DB, userRoleId int64) error
}
