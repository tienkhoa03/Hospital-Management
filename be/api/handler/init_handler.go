package handler

import (
	authHandler "BE_Hospital_Management/api/handler/auth"
	"BE_Hospital_Management/internal/service"
)

type Handlers struct {
	Auth *authHandler.AuthHandler
}

func NewHandlers(services *service.Service) *Handlers {
	return &Handlers{
		Auth: authHandler.NewAuthHandler(services.Auth),
	}
}
