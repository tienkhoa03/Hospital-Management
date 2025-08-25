package nurse

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_nurse_repository.go

type NurseRepository interface {
	GetDB() *gorm.DB
	CreateNurse(tx *gorm.DB, nurse *entity.Nurse) (*entity.Nurse, error)
	GetAllNurse() ([]*entity.Nurse, error)
	GetNurseById(nurseId int64) (*entity.Nurse, error)
	GetNurseByStaffId(staffId int64) (*entity.Nurse, error)
	GetNursesFromIds(nurseIds []int64) ([]*entity.Nurse, error)
	UpdateNurse(tx *gorm.DB, nurse *entity.Nurse) (*entity.Nurse, error)
}
