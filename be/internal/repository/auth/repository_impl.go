package auth

import (
	"BE_Hospital_Management/internal/domain/entity"

	"gorm.io/gorm"
)

type PostgreSQLAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &PostgreSQLAuthRepository{db: db}
}
func (r *PostgreSQLAuthRepository) GetDB() *gorm.DB {
	return r.db
}

func (r *PostgreSQLAuthRepository) CreateToken(tx *gorm.DB, token *entity.UserToken) error {
	result := tx.Create(token)
	return result.Error
}

func (r *PostgreSQLAuthRepository) FindByRefreshToken(refreshToken string) (*entity.UserToken, error) {
	var userToken = entity.UserToken{}
	err := r.db.Model(&entity.UserToken{}).Where("refresh_token = ?", refreshToken).First(&userToken).Error
	if err != nil {
		return nil, err
	}
	return &userToken, nil
}

func (r *PostgreSQLAuthRepository) SetRefreshTokenIsRevoked(tx *gorm.DB, refreshToken string) error {
	err := tx.Model(&entity.UserToken{}).Where("refresh_token = ?", refreshToken).Update("is_revoked", true).Error
	return err
}
