package entity

import "time"

type UserToken struct {
	Id           int64 `gorm:"primaryKey;autoIncrement"`
	UserId       int64
	RefreshToken string `json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ExpiresAt    time.Time
	IsRevoked    bool

	User *User `gorm:"foreignKey:UserId;references:Id"`
}
