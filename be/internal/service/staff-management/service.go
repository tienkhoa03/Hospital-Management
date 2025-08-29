package staffmanagement

import (
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	"errors"
)

var (
	ErrUserNotFound             = errors.New("user not found")
	ErrTaskNotFound             = errors.New("task not found")
	ErrNotPermitted             = errors.New("you are not permitted to perform this action")
	ErrInvalidTimeRange         = errors.New("invalid time range")
	ErrExistsOverlapTask        = errors.New("there is an overlapping task for the staff in the given time range")
	ErrExistsOverlapAppointment = errors.New("there is an overlapping appointment for the doctor in the given time range")
	ErrOutOfWorkingHours        = errors.New("time is out of working hours")
)

//go:generate mockgen -source=service.go -destination=../mock/mock_staff_management_service.go

type StaffManagementService interface {
	CreateTask(managerId int64, staffUID int64, task *dto.TaskInfoRequest) (*entity.Task, error)
	GetTasksByStaffUID(staffUID int64) ([]*entity.Task, error)
	GetTasksByManagerUID(managerUID int64) ([]*entity.Task, error)
	GetTasksByMangerUIDAndStaffUID(managerUID, staffUID int64) ([]*entity.Task, error)
	GetTaskById(authUserId int64, authUserRole string, taskId int64) (*entity.Task, error)
	DeleteTaskById(authUserId int64, taskId int64) error
}
