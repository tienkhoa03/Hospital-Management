package service

import (
	"BE_Hospital_Management/internal/repository"
	"BE_Hospital_Management/internal/service/auth"
)

type Service struct {
	Auth auth.AuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: auth.NewAuthService(repos.Auth, repos.User, repos.UserRole, repos.Doctor, repos.Manager, repos.Nurse, repos.Staff, repos.Patient, repos.StaffRole),
	}
}
