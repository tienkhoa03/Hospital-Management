package bill

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLBillRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &PostgreSQLBillRepository{db: db}
}

func (r *PostgreSQLBillRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLBillRepository) GetAllBill() ([]*entity.Bill, error) {
	var bills = []*entity.Bill{}
	result := r.db.Model(&entity.Bill{}).Find(&bills)
	if result.Error != nil {
		return nil, result.Error
	}
	return bills, nil
}

func (r *PostgreSQLBillRepository) GetBillById(billId int64) (*entity.Bill, error) {
	var bill = entity.Bill{}
	result := r.db.Model(&entity.Bill{}).Where("id = ?", billId).First(&bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bill, nil
}

func (r *PostgreSQLBillRepository) GetBillByMedicalRecordId(medicalRecordId int64) (*entity.Bill, error) {
	var bill entity.Bill
	result := r.db.Model(&entity.Bill{}).Where("medical_record_id = ?", medicalRecordId).First(&bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return &bill, nil
}

func (r *PostgreSQLBillRepository) GetBillsFromIds(billIds []int64) ([]*entity.Bill, error) {
	var bills []*entity.Bill
	result := r.db.Model(&entity.Bill{}).Where("id IN ?", billIds).Find(&bills)
	if result.Error != nil {
		return nil, result.Error
	}
	return bills, nil
}

func (r *PostgreSQLBillRepository) CreateBill(tx *gorm.DB, bill *entity.Bill) (*entity.Bill, error) {
	result := tx.Create(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	return bill, result.Error
}

func (r *PostgreSQLBillRepository) UpdateBill(tx *gorm.DB, bill *entity.Bill) (*entity.Bill, error) {
	result := tx.Model(&entity.Bill{}).Where("id = ?", bill.Id).Updates(bill)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedBill = entity.Bill{}
	result = tx.First(&updatedBill, bill.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedBill, nil
}
