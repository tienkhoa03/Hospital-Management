package auth

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	"BE_Hospital_Management/internal/domain/entity"
	authRepository "BE_Hospital_Management/internal/repository/auth"
	doctorRepository "BE_Hospital_Management/internal/repository/doctor"
	managerRepository "BE_Hospital_Management/internal/repository/manager"
	nurseRepository "BE_Hospital_Management/internal/repository/nurse"
	patientRepository "BE_Hospital_Management/internal/repository/patient"
	staffRepository "BE_Hospital_Management/internal/repository/staff"
	staffRoleRepository "BE_Hospital_Management/internal/repository/staff_role"
	userRepository "BE_Hospital_Management/internal/repository/user"
	userRoleRepository "BE_Hospital_Management/internal/repository/user_role"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type authService struct {
	repo          authRepository.AuthRepository
	userRepo      userRepository.UserRepository
	patientRepo   patientRepository.PatientRepository
	managerRepo   managerRepository.ManagerRepository
	staffRepo     staffRepository.StaffRepository
	doctorRepo    doctorRepository.DoctorRepository
	nurseRepo     nurseRepository.NurseRepository
	userRoleRepo  userRoleRepository.UserRoleRepository
	staffRoleRepo staffRoleRepository.StaffRoleRepository
}

func NewAuthService(repo authRepository.AuthRepository, userRepo userRepository.UserRepository, userRoleRepo userRoleRepository.UserRoleRepository, doctorRepo doctorRepository.DoctorRepository, managerRepo managerRepository.ManagerRepository, nurseRepo nurseRepository.NurseRepository, staffRepo staffRepository.StaffRepository, patientRepo patientRepository.PatientRepository, staffRoleRepo staffRoleRepository.StaffRoleRepository) AuthService {
	return &authService{
		repo:          repo,
		userRepo:      userRepo,
		userRoleRepo:  userRoleRepo,
		doctorRepo:    doctorRepo,
		managerRepo:   managerRepo,
		nurseRepo:     nurseRepo,
		staffRepo:     staffRepo,
		patientRepo:   patientRepo,
		staffRoleRepo: staffRoleRepo,
	}
}

