package billitem

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_bill_item_repository.go

type BillItemRepository interface {
	GetDB() *gorm.DB
	CreateBillItem(tx *gorm.DB, billItem *entity.BillItem) (*entity.BillItem, error)
	GetAllBillItem() ([]*entity.BillItem, error)
	GetBillItemById(billItemId int64) (*entity.BillItem, error)
	GetBillItemsFromIds(billItemIds []int64) ([]*entity.BillItem, error)
	UpdateBillItem(tx *gorm.DB, billItem *entity.BillItem) (*entity.BillItem, error)
}
