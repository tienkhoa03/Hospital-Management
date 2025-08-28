package user

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	"errors"
)

//go:generate mockgen -source=service.go -destination=../mock/mock_user_service.go

var (
	ErrUserNotFound = errors.New("user not found")
	ErrNotPermitted = errors.New("you are not permitted to perform this action")
)

type UserService interface {
	//GetAllUser() ([]*entity.User, error)
	GetUserById(userId int64) (*dto.UserInfoResponse, error)
	GetAllPatientsByDoctorUID(doctorUID int64) ([]*dto.UserInfoResponse, error)
	GetPatientByUserIdForDoctor(patientUID, doctorUID int64) (*dto.UserInfoResponse, error)
	GetAllStaffsByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error)
	GetAllDoctorsByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error)
	GetAllNursesByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error)
	GetStaffByUserIdForManager(staffUID, managerUID int64) (*dto.UserInfoResponse, error)
	GetAllCashingOfficersByManagerUID(managerUID int64) ([]*dto.UserInfoResponse, error)
	DeleteManagerByUID(managerUID int64) error
	DeleteStaffByUID(staffUID, managerUID int64) error
	UpdateUserProfile(userId int64, request *dto.UpdateUserRequest) (*entity.User, error)
	UpdateManagerProfile(managerUID int64, request *dto.UpdateManagerRequest) (*entity.Manager, error)
	UpdateDoctorProfile(authUserId int64, doctorUID int64, request *dto.UpdateDoctorRequest) (*entity.Staff, error)
	UpdateNurseProfile(authUserId int64, nurseUID int64, request *dto.UpdateNurseRequest) (*entity.Staff, error)
	UpdateCashingOfficerProfile(authUserId int64, cashingOfficerUID int64, request *dto.UpdateCashingOfficerRequest) (*entity.Staff, error)
	UpdatePatientProfile(authUserId int64, patientUID int64, request *dto.UpdatePatientRequest) (*entity.Patient, error)
}