func (service *authService) RegisterUser(authUserId *int64, authUserRole *string, request dto.UserInfoRequest) (*dto.UserInfoResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	_, err = service.userRepo.GetUserByEmail(request.Email)
	if err == nil {
		return nil, ErrAlreadyRegistered
	}
	userRole, err := service.userRoleRepo.GetUserRoleBySlug(request.Role)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidUserRole
		}
		return nil, err
	}
	dob, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		return nil, err
	}
	user := entity.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    string(hashedPassword),
		CitizenId:   request.CitizenId,
		DateOfBirth: dob,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		Gender:      request.Gender,
		RoleId:      userRole.Id,
	}
	var response *dto.UserInfoResponse
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		newUser, err := service.userRepo.CreateUser(tx, &user)
		if err != nil {
			if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
				return ErrUniqueConstraintViolated
			}
			return err
		}
		if userRole.RoleSlug == constant.RolePatient {
			if request.PatientInfo == nil {
				return ErrMissingPatientInfo
			}
			patient := entity.Patient{
				UserId:            newUser.Id,
				InsuranceNumber:   request.PatientInfo.InsuranceNumber,
				BloodType:         request.PatientInfo.BloodType,
				Allergies:         request.PatientInfo.Allergies,
				ChronicConditions: request.PatientInfo.ChronicConditions,
				Status:            request.PatientInfo.Status,
			}
			newPatient, err := service.patientRepo.CreatePatient(tx, &patient)
			if err != nil {
				if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
					return ErrUniqueConstraintViolated
				}
				return err
			}
			response = utils.MapPatientToUserInfoResponse(newUser, newPatient)
		} else if userRole.RoleSlug == constant.RoleStaff {
			if authUserId == nil || authUserRole == nil {
				return ErrNotPermitted
			}
			if *authUserRole != constant.RoleManager {
				return ErrNotPermitted
			}
			if request.StaffInfo == nil {
				return ErrMissingStaffInfo
			}
			staffRole, err := service.staffRoleRepo.GetStaffRoleBySlug(request.StaffInfo.Role)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrInvalidStaffRole
				}
				return err
			}
			manager, err := service.managerRepo.GetManagerByUserId(*authUserId)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return ErrNotPermitted
				}
				return err
			}
			staff := entity.Staff{
				UserId:     newUser.Id,
				ManageBy:   manager.Id,
				Department: request.StaffInfo.Department,
				RoleId:     staffRole.Id,
				Status:     request.StaffInfo.Status,
			}
			newStaff, err := service.staffRepo.CreateStaff(tx, &staff)
			if err != nil {
				if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
					return ErrUniqueConstraintViolated
				}
				return err
			}
			if staffRole.RoleSlug == constant.RoleDoctor && request.StaffInfo.DoctorInfo == nil {
				return ErrMissingStaffInfo
			}
			if staffRole.RoleSlug == constant.RoleNurse && request.StaffInfo.NurseInfo == nil {
				return ErrMissingStaffInfo
			}
			if request.StaffInfo.DoctorInfo != nil {
				doctor := entity.Doctor{
					StaffId:              newStaff.Id,
					Specialization:       request.StaffInfo.DoctorInfo.Specialization,
					MedicalLicenseNumber: request.StaffInfo.DoctorInfo.MedicalLicenseNumber,
				}
				newDoctor, err := service.doctorRepo.CreateDoctor(tx, &doctor)
				if err != nil {
					if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
						return ErrUniqueConstraintViolated
					}
					return err
				}
				response = utils.MapDoctorToUserInfoResponse(newUser, newStaff, newDoctor)
			} else if request.StaffInfo.NurseInfo != nil {
				nurse := entity.Nurse{
					StaffId:              newStaff.Id,
					NursingLicenseNumber: request.StaffInfo.NurseInfo.NursingLicenseNumber,
				}
				newNurse, err := service.nurseRepo.CreateNurse(tx, &nurse)
				if err != nil {
					if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
						return ErrUniqueConstraintViolated
					}
					return err
				}
				response = utils.MapNurseToUserInfoResponse(newUser, newStaff, newNurse)
			}
		} else if userRole.RoleSlug == constant.RoleManager {
			if authUserId == nil || authUserRole == nil {
				return ErrNotPermitted
			}
			if *authUserRole != constant.RoleAdmin {
				return ErrNotPermitted
			}
			if request.ManagerInfo == nil {
				return ErrMissingManagerInfo
			}
			manager := entity.Manager{
				UserId:     newUser.Id,
				Department: request.ManagerInfo.Department,
				Status:     request.ManagerInfo.Status,
			}
			newManager, err := service.managerRepo.CreateManager(tx, &manager)
			if err != nil {
				if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
					return ErrUniqueConstraintViolated
				}
				return err
			}
			response = utils.MapManagerToUserInfoResponse(newUser, newManager)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (service *authService) Login(email, password string) (string, string, error) {
	user, err := service.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", ErrInvalidLoginRequest
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", ErrInvalidLoginRequest
	}
	accessTokenExpiredTime := time.Now().Add(utils.AccessTokenExpiredTime)
	userRole := user.Role
	authUserRole := userRole.RoleSlug
	if userRole.RoleSlug == constant.RoleStaff {
		staff, err := service.staffRepo.GetStaffByUserId(user.Id)
		if err != nil {
			return "", "", err
		}
		staffRole, err := service.staffRoleRepo.GetStaffRoleById(staff.RoleId)
		if err != nil {
			return "", "", err
		}
		authUserRole = staffRole.RoleSlug
	}
	accessToken, err := utils.GenerateAccessToken(user.Id, authUserRole, accessTokenExpiredTime)
	if err != nil {
		return "", "", err
	}
	refreshTokenExpiredTime := time.Now().Add(utils.RefreshTokenExpiredTime)
	refreshToken, err := utils.GenerateRefreshToken(user.Id, authUserRole, refreshTokenExpiredTime)
	if err != nil {
		return "", "", err
	}
	tokenRecord := &entity.UserToken{
		UserId:       user.Id,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshTokenExpiredTime,
		IsRevoked:    false,
	}
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err := service.repo.CreateToken(tx, tokenRecord)
		return err
	})
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service *authService) RefreshAccessToken(rawRefreshToken string) (string, string, error) {
	userToken, err := service.repo.FindByRefreshToken(rawRefreshToken)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", ErrInvalidRefreshToken
	}
	if err != nil {
		return "", "", err
	}
	if userToken.IsRevoked {
		return "", "", ErrRefreshTokenIsRevoked
	}
	claims, err := utils.ParseRefreshToken(rawRefreshToken)
	if errors.Is(err, utils.ErrInvalidRefreshToken) {
		return "", "", ErrInvalidRefreshToken
	}
	if errors.Is(err, utils.ErrInvalidSigningMethod) {
		return "", "", ErrInvalidSigningMethod
	}
	if err != nil {
		return "", "", err
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", "", ErrRefreshTokenExpires
	}
	accessToken, err := utils.GenerateAccessToken(userToken.UserId, claims.Role, time.Now().Add(utils.AccessTokenExpiredTime))
	if err != nil {
		return "", "", err
	}
	refreshTokenExpiredTime := time.Now().Add(utils.RefreshTokenExpiredTime)
	refreshToken, err := utils.GenerateRefreshToken(userToken.Id, claims.Role, refreshTokenExpiredTime)
	if err != nil {
		return "", "", err
	}
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err = service.repo.SetRefreshTokenIsRevoked(tx, rawRefreshToken)
		if err != nil {
			return err
		}
		tokenRecord := &entity.UserToken{
			UserId:       userToken.UserId,
			RefreshToken: refreshToken,
			ExpiresAt:    refreshTokenExpiredTime,
			IsRevoked:    false,
		}
		err = service.repo.CreateToken(tx, tokenRecord)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (service *authService) Logout(rawRefreshToken string) error {
	userToken, err := service.repo.FindByRefreshToken(rawRefreshToken)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInvalidRefreshToken
	}
	if err != nil {
		return err
	}
	if userToken.IsRevoked {
		return ErrRefreshTokenIsRevoked
	}
	claims, err := utils.ParseRefreshToken(rawRefreshToken)
	if errors.Is(err, utils.ErrInvalidRefreshToken) {
		return ErrInvalidRefreshToken
	}
	if errors.Is(err, utils.ErrInvalidSigningMethod) {
		return ErrInvalidSigningMethod
	}
	if err != nil {
		return err
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return ErrRefreshTokenExpires
	}
	db := service.repo.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		err = service.repo.SetRefreshTokenIsRevoked(tx, rawRefreshToken)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
