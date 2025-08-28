package repository

import (
	"BE_Hospital_Management/internal/repository/appointment"
	"BE_Hospital_Management/internal/repository/auth"
	"BE_Hospital_Management/internal/repository/doctor"
	"BE_Hospital_Management/internal/repository/manager"
	"BE_Hospital_Management/internal/repository/nurse"
	"BE_Hospital_Management/internal/repository/patient"
	"BE_Hospital_Management/internal/repository/staff"
	"BE_Hospital_Management/internal/repository/staff_role"
	"BE_Hospital_Management/internal/repository/task"
	"BE_Hospital_Management/internal/repository/user"
	"BE_Hospital_Management/internal/repository/user_role"

	"gorm.io/gorm"
)

type Repository struct {
	Auth        auth.AuthRepository
	Doctor      doctor.DoctorRepository
	Manager     manager.ManagerRepository
	Nurse       nurse.NurseRepository
	Patient     patient.PatientRepository
	UserRole    userrole.UserRoleRepository
	StaffRole   staffrole.StaffRoleRepository
	Staff       staff.StaffRepository
	User        user.UserRepository
	Appointment appointment.AppointmentRepository
	Task        task.TaskRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth:        auth.NewAuthRepository(db),
		Doctor:      doctor.NewDoctorRepository(db),
		Manager:     manager.NewManagerRepository(db),
		Nurse:       nurse.NewNurseRepository(db),
		Patient:     patient.NewPatientRepository(db),
		UserRole:    userrole.NewUserRoleRepository(db),
		StaffRole:   staffrole.NewStaffRoleRepository(db),
		Staff:       staff.NewStaffRepository(db),
		User:        user.NewUserRepository(db),
		Appointment: appointment.NewAppointmentRepository(db),
		Task:        task.NewTaskRepository(db),
	}
}
