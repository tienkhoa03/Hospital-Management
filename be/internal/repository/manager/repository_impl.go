package manager

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

type PostgreSQLManagerRepository struct {
	db *gorm.DB
}

func NewManagerRepository(db *gorm.DB) ManagerRepository {
	return &PostgreSQLManagerRepository{db: db}
}

func (r *PostgreSQLManagerRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLManagerRepository) GetAllManager() ([]*entity.Manager, error) {
	var managers = []*entity.Manager{}
	result := r.db.Model(&entity.Manager{}).Find(&managers)
	if result.Error != nil {
		return nil, result.Error
	}
	return managers, nil
}

func (r *PostgreSQLManagerRepository) GetManagerById(managerId int64) (*entity.Manager, error) {
	var manager = entity.Manager{}
	result := r.db.Model(&entity.Manager{}).Where("id = ?", managerId).First(&manager)
	if result.Error != nil {
		return nil, result.Error
	}
	return &manager, nil
}

func (r *PostgreSQLManagerRepository) GetManagersFromIds(managerIds []int64) ([]*entity.Manager, error) {
	var managers []*entity.Manager
	result := r.db.Model(&entity.Manager{}).Where("id IN ?", managerIds).Find(&managers)
	if result.Error != nil {
		return nil, result.Error
	}
	return managers, nil
}

func (r *PostgreSQLManagerRepository) GetManagerByUserId(userId int64) (*entity.Manager, error) {
	var manager = entity.Manager{}
	result := r.db.Model(&entity.Manager{}).Where("user_id = ?", userId).First(&manager)
	if result.Error != nil {
		return nil, result.Error
	}
	return &manager, nil
}

func (r *PostgreSQLManagerRepository) CreateManager(tx *gorm.DB, manager *entity.Manager) (*entity.Manager, error) {
	result := tx.Create(manager)
	if result.Error != nil {
		return nil, result.Error
	}
	return manager, result.Error
}

func (r *PostgreSQLManagerRepository) DeleteManagerByUserId(tx *gorm.DB, managerUID int64) error {
	result := tx.Model(&entity.Manager{}).Where("user_id = ?", managerUID).Update("status", constant.ManagerStatusInactive)
	return result.Error
}

func (r *PostgreSQLManagerRepository) UpdateManager(tx *gorm.DB, manager *entity.Manager) (*entity.Manager, error) {
	result := tx.Model(&entity.Manager{}).Where("id = ?", manager.Id).Updates(manager)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedManager = entity.Manager{}
	result = tx.First(&updatedManager, manager.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedManager, nil
}
