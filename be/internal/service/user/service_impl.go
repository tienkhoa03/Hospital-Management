package user

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	appointmentRepository "BE_Hospital_Management/internal/repository/appointment"
	doctorRepository "BE_Hospital_Management/internal/repository/doctor"
	managerRepository "BE_Hospital_Management/internal/repository/manager"
	nurseRepository "BE_Hospital_Management/internal/repository/nurse"
	patientRepository "BE_Hospital_Management/internal/repository/patient"
	staffRepository "BE_Hospital_Management/internal/repository/staff"
	staffRoleRepository "BE_Hospital_Management/internal/repository/staff_role"
	userRepository "BE_Hospital_Management/internal/repository/user"
	userRoleRepository "BE_Hospital_Management/internal/repository/user_role"
	"BE_Hospital_Management/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type userService struct {
	repo            userRepository.UserRepository
	userRoleRepo    userRoleRepository.UserRoleRepository
	staffRoleRepo   staffRoleRepository.StaffRoleRepository
	patientRepo     patientRepository.PatientRepository
	staffRepo       staffRepository.StaffRepository
	managerRepo     managerRepository.ManagerRepository
	doctorRepo      doctorRepository.DoctorRepository
	nurseRepo       nurseRepository.NurseRepository
	appointmentRepo appointmentRepository.AppointmentRepository
}

func NewUserService(repo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository, patientRepo patientRepository.PatientRepository, staffRepo staffRepository.StaffRepository, mangerRepo managerRepository.ManagerRepository, doctorRepo doctorRepository.DoctorRepository, nurseRepo nurseRepository.NurseRepository, appointmentRepo appointmentRepository.AppointmentRepository) UserService {
	return &userService{
		repo:            repo,
		userRoleRepo:    userRoleRepo,
		staffRoleRepo:   staffRoleRepo,
		patientRepo:     patientRepo,
		staffRepo:       staffRepo,
		managerRepo:     mangerRepo,
		doctorRepo:      doctorRepo,
		nurseRepo:       nurseRepo,
		appointmentRepo: appointmentRepo,
	}
}

//	func (service *userService) GetAllUser() ([]*entity.User, error) {
//		users, err := service.repo.GetAllUser()
//		return users, err
//	}
func (service *userService) GetUserById(userId int64) (*dto.UserInfoResponse, error) {
	user, err := service.repo.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	var response *dto.UserInfoResponse
	if user.Role.RoleSlug == constant.RolePatient {
		patient, err := service.patientRepo.GetPatientByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = utils.MapPatientToUserInfoResponse(user, patient)
	} else if user.Role.RoleSlug == constant.RoleManager {
		manager, err := service.managerRepo.GetManagerByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = utils.MapManagerToUserInfoResponse(user, manager)
	} else if user.Role.RoleSlug == constant.RoleStaff {
		staff, err := service.staffRepo.GetStaffByUserId(userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		if staff.Role.RoleSlug == constant.RoleDoctor {
			doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrUserNotFound
				}
				return nil, err
			}
			response = utils.MapDoctorToUserInfoResponse(user, staff, doctor)
		} else if staff.Role.RoleSlug == constant.RoleNurse {
			nurse, err := service.nurseRepo.GetNurseByStaffId(staff.Id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrUserNotFound
				}
				return nil, err
			}
			response = utils.MapNurseToUserInfoResponse(user, staff, nurse)
		} else if staff.Role.RoleSlug == constant.RoleCashingOfficer {
			response = utils.MapCashingOfficerToUserInfoResponse(user, staff)
		}
	}
	return response, nil
}

