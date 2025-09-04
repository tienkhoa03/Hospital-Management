package billing

import (
	"BE_Hospital_Management/internal/domain/dto"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrBillNotFound = errors.New("bill not found")
	ErrNotPermitted = errors.New("you are not permitted to perform this action")
)

//go:generate mockgen -source=service.go -destination=../mock/mock_billing_service.go

type BillingService interface {
	UpdateBillStatusPaid(cashingOfficerUID, billId int64) (*dto.BillResponse, error)
	GetAllBills(userId int64, userRole string) ([]*dto.BillResponse, error)
	GetBillById(userId int64, userRole string, billId int64) (*dto.BillResponse, error)
}
