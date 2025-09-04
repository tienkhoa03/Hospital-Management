package appointment

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
	"time"
)

type appointmentService struct {
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

func NewAppointmentService(userRepo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, doctorRepo doctorRepository.DoctorRepository, managerRepo managerRepository.ManagerRepository, nurseRepo nurseRepository.NurseRepository, staffRepo staffRepository.StaffRepository, patientRepo patientRepository.PatientRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository, taskRepo taskRepository.TaskRepository, appointmentRepo appointmentRepository.AppointmentRepository) AppointmentService {
	return &appointmentService{
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
func (service *appointmentService) CreateAppointment(authUserId int64, authUserRole string, appointmentRequest *dto.AppointmentInfoRequest) (*dto.AppointmentInfoResponse, error) {
	if appointmentRequest.BeginTime.After(appointmentRequest.FinishTime) {
		return nil, ErrInvalidTimeRange
	}
	if utils.CheckTimeInWorkingHours(appointmentRequest.BeginTime, appointmentRequest.FinishTime) == false {
		return nil, ErrOutOfWorkingHours
	}
	if authUserRole == constant.RolePatient {
		if appointmentRequest.DoctorUID == nil {
			return nil, ErrMissingDoctorId
		}
		if appointmentRequest.PatientUID != nil {
			if *appointmentRequest.PatientUID != authUserId {
				return nil, ErrNotPermitted
			}
		}
		appointmentRequest.PatientUID = &authUserId
	} else if authUserRole == constant.RoleDoctor {
		if appointmentRequest.PatientUID == nil {
			return nil, ErrMissingPatientId
		}
		if appointmentRequest.DoctorUID != nil {
			if *appointmentRequest.DoctorUID != authUserId {
				return nil, ErrNotPermitted
			}
		}
		appointmentRequest.DoctorUID = &authUserId
	}
	var response *dto.AppointmentInfoResponse
	db := service.staffRepo.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		err := utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, *appointmentRequest.DoctorUID)
		if err != nil {
			return err
		}
		staff, err := service.staffRepo.GetStaffByUserId(*appointmentRequest.DoctorUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
		}
		existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(staff.Id, appointmentRequest.BeginTime, appointmentRequest.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapTask {
			return ErrExistsOverlapTask
		}
		doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, appointmentRequest.BeginTime, appointmentRequest.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapAppointment {
			return ErrExistsOverlapAppointment
		}
		patient, err := service.patientRepo.GetPatientByUserId(*appointmentRequest.PatientUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		appointment := &entity.Appointment{
			PatientId:  patient.Id,
			DoctorId:   doctor.Id,
			BeginTime:  appointmentRequest.BeginTime,
			FinishTime: appointmentRequest.FinishTime,
			Status:     appointmentRequest.Status,
		}
		newAppointment, err := service.appointmentRepo.CreateAppointment(tx, appointment)
		if err != nil {
			return err
		}
		response = utils.MapToAppointmentResponse(newAppointment, patient.UserId, staff.UserId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *appointmentService) UpdateAppointment(patientUID, appointmentId int64, request *dto.UpdateAppointmentRequest) (*dto.AppointmentInfoResponse, error) {
	var response *dto.AppointmentInfoResponse
	db := service.staffRepo.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		appointment, err := service.appointmentRepo.GetAppointmentById(appointmentId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrAppointmentNotFound
			}
			return err
		}

		patient, err := service.patientRepo.GetPatientByUserId(patientUID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		if appointment.PatientId != patient.Id {
			return ErrNotPermitted
		}

		var doctor *entity.Doctor
		var staff *entity.Staff
		if request.DoctorUID != nil {
			err = utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, *request.DoctorUID)
			if err != nil {
				return err
			}
			staff, err = service.staffRepo.GetStaffByUserId(*request.DoctorUID)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			doctor, err = service.doctorRepo.GetDoctorByStaffId(staff.Id)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			appointment.DoctorId = doctor.Id
		} else {
			doctor, err = service.doctorRepo.GetDoctorById(appointment.DoctorId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			staff, err = service.staffRepo.GetStaffById(doctor.StaffId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			err = utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, staff.UserId)
			if err != nil {
				return err
			}
		}
		if request.BeginTime != nil {
			appointment.BeginTime = *request.BeginTime
		}
		if request.FinishTime != nil {
			appointment.FinishTime = *request.FinishTime
		}
		if request.Status != nil {
			appointment.Status = *request.Status
		}

		if appointment.BeginTime.After(appointment.FinishTime) {
			return ErrInvalidTimeRange
		}
		if utils.CheckTimeInWorkingHours(appointment.BeginTime, appointment.FinishTime) == false {
			return ErrOutOfWorkingHours
		}
		existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(doctor.StaffId, appointment.BeginTime, appointment.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapTask {
			return ErrExistsOverlapTask
		}
		existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, appointment.BeginTime, appointment.FinishTime)
		if err != nil {
			return err
		}
		if existsOverlapAppointment {
			return ErrExistsOverlapAppointment
		}
		updatedAppointment, err := service.appointmentRepo.UpdateAppointment(tx, appointment)
		if err != nil {
			return err
		}
		response = utils.MapToAppointmentResponse(updatedAppointment, patient.UserId, staff.UserId)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *appointmentService) DeleteAppointment(requestorUID int64, requestorRole string, appointmentId int64) error {
	db := service.appointmentRepo.GetDB()
	err := db.Transaction(func(tx *gorm.DB) error {
		appointment, err := service.appointmentRepo.GetAppointmentById(appointmentId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrAppointmentNotFound
			}
			return err
		}
		appointmentDoctor, err := service.doctorRepo.GetDoctorById(appointment.DoctorId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		appointmentStaff, err := service.staffRepo.GetStaffById(appointmentDoctor.StaffId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return ErrUserNotFound
			}
			return err
		}
		err = utils.AdvisoryLock(tx, utils.NamespaceStaffSchedule, appointmentStaff.UserId)
		if requestorRole == constant.RolePatient {
			appointmentPatient, err := service.patientRepo.GetPatientByUserId(appointment.PatientId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrUserNotFound
				}
				return err
			}
			if appointmentPatient.UserId != requestorUID {
				return ErrNotPermitted
			}

		} else if requestorRole == constant.RoleDoctor {
			if appointmentStaff.UserId != requestorUID {
				return ErrNotPermitted
			}
		} else {
			return ErrNotPermitted
		}
		err = service.appointmentRepo.DeleteAppointmentById(tx, appointmentId)
		return err
	})
	return err
}

func (service *appointmentService) GetAvailableSlots(doctorUID int64, date time.Time) ([]*dto.AppointmentSlot, error) {
	var availableSlots []*dto.AppointmentSlot
	staff, err := service.staffRepo.GetStaffByUserId(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	if staff.Status != constant.StaffStatusWorking {
		return nil, ErrDoctorNotWorking
	}
	doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	for startTime := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location()); startTime.Before(time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 0, date.Location())); startTime = startTime.Add(constant.SlotInterval) {
		endTime := startTime.Add(constant.SlotInterval)
		if utils.CheckTimeInWorkingHours(startTime, endTime) == false {
			continue
		}
		existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(staff.Id, startTime, endTime)
		if err != nil {
			return nil, err
		}
		if existsOverlapTask {
			continue
		}
		existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, startTime, endTime)
		if err != nil {
			return nil, err
		}
		if existsOverlapAppointment {
			continue
		}
		slot := &dto.AppointmentSlot{
			BeginTime:  startTime,
			FinishTime: endTime,
		}
		availableSlots = append(availableSlots, slot)
	}
	return availableSlots, nil
}

func (service *appointmentService) CheckAvailableSlot(doctorUID int64, beginTime, finishTime time.Time) (bool, error) {
	if beginTime.After(finishTime) {
		return false, ErrInvalidTimeRange
	}
	if utils.CheckTimeInWorkingHours(beginTime, finishTime) == false {
		return false, ErrOutOfWorkingHours
	}
	staff, err := service.staffRepo.GetStaffByUserId(doctorUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrUserNotFound
		}
		return false, err
	}
	doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, ErrUserNotFound
		}
		return false, err
	}
	existsOverlapTask, err := service.taskRepo.ExistsOverlapTaskOfStaff(staff.Id, beginTime, finishTime)
	if err != nil {
		return false, err
	}
	if existsOverlapTask {
		return false, nil
	}
	existsOverlapAppointment, err := service.appointmentRepo.ExistsOverlapAppointmentOfDoctor(doctor.Id, beginTime, finishTime)
	if err != nil {
		return false, err
	}
	if existsOverlapAppointment {
		return false, nil
	}
	return true, nil
}

