package utils

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
)

func MapToTreatmentPlanResponse(medicalRecord *entity.MedicalRecord, prescriptions []*entity.Prescription, bill *entity.Bill, patientUID, doctorUID int64) *dto.TreatmentPlanResponse {
	if medicalRecord == nil || bill == nil {
		return nil
	}
	var prescriptionResponses []dto.PrescriptionResponse
	for _, prescription := range prescriptions {
		if prescription != nil {
			prescriptionResponses = append(prescriptionResponses, dto.PrescriptionResponse{
				PresciptionId:           prescription.Id,
				MedicineId:              prescription.MedicineId,
				MedicineName:            prescription.Medicine.Name,
				MedicineUsesInstruction: prescription.Medicine.UsesInstruction,
				Price:                   prescription.Medicine.Price,
				Instruction:             prescription.Instruction,
				Amount:                  prescription.Amount,
			})
		}
	}
	return &dto.TreatmentPlanResponse{
		MedicalRecordId: medicalRecord.Id,
		PatientUID:      patientUID,
		DoctorUID:       doctorUID,
		AppointmentId:   medicalRecord.AppointmentId,
		Symptom:         medicalRecord.Symptoms,
		MedicalServices: medicalRecord.MedicalServices,
		Note:            medicalRecord.Note,
		BillResponse: dto.BillInTreatmentPlanResponse{
			BillId:     bill.Id,
			TotalPrice: bill.TotalPrice,
			Status:     bill.Status,
		},
		Prescriptions: prescriptionResponses,
	}
}
