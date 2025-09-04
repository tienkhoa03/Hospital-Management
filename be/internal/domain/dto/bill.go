package dto

import "time"

type BillResponse struct {
	BillId            int64                  `json:"bill_id"`
	PatientUID        int64                  `json:"patient_uid"`
	DoctorUID         int64                  `json:"doctor_uid"`
	CashingOfficerUID *int64                 `json:"cashing_officer_uid"`
	MedicalRecordId   int64                  `json:"medical_record_id"`
	TotalPrice        float32                `json:"total_price"`
	Status            string                 `json:"status"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	Medicines         []BillMedicineResponse `json:"medicines"`
}

type BillMedicineResponse struct {
	BillItemId      int64   `json:"bill_item_id"`
	MedicineId      int64   `json:"medicine_id"`
	MedicineName    string  `json:"medicine_name"`
	UsesInstruction string  `json:"uses_instruction"`
	Price           float32 `json:"price"`
	Amount          int     `json:"amount"`
}
