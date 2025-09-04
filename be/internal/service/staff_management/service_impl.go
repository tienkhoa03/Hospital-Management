package staffmanagement

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
	taskRepository "BE_Hospital_Management/internal/repository/task"
	userRepository "BE_Hospital_Management/internal/repository/user"
	userRoleRepository "BE_Hospital_Management/internal/repository/user_role"
	"BE_Hospital_Management/pkg/utils"
	"errors"

	"gorm.io/gorm"
)

type staffManagementService struct {
	userRepo        userRepository.UserRepository
	patientRepo     patientRepository.PatientRepository
	managerRepo     managerRepository.ManagerRepository
	staffRepo       staffRepository.StaffRepository
	doctorRepo      doctorRepository.DoctorRepository
	nurseRepo       nurseRepository.NurseRepository
	userRoleRepo    userRoleRepository.UserRoleRepository
	staffRoleRepo   staffRoleRepository.StaffRoleRepository
	taskRepo        taskRepository.TaskRepository
	appointmentRepo appointmentRepository.AppointmentRepository
}

func NewStaffManagementService(userRepo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, doctorRepo doctorRepository.DoctorRepository, managerRepo managerRepository.ManagerRepository, nurseRepo nurseRepository.NurseRepository, staffRepo staffRepository.StaffRepository, patientRepo patientRepository.PatientRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository, taskRepo taskRepository.TaskRepository, appointmentRepo appointmentRepository.AppointmentRepository) StaffManagementService {
	return &staffManagementService{
		userRepo:        userRepo,
		userRoleRepo:    userRoleRepo,
		doctorRepo:      doctorRepo,
		managerRepo:     managerRepo,
		nurseRepo:       nurseRepo,
		staffRepo:       staffRepo,
		patientRepo:     patientRepo,
		staffRoleRepo:   staffRoleRepo,
		taskRepo:        taskRepo,
		appointmentRepo: appointmentRepo,
	}
}
func (service *staffManagementService) CreateTask(managerId int64, staffUID int64, taskRequest *dto.TaskInfoRequest) (*dto.TaskInfoResponse, error) {
	var response *dto.TaskInfoResponse
	if taskRequest.BeginTime.After(taskRequest.FinishTime) {
		return nil, ErrInvalidTimeRange
	}
	if utils.CheckTimeInWorkingHours(taskRequest.BeginTime, taskRequest.FinishTime) == false {
		return nil, ErrOutOfWorkingHours
	}
	db := service.staffRepo.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, staffUID)
		if err != nil {
			return err
		}
		staff, err := service.staffRepo.GetStaffByUserId(staffUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
		}
		existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(staff.Id, taskRequest.BeginTime, taskRequest.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapTask {
			return ErrExistsOverlapTask
		}
		if staff.Role.RoleSlug == constant.RoleDoctor {
			doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, taskRequest.BeginTime, taskRequest.FinishTime)
			if err != nil {
				return err
			}
			if existsOverlapAppointment {
				return ErrExistsOverlapAppointment
			}
		}
		manager, err := service.managerRepo.GetManagerByUserId(managerId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		if staff.ManageBy != manager.Id {
			return ErrNotPermitted
		}
		task := &entity.Task{
			StaffId:     staff.Id,
			AssignerId:  manager.Id,
			Title:       taskRequest.Title,
			Description: &taskRequest.Description,
			BeginTime:   taskRequest.BeginTime,
			FinishTime:  taskRequest.FinishTime,
			Status:      taskRequest.Status,
		}
		newTask, err := service.taskRepo.CreateTask(tx, task)
		if err != nil {
			return err
		}
		response = utils.MapToTaskResponse(newTask, staff.UserId, manager.UserId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *staffManagementService) GetTasksByStaffUID(staffUID int64) ([]*dto.TaskInfoResponse, error) {
	staff, err := service.staffRepo.GetStaffByUserId(staffUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	tasks, err := service.taskRepo.GetTasksByStaffId(staff.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.TaskInfoResponse
	for _, task := range tasks {
		manager, err := service.managerRepo.GetManagerById(task.AssignerId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = append(response, utils.MapToTaskResponse(task, staff.UserId, manager.UserId))
	}
	return response, nil
}

func (service *staffManagementService) GetTasksByManagerUID(managerUID int64) ([]*dto.TaskInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	tasks, err := service.taskRepo.GetTasksByManagerId(manager.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.TaskInfoResponse
	for _, task := range tasks {
		staff, err := service.staffRepo.GetStaffById(task.StaffId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = append(response, utils.MapToTaskResponse(task, staff.UserId, manager.UserId))
	}
	return response, nil
}

func (service *staffManagementService) GetTasksByMangerUIDAndStaffUID(managerUID, staffUID int64) ([]*dto.TaskInfoResponse, error) {
	manager, err := service.managerRepo.GetManagerByUserId(managerUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	staff, err := service.staffRepo.GetStaffByUserId(staffUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.ManageBy != manager.Id {
		return nil, ErrNotPermitted
	}
	tasks, err := service.taskRepo.GetTasksByManagerIdAndStaffId(manager.Id, staff.Id)
	if err != nil {
		return nil, err
	}
	var response []*dto.TaskInfoResponse
	for _, task := range tasks {
		response = append(response, utils.MapToTaskResponse(task, staff.UserId, manager.UserId))
	}
	return response, nil
}

func (service *staffManagementService) GetTaskById(authUserId int64, authUserRole string, taskId int64) (*dto.TaskInfoResponse, error) {
	task, err := service.taskRepo.GetTaskById(taskId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	manager, err := service.managerRepo.GetManagerByUserId(authUserId)
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
	if authUserRole == constant.RoleManager {
		if task.AssignerId != manager.Id {
			return nil, ErrNotPermitted
		}
	} else {
		if task.StaffId != staff.Id {
			return nil, ErrNotPermitted
		}
	}
	response := utils.MapToTaskResponse(task, staff.UserId, manager.UserId)
	return response, nil
}

func (service *staffManagementService) DeleteTaskById(authUserId int64, taskId int64) error {
	task, err := service.taskRepo.GetTaskById(taskId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrTaskNotFound
		}
		return err
	}
	manager, err := service.managerRepo.GetManagerByUserId(authUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUserNotFound
		}
		return err
	}
	if task.AssignerId != manager.Id {
		return ErrNotPermitted
	}
	db := service.staffRepo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err = service.taskRepo.DeleteTaskById(tx, taskId)
		return err
	})
	return err
}

func (service *staffManagementService) UpdateTaskById(authUserId int64, taskId int64, updateRequest *dto.UpdateTaskInfoRequest) (*dto.TaskInfoResponse, error) {
	task, err := service.taskRepo.GetTaskById(taskId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, err
	}
	manager, err := service.managerRepo.GetManagerByUserId(authUserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if task.AssignerId != manager.Id {
		return nil, ErrNotPermitted
	}
	var staffUID int64
	if updateRequest.StaffUID != nil {
		staff, err := service.staffRepo.GetStaffByUserId(*updateRequest.StaffUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		staffUID = staff.UserId
		task.StaffId = staff.Id
	} else {
		staff, err := service.staffRepo.GetStaffById(task.StaffId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		staffUID = staff.UserId
	}
	if updateRequest.Title != nil {
		task.Title = *updateRequest.Title
	}
	if updateRequest.Description != nil {
		task.Description = updateRequest.Description
	}
	if updateRequest.BeginTime != nil {
		task.BeginTime = *updateRequest.BeginTime
	}
	if updateRequest.FinishTime != nil {
		task.FinishTime = *updateRequest.FinishTime
	}
	if updateRequest.Status != nil {
		task.Status = *updateRequest.Status
	}

	if task.BeginTime.After(task.FinishTime) {
		return nil, ErrInvalidTimeRange
	}
	if utils.CheckTimeInWorkingHours(task.BeginTime, task.FinishTime) == false {
		return nil, ErrOutOfWorkingHours
	}
	var response *dto.TaskInfoResponse
	db := service.staffRepo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err := utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, staffUID)
		if err != nil {
			return err
		}
		staff, err := service.staffRepo.GetStaffByUserId(staffUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
		}
		existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(staff.Id, task.BeginTime, task.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapTask {
			return ErrExistsOverlapTask
		}
		if staff.Role.RoleSlug == constant.RoleDoctor {
			doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, task.BeginTime, task.FinishTime)
			if err != nil {
				return err
			}
			if existsOverlapAppointment {
				return ErrExistsOverlapAppointment
			}
		}
		manager, err := service.managerRepo.GetManagerByUserId(authUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		if staff.ManageBy != manager.Id {
			return ErrNotPermitted
		}
		task := &entity.Task{
			StaffId:     staff.Id,
			AssignerId:  manager.Id,
			Title:       task.Title,
			Description: task.Description,
			BeginTime:   task.BeginTime,
			FinishTime:  task.FinishTime,
			Status:      task.Status,
		}
		updatedTask, err := service.taskRepo.UpdateTask(tx, task)
		if err != nil {
			return err
		}
		response = utils.MapToTaskResponse(updatedTask, staff.UserId, manager.UserId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}
