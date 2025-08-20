package user

import (
	"BE_Hospital_Management/internal/domain/entity"
	"gorm.io/gorm"
)

//go:generate mockgen -source=repository.go -destination=../mock/mock_user_repository.go

type UserRepository interface {
	GetDB() *gorm.DB
	CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error)
	GetAllUser() ([]*entity.User, error)
	GetUserById(userId int64) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	GetUsersFromIds(userIds []int64) ([]*entity.User, error)
	GetUsersFromEmails(emails []string) ([]*entity.User, error)
	UpdateUser(tx *gorm.DB, user *entity.User) (*entity.User, error)
	DeleteUserById(tx *gorm.DB, userId int64) error
}
