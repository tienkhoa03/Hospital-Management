package user

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &PostgreSQLUserRepository{db: db}
}

func (r *PostgreSQLUserRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLUserRepository) GetAllUser() ([]*entity.User, error) {
	var users = []*entity.User{}
	result := r.db.Model(&entity.User{}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *PostgreSQLUserRepository) GetUserById(userId int64) (*entity.User, error) {
	var user = entity.User{}
	result := r.db.Model(&entity.User{}).Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgreSQLUserRepository) GetUsersFromIds(userIds []int64) ([]*entity.User, error) {
	var users []*entity.User
	result := r.db.Model(&entity.User{}).Where("id IN ?", userIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *PostgreSQLUserRepository) GetUsersFromEmails(emails []string) ([]*entity.User, error) {
	var users []*entity.User
	result := r.db.Model(&entity.User{}).Where("email IN ?", emails).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *PostgreSQLUserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user = entity.User{}
	result := r.db.Model(&entity.User{}).Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *PostgreSQLUserRepository) CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	result := tx.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, result.Error
}

func (r *PostgreSQLUserRepository) DeleteUserById(tx *gorm.DB, userId int64) error {
	result := tx.Model(&entity.User{}).Where("id = ?", userId).Delete(entity.User{})
	return result.Error
}

func (r *PostgreSQLUserRepository) UpdateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	result := tx.Model(&entity.User{}).Where("id = ?", user.Id).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}
	var updatedUser = entity.User{}
	result = tx.First(&updatedUser, user.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updatedUser, nil
}
