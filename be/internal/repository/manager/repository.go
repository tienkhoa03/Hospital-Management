package manager

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_manager_repository.go

type ManagerRepository interface {
	GetDB() *gorm.DB
	CreateManager(tx *gorm.DB, manager *entity.Manager) (*entity.Manager, error)
	GetAllManager() ([]*entity.Manager, error)
	GetManagerById(managerId int64) (*entity.Manager, error)
	GetManagersFromIds(managerIds []int64) ([]*entity.Manager, error)
	GetManagerByUserId(userId int64) (*entity.Manager, error)
	UpdateManager(tx *gorm.DB, manager *entity.Manager) (*entity.Manager, error)
	DeleteManagerById(tx *gorm.DB, managerId int64) error
}
