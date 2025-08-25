package staff

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_staff_repository.go

type StaffRepository interface {
	GetDB() *gorm.DB
	CreateStaff(tx *gorm.DB, staff *entity.Staff) (*entity.Staff, error)
	GetAllStaff() ([]*entity.Staff, error)
	GetStaffById(staffId int64) (*entity.Staff, error)
	GetStaffsFromIds(staffIds []int64) ([]*entity.Staff, error)
	GetStaffsByManagerIdWithInformation(managerId int64) ([]*entity.Staff, error)
	GetStaffsByUserIdWithInformation(userId int64) (*entity.Staff, error)
	GetStaffByUserId(userId int64) (*entity.Staff, error)
	UpdateStaff(tx *gorm.DB, staff *entity.Staff) (*entity.Staff, error)
	DeleteStaffByUserId(tx *gorm.DB, staffId int64) error
}
