package utils

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
)

func MapPatientToUserInfoResponse(user *entity.User, patient *entity.Patient) *dto.UserInfoResponse {
	if user == nil || patient == nil {
		return nil
	}
	return &dto.UserInfoResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CitizenId:   user.CitizenId,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Gender:      user.Gender,
		RoleId:      user.RoleId,
		Role:        constant.RolePatient,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		PatientInfo: &dto.PatientInfoResponse{
			Id:                patient.Id,
			InsuranceNumber:   patient.InsuranceNumber,
			BloodType:         patient.BloodType,
			Allergies:         patient.Allergies,
			ChronicConditions: patient.ChronicConditions,
			Status:            patient.Status,
		},
	}
}

func MapManagerToUserInfoResponse(user *entity.User, manager *entity.Manager) *dto.UserInfoResponse {
	if user == nil || manager == nil {
		return nil
	}
	return &dto.UserInfoResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CitizenId:   user.CitizenId,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Gender:      user.Gender,
		RoleId:      user.RoleId,
		Role:        constant.RoleManager,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		ManagerInfo: &dto.ManagerInfoResponse{
			Id:         manager.Id,
			Department: manager.Department,
			Status:     manager.Status,
		},
	}
}

func MapDoctorToUserInfoResponse(user *entity.User, staff *entity.Staff, doctor *entity.Doctor) *dto.UserInfoResponse {
	if user == nil || staff == nil || doctor == nil {
		return nil
	}
	return &dto.UserInfoResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CitizenId:   user.CitizenId,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Gender:      user.Gender,
		RoleId:      user.RoleId,
		Role:        constant.RoleStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		StaffInfo: &dto.StaffInfoResponse{
			Id:         staff.Id,
			Department: staff.Department,
			Status:     staff.Status,
			RoleId:     staff.RoleId,
			Role:       constant.RoleDoctor,
			DoctorInfo: &dto.DoctorInfoResponse{
				Id:                   doctor.Id,
				Specialization:       doctor.Specialization,
				MedicalLicenseNumber: doctor.MedicalLicenseNumber,
			},
		},
	}
}

func MapNurseToUserInfoResponse(user *entity.User, staff *entity.Staff, nurse *entity.Nurse) *dto.UserInfoResponse {
	if user == nil || staff == nil || nurse == nil {
		return nil
	}
	return &dto.UserInfoResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CitizenId:   user.CitizenId,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Gender:      user.Gender,
		RoleId:      user.RoleId,
		Role:        constant.RoleStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		StaffInfo: &dto.StaffInfoResponse{
			Id:         staff.Id,
			Department: staff.Department,
			Status:     staff.Status,
			RoleId:     staff.RoleId,
			Role:       constant.RoleNurse,
			NurseInfo: &dto.NurseInfoResponse{
				Id:                   nurse.Id,
				NursingLicenseNumber: nurse.NursingLicenseNumber,
			},
		},
	}
}

func MapCashingOfficerToUserInfoResponse(user *entity.User, staff *entity.Staff) *dto.UserInfoResponse {
	if user == nil || staff == nil {
		return nil
	}
	return &dto.UserInfoResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CitizenId:   user.CitizenId,
		DateOfBirth: user.DateOfBirth,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Gender:      user.Gender,
		RoleId:      user.RoleId,
		Role:        constant.RoleStaff,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		StaffInfo: &dto.StaffInfoResponse{
			Id:         staff.Id,
			Department: staff.Department,
			Status:     staff.Status,
			RoleId:     staff.RoleId,
			Role:       constant.RoleCashingOfficer,
		},
	}
}
