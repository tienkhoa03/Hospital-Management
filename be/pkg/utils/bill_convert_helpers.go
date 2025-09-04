package utils

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
)

func MapToBillResponse(bill *entity.Bill, billItems []*entity.BillItem, patientUID, doctorUID int64, cashingOfficerUID *int64) *dto.BillResponse {
	if bill == nil {
		return nil
	}
	var billMedicineResponses []dto.BillMedicineResponse
	for _, billItem := range billItems {
		if billItem != nil {
			billMedicineResponses = append(billMedicineResponses, dto.BillMedicineResponse{
				BillItemId:      billItem.Id,
				MedicineId:      billItem.MedicineId,
				MedicineName:    billItem.Medicine.Name,
				UsesInstruction: billItem.Medicine.UsesInstruction,
				Price:           billItem.Medicine.Price,
				Amount:          billItem.Amount,
			})
		}
	}
	return &dto.BillResponse{
		BillId:            bill.Id,
		PatientUID:        patientUID,
		DoctorUID:         doctorUID,
		CashingOfficerUID: cashingOfficerUID,
		MedicalRecordId:   bill.MedicalRecordId,
		TotalPrice:        bill.TotalPrice,
		Status:            bill.Status,
		CreatedAt:         bill.CreatedAt,
		UpdatedAt:         bill.UpdatedAt,
		Medicines:         billMedicineResponses,
	}
}
