package patientmanagement

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	appointmentRepository "BE_Hospital_Management/internal/repository/appointment"
	billRepository "BE_Hospital_Management/internal/repository/bill"
	billItemRepository "BE_Hospital_Management/internal/repository/bill_item"
	doctorRepository "BE_Hospital_Management/internal/repository/doctor"
	managerRepository "BE_Hospital_Management/internal/repository/manager"
	medicalRecordRepository "BE_Hospital_Management/internal/repository/medical_record"
	nurseRepository "BE_Hospital_Management/internal/repository/nurse"
	patientRepository "BE_Hospital_Management/internal/repository/patient"
	prescriptionRepository "BE_Hospital_Management/internal/repository/prescription"
	staffRepository "BE_Hospital_Management/internal/repository/staff"
	staffRoleRepository "BE_Hospital_Management/internal/repository/staff_role"
	taskRepository "BE_Hospital_Management/internal/repository/task"
	userRepository "BE_Hospital_Management/internal/repository/user"
	userRoleRepository "BE_Hospital_Management/internal/repository/user_role"
	"errors"
	"gorm.io/gorm"
)

type patientManagementService struct {
	userRepo          userRepository.UserRepository
	patientRepo       patientRepository.PatientRepository
	managerRepo       managerRepository.ManagerRepository
	staffRepo         staffRepository.StaffRepository
	doctorRepo        doctorRepository.DoctorRepository
	nurseRepo         nurseRepository.NurseRepository
	userRoleRepo      userRoleRepository.UserRoleRepository
	staffRoleRepo     staffRoleRepository.StaffRoleRepository
	taskRepo          taskRepository.TaskRepository
	appointmentRepo   appointmentRepository.AppointmentRepository
	medicalRecordRepo medicalRecordRepository.MedicalRecordRepository
	prescriptionRepo  prescriptionRepository.PrescriptionRepository
	billRepo          billRepository.BillRepository
	billItemRepo      billItemRepository.BillItemRepository
}

func NewPatientManagementService(userRepo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, doctorRepo doctorRepository.DoctorRepository, managerRepo managerRepository.ManagerRepository, nurseRepo nurseRepository.NurseRepository, staffRepo staffRepository.StaffRepository, patientRepo patientRepository.PatientRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository, taskRepo taskRepository.TaskRepository, appointmentRepo appointmentRepository.AppointmentRepository, medicalRecordRepo medicalRecordRepository.MedicalRecordRepository, prescriptionRepo prescriptionRepository.PrescriptionRepository, billRepo billRepository.BillRepository, billItemRepo billItemRepository.BillItemRepository) PatientManagementService {
	return &patientManagementService{
		userRepo:          userRepo,
		userRoleRepo:      userRoleRepo,
		doctorRepo:        doctorRepo,
		managerRepo:       managerRepo,
		nurseRepo:         nurseRepo,
		staffRepo:         staffRepo,
		patientRepo:       patientRepo,
		staffRoleRepo:     staffRoleRepo,
		taskRepo:          taskRepo,
		appointmentRepo:   appointmentRepo,
		medicalRecordRepo: medicalRecordRepo,
		prescriptionRepo:  prescriptionRepo,
		billRepo:          billRepo,
		billItemRepo:      billItemRepo,
	}
}
func (service *patientManagementService) CreateTreatmentPlan(doctorUID int64, treatmentPlan dto.TreatmentPlanRequest) (*dto.TreatmentPlanResponse, error) {
	db := service.medicalRecordRepo.GetDB()
	var response *dto.TreatmentPlanResponse
	err := db.Transaction(func(tx *gorm.DB) error {
		staff, err := service.staffRepo.GetStaffByUserId(doctorUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		patient, err := service.patientRepo.GetPatientByUserId(treatmentPlan.PatientUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		medicalRecord := &entity.MedicalRecord{
			PatientId:       patient.Id,
			DoctorId:        doctor.Id,
			AppointmentId:   treatmentPlan.AppointmentId,
			Symptoms:        treatmentPlan.Symptom,
			MedicalServices: treatmentPlan.MedicalServices,
			Diagnosis:       treatmentPlan.Diagnosis,
			Treatment:       treatmentPlan.Treatment,
			Note:            treatmentPlan.Note,
		}
		newMedicalRecord, err := service.medicalRecordRepo.CreateMedicalRecord(tx, medicalRecord)
		if err != nil {
			return err
		}
		response.MedicalRecordId = newMedicalRecord.Id
		response.PatientUID = treatmentPlan.PatientUID
		response.AppointmentId = treatmentPlan.AppointmentId
		response.Symptom = treatmentPlan.Symptom
		response.MedicalServices = treatmentPlan.MedicalServices
		response.Diagnosis = treatmentPlan.Diagnosis
		response.Treatment = treatmentPlan.Treatment
		response.Note = treatmentPlan.Note
		var totalPrice float32 = 0
		for _, prescriptionRequest := range treatmentPlan.Prescriptions {
			prescription := &entity.Prescription{
				MedicalRecordId: newMedicalRecord.Id,
				MedicineId:      prescriptionRequest.MedicineId,
				Instruction:     prescriptionRequest.Instruction,
				Amount:          prescriptionRequest.Amount,
			}
			newPrescription, err := service.prescriptionRepo.CreatePrescription(tx, prescription)
			if err != nil {
				return err
			}
			prescriptionResponse := dto.PrescriptionResponse{
				PresciptionId:           newPrescription.Id,
				MedicineId:              prescriptionRequest.MedicineId,
				MedicineName:            newPrescription.Medicine.Name,
				MedicineUsesInstruction: newPrescription.Medicine.UsesInstruction,
				Price:                   newPrescription.Medicine.Price,
				Instruction:             prescriptionRequest.Instruction,
				Amount:                  prescriptionRequest.Amount,
			}
			response.Prescriptions = append(response.Prescriptions, prescriptionResponse)
			totalPrice += newPrescription.Medicine.Price * float32(prescriptionRequest.Amount)
		}
		bill := &entity.Bill{
			PatientId:       patient.Id,
			DoctorId:        doctor.Id,
			MedicalRecordId: newMedicalRecord.Id,
			TotalPrice:      totalPrice,
		}
		newBill, err := service.billRepo.CreateBill(tx, bill)
		if err != nil {
			return err
		}
		billResponse := dto.BillResponse{
			BillId:     newBill.Id,
			TotalPrice: newBill.TotalPrice,
			Status:     newBill.Status,
		}
		response.BillResponse = billResponse
		for _, prescriptionRequest := range treatmentPlan.Prescriptions {
			billItem := &entity.BillItem{
				BillId:     newBill.Id,
				MedicineId: prescriptionRequest.MedicineId,
				Amount:     prescriptionRequest.Amount,
			}
			_, err := service.billItemRepo.CreateBillItem(tx, billItem)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
