package dto

type TreatmentPlanRequest struct {
	PatientUID      int64                 `json:"patient_uid"`
	AppointmentId   *int64                `json:"appointment_id"`
	Symptom         string                `json:"symptom"`
	MedicalServices *string               `json:"medical_services"`
	Diagnosis       string                `json:"diagnosis"`
	Treatment       string                `json:"treatment"`
	Note            *string               `json:"note"`
	Prescriptions   []PrescriptionRequest `json:"prescriptions"`
}

type PrescriptionRequest struct {
	MedicineId  int64  `json:"medicine_id"`
	Instruction string `json:"instruction"`
	Amount      int    `json:"amount"`
}

type TreatmentPlanResponse struct {
	MedicalRecordId int64                  `json:"treatment_plan_id"`
	PatientUID      int64                  `json:"patient_uid"`
	DoctorUID       int64                  `json:"doctor_uid"`
	AppointmentId   *int64                 `json:"appointment_id"`
	Symptom         string                 `json:"symptom"`
	MedicalServices *string                `json:"medical_services"`
	Diagnosis       string                 `json:"diagnosis"`
	Treatment       string                 `json:"treatment"`
	Note            *string                `json:"note"`
	Prescriptions   []PrescriptionResponse `json:"prescriptions"`
	BillResponse    BillResponse           `json:"bill"`
}

type PrescriptionResponse struct {
	PresciptionId           int64   `json:"prescription_id"`
	MedicineId              int64   `json:"medicine_id"`
	MedicineName            string  `json:"medicine_name"`
	MedicineUsesInstruction string  `json:"medicine_uses_instruction"`
	Price                   float32 `json:"price"`
	Instruction             string  `json:"instruction"`
	Amount                  int     `json:"amount"`
}

type BillResponse struct {
	BillId     int64   `json:"bill_id"`
	TotalPrice float32 `json:"total_price"`
	Status     string  `json:"status"`
}
