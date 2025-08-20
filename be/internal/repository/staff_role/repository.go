package staffrole

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_staffRole_repository.go

type StaffRoleRepository interface {
	GetDB() *gorm.DB
	CreateStaffRole(tx *gorm.DB, staffRole *entity.StaffRole) (*entity.StaffRole, error)
	GetAllStaffRole() ([]*entity.StaffRole, error)
	GetStaffRoleById(staffRoleId int64) (*entity.StaffRole, error)
	GetStaffRolesFromIds(staffRoleIds []int64) ([]*entity.StaffRole, error)
	GetStaffRoleBySlug(staffRoleSlug string) (*entity.StaffRole, error)
	GetStaffRolesFromSlugs(staffRoleSlugs []string) ([]*entity.StaffRole, error)
	UpdateStaffRole(tx *gorm.DB, staffRole *entity.StaffRole) (*entity.StaffRole, error)
	DeleteStaffRoleById(tx *gorm.DB, staffRoleId int64) error
}
