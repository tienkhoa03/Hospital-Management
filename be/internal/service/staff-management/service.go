package staffmanagement

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	"errors"
)

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrNotPermitted             = errors.New("you are not permitted to perform this action")
	ErrInvalidTimeRange         = errors.New("invalid time range")
	ErrExistsOverlapTask        = errors.New("there is an overlapping task for the staff in the given time range")
	ErrExistsOverlapAppointment = errors.New("there is an overlapping appointment for the doctor in the given time range")
)

//go:generate mockgen -source=service.go -destination=../mock/mock_staff_management_service.go

type StaffManagementService interface {
	AssignTask(authUserId int64, staffUID int64, task *dto.TaskInfoRequest) (*entity.Task, error)
}