func (service *userService) GetAllPatientsByDoctorUID(doctorUID int64) ([]*dto.UserInfoResponse, error) {
	user, err := service.repo.GetUserById(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if user.Role.RoleSlug != constant.RoleDoctor {
		return nil, ErrNotPermitted
	}
	var response []*dto.UserInfoResponse
	staff, err := service.staffRepo.GetStaffByUserId(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.Status == constant.StaffStatusInactive {
		return nil, ErrNotPermitted
	}
	doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	patientIds, err := service.appointmentRepo.GetPatientIdsByDoctorId(doctor.Id)
	if err != nil {
		return nil, err
	}
	patients, err := service.patientRepo.GetPatientsFromIdsWithUserInfo(patientIds)
	if err != nil {
		return nil, err
	}
	for _, patient := range patients {
		userInfoResponse := utils.MapPatientToUserInfoResponse(patient.User, patient)
		response = append(response, userInfoResponse)
	}
	return response, nil
}

func (service *userService) GetPatientByUserIdForDoctor(patientUID, doctorUID int64) (*dto.UserInfoResponse, error) {
	patient, err := service.patientRepo.GetPatientByUserIdWithUserInfo(patientUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	staff, err := service.staffRepo.GetStaffByUserId(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.Status == constant.StaffStatusInactive {
		return nil, ErrNotPermitted
	}
	doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	_, err = service.appointmentRepo.GetAppointmentByPatientIdAndDoctorId(patient.Id, doctor.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotPermitted
		}
		return nil, err
	}
	response := utils.MapPatientToUserInfoResponse(patient.User, patient)
	return response, nil
}

func (service *userService) GetAllStaffsByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	staffs, err := service.staffRepo.GetStaffsByManagerIdWithInformation(manager.Id)
	if err != nil {
		return nil, err
	}
	if manager.Status == constant.ManagerStatusInactive {
		return nil, ErrNotPermitted
	}
	var response []*dto.UserInfoResponse
	for _, staff := range staffs {
		var userInfoResponse *dto.UserInfoResponse
		if staff.Role.RoleSlug == constant.RoleDoctor {
			if staff.User == nil || staff.Doctor == nil {
				continue
			}
			userInfoResponse = utils.MapDoctorToUserInfoResponse(staff.User, staff, staff.Doctor)
		} else if staff.Role.RoleSlug == constant.RoleNurse {
			if staff.User == nil || staff.Nurse == nil {
				continue
			}
			userInfoResponse = utils.MapNurseToUserInfoResponse(staff.User, staff, staff.Nurse)
		} else if staff.Role.RoleSlug == constant.RoleCashingOfficer {
			if staff.User == nil {
				continue
			}
			userInfoResponse = utils.MapCashingOfficerToUserInfoResponse(staff.User, staff)
		}
		response = append(response, userInfoResponse)
	}
	return response, nil
}

func (service *userService) GetAllDoctorsByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if manager.Status == constant.ManagerStatusInactive {
		return nil, ErrNotPermitted
	}
	staffs, err := service.staffRepo.GetStaffsByManagerIdWithInformation(manager.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.UserInfoResponse
	for _, staff := range staffs {
		var userInfoResponse *dto.UserInfoResponse
		if staff.Role.RoleSlug == constant.RoleDoctor {
			if staff.User == nil || staff.Doctor == nil {
				continue
			}
			userInfoResponse = utils.MapDoctorToUserInfoResponse(staff.User, staff, staff.Doctor)
			response = append(response, userInfoResponse)
		}
	}
	return response, nil
}

func (service *userService) GetAllNursesByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if manager.Status == constant.ManagerStatusInactive {
		return nil, ErrNotPermitted
	}
	staffs, err := service.staffRepo.GetStaffsByManagerIdWithInformation(manager.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.UserInfoResponse
	for _, staff := range staffs {
		var userInfoResponse *dto.UserInfoResponse
		if staff.Role.RoleSlug == constant.RoleNurse {
			if staff.User == nil || staff.Nurse == nil {
				continue
			}
			userInfoResponse = utils.MapNurseToUserInfoResponse(staff.User, staff, staff.Nurse)
			response = append(response, userInfoResponse)
		}
	}
	return response, nil
}

func (service *userService) GetAllCashingOfficersByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if manager.Status == constant.ManagerStatusInactive {
		return nil, ErrNotPermitted
	}
	staffs, err := service.staffRepo.GetStaffsByManagerIdWithInformation(manager.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.UserInfoResponse
	for _, staff := range staffs {
		var userInfoResponse *dto.UserInfoResponse
		if staff.Role.RoleSlug == constant.RoleCashingOfficer {
			if staff.User == nil {
				continue
			}
			userInfoResponse = utils.MapCashingOfficerToUserInfoResponse(staff.User, staff)
			response = append(response, userInfoResponse)
		}
	}
	return response, nil
}

func (service *userService) GetStaffByUserIdForManager(staffUID, managerUID int64) (*dto.UserInfoResponse, error) {
	staff, err := service.staffRepo.GetStaffsByUserIdWithInformation(staffUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.ManageBy != manager.Id {
		return nil, ErrNotPermitted
	}
	var response *dto.UserInfoResponse
	if staff.Role.RoleSlug == constant.RoleDoctor {
		if staff.User == nil || staff.Doctor == nil {
			return nil, ErrUserNotFound
		}
		response = utils.MapDoctorToUserInfoResponse(staff.User, staff, staff.Doctor)
	} else if staff.Role.RoleSlug == constant.RoleNurse {
		if staff.User == nil || staff.Nurse == nil {
			return nil, ErrUserNotFound
		}
		response = utils.MapNurseToUserInfoResponse(staff.User, staff, staff.Nurse)
	} else if staff.Role.RoleSlug == constant.RoleCashingOfficer {
		if staff.User == nil {
			return nil, ErrUserNotFound
		}
		response = utils.MapCashingOfficerToUserInfoResponse(staff.User, staff)
	}
	return response, nil
}

func (service *userService) DeleteManagerByUID(managerUID int64) error {
	_, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err := service.repo.DeleteUserById(tx, managerUID)
		if err != nil {
			return err
		}
		err = service.managerRepo.DeleteManagerByUserId(tx, managerUID)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}
	return err
}

func (service *userService) DeleteStaffByUID(staffUID, managerUID int64) error {
	staff, err := service.staffRepo.GetStaffByUserId(staffUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	if staff.ManageBy != manager.Id {
		return ErrNotPermitted
	}
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err := service.repo.DeleteUserById(tx, staffUID)
		if err != nil {
			return err
		}
		err = service.staffRepo.DeleteStaffByUserId(tx, staffUID)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrUserNotFound
	}
	return err
}

func (service *userService) UpdateUserProfile(userId int64, request *dto.UpdateUserRequest) (*entity.User, error) {
	user := entity.User{
		Id:          userId,
		Name:        request.Name,
		DateOfBirth: request.DateOfBirth,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		Gender:      request.Gender,
	}
	var updatedUser *entity.User
	db := service.repo.GetDB()
	var err error
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedUser, err = service.repo.UpdateUser(tx, &user)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (service *userService) UpdateManagerProfile(managerUID int64, request *dto.UpdateManagerRequest) (*entity.Manager, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	manager.Department = request.Department
	manager.Status = request.Status
	var updatedManager *entity.Manager
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedManager, err = service.managerRepo.UpdateManager(tx, manager)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedManager, nil
}

func (service *userService) UpdateDoctorProfile(authUserId int64, doctorUID int64, request *dto.UpdateDoctorRequest) (*entity.Staff, error) {
	staff, err := service.staffRepo.GetStaffsByUserIdWithInformation(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.ManageBy != authUserId {
		return nil, ErrNotPermitted
	}
	updateStaff := &entity.Staff{
		Id:         staff.Id,
		Department: request.Department,
		Status:     request.Status,
	}
	updateDoctor := &entity.Doctor{
		Id:                   staff.Doctor.Id,
		Specialization:       request.Specialization,
		MedicalLicenseNumber: request.MedicalLicenseNumber,
	}
	var updatedStaff *entity.Staff
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedStaff, err = service.staffRepo.UpdateStaff(tx, updateStaff)
		if err != nil {
			return err
		}
		updatedStaff.Doctor, err = service.doctorRepo.UpdateDoctor(tx, updateDoctor)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedStaff, nil
}

func (service *userService) UpdateNurseProfile(authUserId int64, nurseUID int64, request *dto.UpdateNurseRequest) (*entity.Staff, error) {
	staff, err := service.staffRepo.GetStaffsByUserIdWithInformation(nurseUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.ManageBy != authUserId {
		return nil, ErrNotPermitted
	}
	updateStaff := &entity.Staff{
		Id:         staff.Id,
		Department: request.Department,
		Status:     request.Status,
	}
	updateNurse := &entity.Nurse{
		Id:                   staff.Doctor.Id,
		NursingLicenseNumber: request.NursingLicenseNumber,
	}
	var updatedStaff *entity.Staff
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedStaff, err = service.staffRepo.UpdateStaff(tx, updateStaff)
		if err != nil {
			return err
		}
		updatedStaff.Nurse, err = service.nurseRepo.UpdateNurse(tx, updateNurse)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedStaff, nil
}

func (service *userService) UpdateCashingOfficerProfile(authUserId int64, cashingOfficerUID int64, request *dto.UpdateCashingOfficerRequest) (*entity.Staff, error) {
	staff, err := service.staffRepo.GetStaffsByUserIdWithInformation(cashingOfficerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.ManageBy != authUserId {
		return nil, ErrNotPermitted
	}
	updateStaff := &entity.Staff{
		Id:         staff.Id,
		Department: request.Department,
		Status:     request.Status,
	}
	var updatedStaff *entity.Staff
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedStaff, err = service.staffRepo.UpdateStaff(tx, updateStaff)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedStaff, nil
}

func (service *userService) UpdatePatientProfile(authUserId int64, patientUID int64, request *dto.UpdatePatientRequest) (*entity.Patient, error) {
	patient, err := service.patientRepo.GetPatientByUserId(patientUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	staff, err := service.staffRepo.GetStaffByUserId(authUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.Doctor == nil {
		return nil, ErrNotPermitted
	}
	_, err = service.appointmentRepo.GetAppointmentByPatientIdAndDoctorId(patient.Id, staff.Doctor.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotPermitted
		}
		return nil, err
	}
	updatePatient := &entity.Patient{
		Id:                patient.Id,
		InsuranceNumber:   &request.InsuranceNumber,
		BloodType:         &request.BloodType,
		Allergies:         &request.Allergies,
		ChronicConditions: &request.ChronicConditions,
		Status:            request.Status,
	}
	var updatedPatient *entity.Patient
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		updatedPatient, err = service.patientRepo.UpdatePatient(tx, updatePatient)
		return err
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return updatedPatient, nil
}
