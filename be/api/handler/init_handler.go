package handler

import (
	authHandler "BE_Hospital_Management/api/handler/auth"
	userHandler "BE_Hospital_Management/api/handler/user"
	"BE_Hospital_Management/internal/service"
)

type Handlers struct {
	Auth *authHandler.AuthHandler
	User *userHandler.UserHandler
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{
		Auth: authHandler.NewAuthHandler(services.Auth),
		User: userHandler.NewUserHandler(services.User),
	}
}
