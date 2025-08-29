package handler

import (
	appointmentHandler "BE_Hospital_Management/api/handler/appointment"
	authHandler "BE_Hospital_Management/api/handler/auth"
	staffManagementHandler "BE_Hospital_Management/api/handler/staff-management"
	userHandler "BE_Hospital_Management/api/handler/user"
	"BE_Hospital_Management/internal/service"
)

type Handlers struct {
	Auth            *authHandler.AuthHandler
	User            *userHandler.UserHandler
	StaffManagement *staffManagementHandler.StaffManagementHandler
	Appointment     *appointmentHandler.AppointmentHandler
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{
		Auth:            authHandler.NewAuthHandler(services.Auth),
		User:            userHandler.NewUserHandler(services.User),
		StaffManagement: staffManagementHandler.NewStaffManagementHandler(services.StaffManagement),
		Appointment:     appointmentHandler.NewAppointmentHandler(services.Appointment),
	}
}
