package auth

import (
	"BE_Hospital_Management/internal/domain/dto"
	"errors"
)

var (
	ErrInvalidUserRole          = errors.New("invalid user role")
	ErrInvalidStaffRole         = errors.New("invalid staff role")
	ErrMissingPatientInfo       = errors.New("missing patient information")
	ErrMissingDoctorInfo        = errors.New("missing doctor information")
	ErrMissingManagerInfo       = errors.New("missing manager information")
	ErrMissingNurseInfo         = errors.New("missing nurse information")
	ErrMissingStaffInfo         = errors.New("missing staff information")
	ErrAlreadyRegistered        = errors.New("email has already been registed")
	ErrInvalidLoginRequest      = errors.New("email or password is incorrect")
	ErrInvalidRefreshToken      = errors.New("invalid refresh token")
	ErrRefreshTokenIsRevoked    = errors.New("refresh token is revoked")
	ErrRefreshTokenExpires      = errors.New("refresh token has expired")
	ErrInvalidSigningMethod     = errors.New("unexpected signing method")
	ErrNotPermitted             = errors.New("you are not permitted to perform this action")
	ErrUniqueConstraintViolated = errors.New("unique constraint violated")
)

//go:generate mockgen -source=interface.go -destination=../mock/mock_auth_service.go

type AuthService interface {
	RegisterUser(authUserId *int64, authUserRole *string, request dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(email, password string) (string, string, error)
	RefreshAccessToken(rawRefreshToken string) (string, string, error)
	Logout(rawRefreshToken string) error
}
