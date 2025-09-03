package billitem

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLBillItemRepository struct {
	db *gorm.DB
}

func NewBillItemRepository(db *gorm.DB) BillItemRepository {
	return &PostgreSQLBillItemRepository{db: db}
}

func (r *PostgreSQLBillItemRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLBillItemRepository) GetAllBillItem() ([]*entity.BillItem, error) {
	var billItems = []*entity.BillItem{}
	result := r.db.Model(&entity.BillItem{}).Find(&billItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return billItems, nil
}

func (r *PostgreSQLBillItemRepository) GetBillItemById(billItemId int64) (*entity.BillItem, error) {
	var billItem = entity.BillItem{}
	result := r.db.Model(&entity.BillItem{}).Where("id = ?", billItemId).First(&billItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &billItem, nil
}

func (r *PostgreSQLBillItemRepository) GetBillItemsFromIds(billItemIds []int64) ([]*entity.BillItem, error) {
	var billItems []*entity.BillItem
	result := r.db.Model(&entity.BillItem{}).Where("id IN ?", billItemIds).Find(&billItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return billItems, nil
}

func (r *PostgreSQLBillItemRepository) CreateBillItem(tx *gorm.DB, billItem *entity.BillItem) (*entity.BillItem, error) {
	result := tx.Create(billItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return billItem, result.Error
}

func (r *PostgreSQLBillItemRepository) UpdateBillItem(tx *gorm.DB, billItem *entity.BillItem) (*entity.BillItem, error) {
	result := tx.Model(&entity.BillItem{}).Where("id = ?", billItem.Id).Updates(billItem)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedBillItem = entity.BillItem{}
	result = tx.First(&updatedBillItem, billItem.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedBillItem, nil
}
