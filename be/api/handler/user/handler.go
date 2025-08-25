package auth

import (
	"BE_Hospital_Management/constant"
	service "BE_Hospital_Management/internal/service/user"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// User godoc
// @Summary      Get current user information
// @Description  Get current user information
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting current user. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	userInfo, err := h.service.GetUserById(*authUserId)
	if err != nil {
		log.Error("Happened error when getting current user. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, "User not found")
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting current user")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get all patients for current doctor
// @Description  Get all patients for current doctor
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/patients [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetAllMyPatients(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting patients. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	userInfo, err := h.service.GetAllPatientsByDoctorUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting patients. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting patients")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get patients by user ID for current doctor
// @Description  Get patients by user ID for current doctor
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/patients/{id} [get]
// @Param id path int true "Patient UID"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetMyPatientByUID(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting patient's information. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	patientUIDStr := c.Param("id")
	patientUID, err := strconv.ParseInt(patientUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	userInfo, err := h.service.GetPatientByUserIdForDoctor(patientUID, *authUserId)
	if err != nil {
		log.Error("Happened error when getting patient's information. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting patient's information")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get all staffs for current manager
// @Description  Get all staffs for current manager
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/staffs [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetAllMyStaffs(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting staffs. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	userInfo, err := h.service.GetAllStaffsByManagerUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting staffs. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting staffs.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get all doctors for current manager
// @Description  Get all doctors for current manager
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/doctors [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetAllMyDoctors(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting doctors. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	userInfo, err := h.service.GetAllDoctorsByManagerUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting doctors. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting doctors.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get all nurses for current manager
// @Description  Get all nurses for current manager
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/nurses [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetAllMyNurses(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting nurses. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	userInfo, err := h.service.GetAllNursesByManagerUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting nurses. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting nurses.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}

// User godoc
// @Summary      Get staff by user ID for current manager
// @Description  Get staff by user ID for current manager
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/staffs/{id} [get]
// @Param id path int true "Staff UID"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetMyStaffByUID(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting staff's information. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	staffUIDStr := c.Param("id")
	staffUID, err := strconv.ParseInt(staffUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	userInfo, err := h.service.GetStaffByUserIdForManager(staffUID, *authUserId)
	if err != nil {
		log.Error("Happened error when getting staff's information. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting staff's information")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, userInfo))
}
