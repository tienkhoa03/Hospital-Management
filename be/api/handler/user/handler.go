package auth

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
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
		return
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
		return
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
// @Router       /api/users/me/patients/{uid} [get]
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
		return
	}
	patientUIDStr := c.Param("uid")
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
		return
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
		return
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
		return
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
// @Summary      Get all cashing officers for current manager
// @Description  Get all cashing officers for current manager
// @Tags         User
// @Accept 		 json
// @Produce      json
// @Router       /api/users/me/cashing-officers [get]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @param Authorization header string true "User Authorization"
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) GetAllMyCashingOfficers(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting cashing officers. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	userInfo, err := h.service.GetAllCashingOfficersByManagerUID(*authUserId)
	if err != nil {
		log.Error("Happened error when getting cashing officers. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting cashing officers.")
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
// @Router       /api/users/me/staffs/{uid} [get]
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
		return
	}
	staffUIDStr := c.Param("uid")
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

// // User godoc
// // @Summary      Delete manager by user ID
// // @Description  Delete manager by user ID
// // @Tags         User
// // @Accept 		 json
// // @Produce      json
// // @Router       /api/users/managers/{uid} [delete]
// // @Param id path int true "manager UID"
// // @Success      200   {object}  dto.ApiResponseSuccessStruct
// // @param Authorization header string true "User Authorization"
// // @securityDefinitions.apiKey token
// // @in header
// // @name Authorization
// // @Security JWT
func (h *UserHandler) DeleteManagerByUID(c *gin.Context) {
	defer pkg.PanicHandler(c)
	managerUIDStr := c.Param("uid")
	managerUID, err := strconv.ParseInt(managerUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	err = h.service.DeleteManagerByUID(managerUID)
	if err != nil {
		log.Error("Happened error when deleting manager. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when deleting manager")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessNoData())
}

// // User godoc
// // @Summary      Delete staff by user ID for manager
// // @Description  Delete staff by user ID for manager
// // @Tags         User
// // @Accept 		 json
// // @Produce      json
// // @Router       /api/users/me/staffs/{uid} [delete]
// // @Param id path int true "staff UID"
// // @Success      200   {object}  dto.ApiResponseSuccessStruct
// // @param Authorization header string true "User Authorization"
// // @securityDefinitions.apiKey token
// // @in header
// // @name Authorization
// // @Security JWT
func (h *UserHandler) DeleteStaffByUID(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when deleting staff. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	staffUIDStr := c.Param("uid")
	staffUID, err := strconv.ParseInt(staffUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	err = h.service.DeleteStaffByUID(staffUID, *authUserId)
	if err != nil {
		log.Error("Happened error when deleting staff. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when deleting staff")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessNoData())
}

// User godoc
// @Summary      Update current user profile
// @Description  Update current user profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateUserRequest			true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/users/me [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdateMyProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating user profile. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	var request dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedUser, err := h.service.UpdateUserProfile(*authUserId, &request)
	if err != nil {
		log.Error("Happened error when updating user profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating user profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedUser))
}

// User godoc
// @Summary      Update manager profile
// @Description  Update manager profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateManagerRequest		true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/managers/{uid} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdateManagerProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	managerUIDStr := c.Param("uid")
	managerUID, err := strconv.ParseInt(managerUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateManagerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedManager, err := h.service.UpdateManagerProfile(managerUID, &request)
	if err != nil {
		log.Error("Happened error when updating manager profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating manager profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedManager))
}

// User godoc
// @Summary      Update doctor profile
// @Description  Update doctor profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateDoctorRequest		true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/doctors/{uid} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdateDoctorProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating doctor profile. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	doctorUIDStr := c.Param("uid")
	doctorUID, err := strconv.ParseInt(doctorUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateDoctorRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedDoctor, err := h.service.UpdateDoctorProfile(*authUserId, doctorUID, &request)
	if err != nil {
		log.Error("Happened error when updating doctor profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating doctor profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedDoctor))
}

// User godoc
// @Summary      Update nurse profile
// @Description  Update nurse profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateNurseRequest		true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/nurses/{uid} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdateNurseProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating nurse profile. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	nurseUIDStr := c.Param("uid")
	nurseUID, err := strconv.ParseInt(nurseUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateNurseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedNurse, err := h.service.UpdateNurseProfile(*authUserId, nurseUID, &request)
	if err != nil {
		log.Error("Happened error when updating nurse profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating nurse profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedNurse))
}

// User godoc
// @Summary      Update cashing officer profile
// @Description  Update cashing officer profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateCashingOfficerRequest		true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/cashing-officer/{uid} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdateCashingOfficerProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating cashing officer profile. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	cashingOfficerUIDStr := c.Param("uid")
	cashingOfficerUID, err := strconv.ParseInt(cashingOfficerUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateCashingOfficerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedCashingOfficer, err := h.service.UpdateCashingOfficerProfile(*authUserId, cashingOfficerUID, &request)
	if err != nil {
		log.Error("Happened error when updating cashing officer profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating cashing officer profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedCashingOfficer))
}

// User godoc
// @Summary      Update patient profile
// @Description  Update patient profile
// @Tags         User
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdatePatientRequest		true	"New information"
// @param Authorization header string true "Authorization"
// @Router       /api/patients/{uid} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *UserHandler) UpdatePatientProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating patient profile. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	patientUIDStr := c.Param("uid")
	patientUID, err := strconv.ParseInt(patientUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdatePatientRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	updatedPatient, err := h.service.UpdatePatientProfile(*authUserId, patientUID, &request)
	if err != nil {
		log.Error("Happened error when updating patient profile. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating patient profile.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedPatient))
}