func (service *appointmentService) GetAllAppointments(authUserId int64, authUserRole string) ([]*dto.AppointmentInfoResponse, error) {
	var response []*dto.AppointmentInfoResponse
	if authUserRole == constant.RolePatient {
		patient, err := service.patientRepo.GetPatientByUserId(authUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		appointments, err := service.appointmentRepo.GetAppointmentsByPatientId(patient.Id)
		if err != nil {
			return nil, err
		}
		for _, appointment := range appointments {
			doctor, err := service.doctorRepo.GetDoctorById(appointment.DoctorId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrUserNotFound
				}
				return nil, err
			}
			staff, err := service.staffRepo.GetStaffById(doctor.StaffId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrUserNotFound
				}
				return nil, err
			}
			response = append(response, utils.MapToAppointmentResponse(appointment, patient.UserId, staff.UserId))
		}
	} else if authUserRole == constant.RoleDoctor {
		staff, err := service.staffRepo.GetStaffByUserId(authUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		appointments, err := service.appointmentRepo.GetAppointmentsByDoctorId(doctor.Id)
		if err != nil {
			return nil, err
		}
		for _, appointment := range appointments {
			patient, err := service.patientRepo.GetPatientById(appointment.PatientId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, ErrUserNotFound
				}
				return nil, err
			}
			response = append(response, utils.MapToAppointmentResponse(appointment, patient.UserId, staff.UserId))
		}
	} else {
		return nil, ErrNotPermitted
	}
	return response, nil
}

