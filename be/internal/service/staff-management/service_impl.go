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
func (service *staffManagementService) AssignTask(authUserId int64, staffUID int64, taskRequest *dto.TaskInfoRequest) (*entity.Task, error) {
	if taskRequest.BeginTime.After(taskRequest.FinishTime) {
		return nil, ErrInvalidTimeRange
	}
	var newTask *entity.Task
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
			return ErrExistsOverlapAppointment
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
			Title:       taskRequest.Title,
			Description: &taskRequest.Description,
			BeginTime:   taskRequest.BeginTime,
			FinishTime:  taskRequest.FinishTime,
			Status:      taskRequest.Status,
		}
		newTask, err = service.taskRepo.CreateTask(tx, task)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return newTask, nil
}
