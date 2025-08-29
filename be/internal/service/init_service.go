package service

import (
	"BE_Hospital_Management/internal/repository"
	"BE_Hospital_Management/internal/service/appointment"
	"BE_Hospital_Management/internal/service/auth"
	"BE_Hospital_Management/internal/service/staff-management"
	"BE_Hospital_Management/internal/service/user"
)

type Service struct {
	Auth            auth.AuthService
	User            user.UserService
	StaffManagement staffmanagement.StaffManagementService
	Appointment     appointment.AppointmentService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth:            auth.NewAuthService(repos.Auth, repos.User, repos.UserRole, repos.Doctor, repos.Manager, repos.Nurse, repos.Staff, repos.Patient, repos.StaffRole),
		User:            user.NewUserService(repos.User, repos.UserRole, repos.StaffRole, repos.Patient, repos.Staff, repos.Manager, repos.Doctor, repos.Nurse, repos.Appointment),
		StaffManagement: staffmanagement.NewStaffManagementService(repos.User, repos.UserRole, repos.Doctor, repos.Manager, repos.Nurse, repos.Staff, repos.Patient, repos.StaffRole, repos.Task, repos.Appointment),
		Appointment:     appointment.NewAppointmentService(repos.User, repos.UserRole, repos.Doctor, repos.Manager, repos.Nurse, repos.Staff, repos.Patient, repos.StaffRole, repos.Task, repos.Appointment),
	}
}
