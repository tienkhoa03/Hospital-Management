package auth

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_auth_repository.go

type AuthRepository interface {
	GetDB() *gorm.DB
	CreateToken(tx *gorm.DB, token *entity.UserToken) error
	FindByRefreshToken(refreshToken string) (*entity.UserToken, error)
	SetRefreshTokenIsRevoked(tx *gorm.DB, refreshToken string) error
}
