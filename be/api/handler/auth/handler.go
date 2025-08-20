package auth

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/auth"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Auth godoc
// @Summary      Register new user
// @Description  Register new user
// @Tags         Auth
// @Accept 		 json
// @Produce      json
// @Param 		 request body dto.RegisterRequest true "Request Body"
// @Router       /api/auth/register [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string false "Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	authUserRole := utils.GetAuthUserRole(c)
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newUser, err := h.service.RegisterUser(authUserId, authUserRole, request)
	if err != nil {
		log.Error("Happened error when registing new user. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUniqueConstraintViolated):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrAlreadyRegistered):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrInvalidUserRole):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrInvalidStaffRole):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrMissingPatientInfo):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrMissingManagerInfo):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrMissingStaffInfo):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when registing new user")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, newUser))
}

// Auth godoc
// @Summary      Login
// @Description  Login
// @Tags         Auth
// @Accept 		 json
// @Produce      json
// @Param 		 request body dto.LoginRequest true "User's email and password"
// @Router       /api/auth/login [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
func (h *AuthHandler) Login(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	accessToken, refreshToken, err := h.service.Login(request.Email, request.Password)
	if err != nil {
		log.Error("Happened error when creating new friendship. Error: ", err)
		switch {
		case errors.Is(err, service.ErrInvalidLoginRequest):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when creating new friendship.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessWithTokens(accessToken, refreshToken))
}

// Auth godoc
// @Summary      Refresh Access Token
// @Description  Refresh Access Token
// @Tags         Auth
// @Accept 		 json
// @Produce      json
// @Param 		 request body dto.RefreshRequest true "User's refresh token"
// @Router       /api/auth/refresh [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
func (h *AuthHandler) RefreshAccessToken(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.RefreshRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newAccessToken, newRefreshToken, err := h.service.RefreshAccessToken(request.RefreshToken)
	if err != nil {
		log.Error("Happened error when creating new friendship. Error: ", err)
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrRefreshTokenIsRevoked):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrRefreshTokenExpires):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrInvalidSigningMethod):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when refreshing access token.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessWithTokens(newAccessToken, newRefreshToken))
}

// Auth godoc
// @Summary      Logout
// @Description  Logout
// @Tags         Auth
// @Accept 		 json
// @Produce      json
// @Param 		 request body dto.LogoutRequest true "User's refresh token"
// @Router       /api/auth/logout [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
func (h *AuthHandler) Logout(c *gin.Context) {
	defer pkg.PanicHandler(c)
	var request dto.LogoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	err := h.service.Logout(request.RefreshToken)
	if err != nil {
		log.Error("Happened error when logging out. Error: ", err)
		switch {
		case errors.Is(err, service.ErrInvalidRefreshToken):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrRefreshTokenIsRevoked):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrRefreshTokenExpires):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrInvalidSigningMethod):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when logging out.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessNoData())
}