func (service *appointmentService) GetAppointmentById(authUserId int64, authUserRole string, appointmentId int64) (*dto.AppointmentInfoResponse, error) {
	var response *dto.AppointmentInfoResponse
	appointment, err := service.appointmentRepo.GetAppointmentById(appointmentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAppointmentNotFound
		}
		return nil, err
	}
	if authUserRole == constant.RolePatient {
		patient, err := service.patientRepo.GetPatientByUserId(authUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		if appointment.PatientId != patient.Id {
			return nil, ErrNotPermitted
		}
		doctor, err := service.doctorRepo.GetDoctorById(appointment.DoctorId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		staff, err := service.staffRepo.GetStaffById(doctor.StaffId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = utils.MapToAppointmentResponse(appointment, patient.UserId, staff.UserId)
	} else if authUserRole == constant.RoleDoctor {
		staff, err := service.staffRepo.GetStaffByUserId(authUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		doctor, err := service.doctorRepo.GetDoctorByStaffId(staff.Id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		if appointment.DoctorId != doctor.Id {
			return nil, ErrNotPermitted
		}
		patient, err := service.patientRepo.GetPatientById(appointment.PatientId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, ErrUserNotFound
			}
			return nil, err
		}
		response = utils.MapToAppointmentResponse(appointment, patient.UserId, staff.UserId)
	} else {
		return nil, ErrNotPermitted
	}

	return response, nil
}
