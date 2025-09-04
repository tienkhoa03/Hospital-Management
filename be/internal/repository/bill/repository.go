package bill

import (
	"BE_Hospital_Management/internal/domain/entity"
	"BE_Hospital_Management/internal/domain/filter"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_bill_repository.go

type BillRepository interface {
	GetDB() *gorm.DB
	CreateBill(tx *gorm.DB, bill *entity.Bill) (*entity.Bill, error)
	GetAllBill() ([]*entity.Bill, error)
	GetBillById(billId int64) (*entity.Bill, error)
	GetBillsByPatientId(patientId int64) ([]*entity.Bill, error)
	GetBillsByCashingOfficerId(cashingOfficerId int64) ([]*entity.Bill, error)
	GetBillByMedicalRecordId(medicalRecordId int64) (*entity.Bill, error)
	GetBillsFromIds(billIds []int64) ([]*entity.Bill, error)
	UpdateBill(tx *gorm.DB, bill *entity.Bill) (*entity.Bill, error)
	GetBillsByPatientIdWithFilter(patientId int64, billFilter *filter.BillFilter) ([]*entity.Bill, error)
	GetBillsByCashingOfficerIdWithFilter(cashingOfficerId int64, billFilter *filter.BillFilter) ([]*entity.Bill, error)
}
