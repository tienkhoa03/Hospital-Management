package dto

type TreatmentPlanRequest struct {
	PatientUID      int64                 `json:"patient_uid"`
	AppointmentId   int64                 `json:"appointment_id"`
	Symptom         string                `json:"symptom"`
	MedicalServices string                `json:"medical_services"`
	Diagnosis       string                `json:"diagnosis"`
	Treatment       string                `json:"treatment"`
	Prescriptions   []PrescriptionRequest `json:"prescriptions"`
}

type PrescriptionRequest struct {
	MedicineId  int64  `json:"medicine_id"`
	Instruction string `json:"instruction"`
	Amount      int    `json:"amount"`
}
