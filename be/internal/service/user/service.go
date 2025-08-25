package user

import (
	"BE_Hospital_Management/internal/domain/dto"
	"errors"
)

//go:generate mockgen -source=interface.go -destination=../mock/mock_user_service.go

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
	//UpdateUser(userId int64, email string, password string) (*entity.User, error)
}
