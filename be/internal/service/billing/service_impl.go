package billing

import (
	"BE_Hospital_Management/constant"
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
	"BE_Hospital_Management/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type billingService struct {
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

func NewBillingService(userRepo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, doctorRepo doctorRepository.DoctorRepository, managerRepo managerRepository.ManagerRepository, nurseRepo nurseRepository.NurseRepository, staffRepo staffRepository.StaffRepository, patientRepo patientRepository.PatientRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository, taskRepo taskRepository.TaskRepository, appointmentRepo appointmentRepository.AppointmentRepository, medicalRecordRepo medicalRecordRepository.MedicalRecordRepository, prescriptionRepo prescriptionRepository.PrescriptionRepository, billRepo billRepository.BillRepository, billItemRepo billItemRepository.BillItemRepository) BillingService {
	return &billingService{
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
func (service *billingService) UpdateBillStatusPaid(cashingOfficerUID, billId int64) (*dto.BillResponse, error) {
	cashingOfficerStaff, err := service.staffRepo.GetStaffByUserId(cashingOfficerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotPermitted
		}
		return nil, err
	}
	if cashingOfficerStaff.Role.RoleSlug != constant.RoleCashingOfficer {
		return nil, ErrNotPermitted
	}
	db := service.billRepo.GetDB()
	var response *dto.BillResponse
	err = db.Transaction(func(tx *gorm.DB) error {
		bill := &entity.Bill{
			Id:               billId,
			Status:           constant.BillStatusPaid,
			CashingOfficerId: &cashingOfficerStaff.Id,
		}
		updatedBill, err := service.billRepo.UpdateBill(tx, bill)
		if err != nil {
			return err
		}
		billItems, err := service.billItemRepo.GetBillItemsByBillId(billId)
		if err != nil {
			return err
		}
		patient, err := service.patientRepo.GetPatientById(updatedBill.PatientId)
		if err != nil {
			return err
		}
		doctor, err := service.doctorRepo.GetDoctorById(updatedBill.DoctorId)
		if err != nil {
			return err
		}
		doctorStaff, err := service.staffRepo.GetStaffById(doctor.StaffId)
		if err != nil {
			return err
		}
		response = utils.MapToBillResponse(updatedBill, billItems, patient.UserId, doctorStaff.UserId, &cashingOfficerUID)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *billingService) GetAllBills(userId int64, userRole string) ([]*dto.BillResponse, error) {
	var responses []*dto.BillResponse
	if userRole == constant.RolePatient {
		patient, err := service.patientRepo.GetPatientByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		bills, err := service.billRepo.GetBillsByPatientId(patient.Id)
		if err != nil {
			return nil, err
		}
		for _, bill := range bills {
			doctor, err := service.doctorRepo.GetDoctorById(bill.DoctorId)
			if err != nil {
				return nil, err
			}
			doctorStaff, err := service.staffRepo.GetStaffById(doctor.StaffId)
			if err != nil {
				return nil, err
			}
			var cashingOfficerUID *int64
			if bill.CashingOfficerId != nil {
				cashingOfficer, err := service.staffRepo.GetStaffById(*bill.CashingOfficerId)
				if err != nil {
					return nil, err
				}
				cashingOfficerUID = &cashingOfficer.UserId
			}
			billItems, err := service.billItemRepo.GetBillItemsByBillId(bill.Id)
			if err != nil {
				return nil, err
			}
			response := utils.MapToBillResponse(bill, billItems, patient.UserId, doctorStaff.UserId, cashingOfficerUID)
			responses = append(responses, response)
		}
	} else if userRole == constant.RoleCashingOfficer {
		cashingOfficerStaff, err := service.staffRepo.GetStaffByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		bills, err := service.billRepo.GetBillsByCashingOfficerId(cashingOfficerStaff.Id)
		if err != nil {
			return nil, err
		}
		for _, bill := range bills {
			patient, err := service.patientRepo.GetPatientById(bill.PatientId)
			if err != nil {
				return nil, err
			}
			doctor, err := service.doctorRepo.GetDoctorById(bill.DoctorId)
			if err != nil {
				return nil, err
			}
			doctorStaff, err := service.staffRepo.GetStaffById(doctor.StaffId)
			if err != nil {
				return nil, err
			}
			billItems, err := service.billItemRepo.GetBillItemsByBillId(bill.Id)
			if err != nil {
				return nil, err
			}
			response := utils.MapToBillResponse(bill, billItems, patient.UserId, doctorStaff.UserId, &userId)
			responses = append(responses, response)
		}
	} else {
		return nil, ErrNotPermitted
	}
	return responses, nil
}

func (service *billingService) GetBillById(userId int64, userRole string, billId int64) (*dto.BillResponse, error) {
	var response *dto.BillResponse
	bill, err := service.billRepo.GetBillById(billId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrBillNotFound
		}
		return nil, err
	}
	patient, err := service.patientRepo.GetPatientById(bill.PatientId)
	if err != nil {
		return nil, err
	}
	doctor, err := service.doctorRepo.GetDoctorById(bill.DoctorId)
	if err != nil {
		return nil, err
	}
	doctorStaff, err := service.staffRepo.GetStaffById(doctor.StaffId)
	if err != nil {
		return nil, err
	}
	billItems, err := service.billItemRepo.GetBillItemsByBillId(bill.Id)
	if err != nil {
		return nil, err
	}
	var cashingOfficerUID *int64
	if bill.CashingOfficerId != nil {
		cashingOfficer, err := service.staffRepo.GetStaffById(*bill.CashingOfficerId)
		if err != nil {
			return nil, err
		}
		cashingOfficerUID = &cashingOfficer.UserId
	}
	if userRole == constant.RolePatient {
		if patient.UserId != userId {
			return nil, ErrNotPermitted
		}
		response = utils.MapToBillResponse(bill, billItems, patient.UserId, doctorStaff.UserId, cashingOfficerUID)
	} else if userRole == constant.RoleCashingOfficer {
		cashingOfficerStaff, err := service.staffRepo.GetStaffByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		if bill.CashingOfficerId == nil || *bill.CashingOfficerId != cashingOfficerStaff.Id {
			return nil, ErrNotPermitted
		}
		response = utils.MapToBillResponse(bill, billItems, patient.UserId, doctorStaff.UserId, &userId)
	} else {
		return nil, ErrNotPermitted
	}
	return response, nil
}
