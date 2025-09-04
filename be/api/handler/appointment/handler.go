package appointment

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/appointment"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AppointmentHandler struct {
	service service.AppointmentService
}

func NewAppointmentHandler(service service.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

// Appointment godoc
// @Summary      Create an appointment
// @Description  Create an appointment
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.AppointmentInfoRequest		true	"Appointment information"
// @param Authorization header string true "Authorization"
// @Router       /api/appointments [POST]
// @Success      201   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when creating appointment. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when creating appointment. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	var request dto.AppointmentInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newAppointment, err := h.service.CreateAppointment(*authUserId, *authUserRole, &request)
	if err != nil {
		log.Error("Happened error when creating appointment. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrInvalidTimeRange):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrExistsOverlapTask):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrExistsOverlapAppointment):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrMissingDoctorId):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrMissingPatientId):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when creating appointment.")
		}
	}
	c.JSON(http.StatusCreated, pkg.BuildResponseSuccess(constant.Success, newAppointment))
}

// Appointment godoc
// @Summary      Update appointment time
// @Description  Update appointment time
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.UpdateAppointmentRequest		true	"Appointment time"
// @param Authorization header string true "Authorization"
// @Router       /api/appointments/{id} [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating appointment. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	appointmentIdStr := c.Param("id")
	appointmentId, err := strconv.ParseInt(appointmentIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.UpdateAppointmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newAppointment, err := h.service.UpdateAppointment(*authUserId, appointmentId, &request)
	if err != nil {
		log.Error("Happened error when updating appointment. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrInvalidTimeRange):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		case errors.Is(err, service.ErrExistsOverlapTask):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrExistsOverlapAppointment):
			pkg.PanicExeption(constant.Conflict, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating appointment.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, newAppointment))
}

// // Appointment godoc
// // @Summary      Delete appointment
// // @Description  Delete appointment
// // @Tags         Appointment
// // @Accept 		json
// // @Produce      json
// // @Param		id	 	path		int		true	"Appointment id"
// // @param Authorization header string true "Authorization"
// // @Router       /api/appointments/{id} [DELETE]
// // @Success      200   {object}  dto.ApiResponseSuccessStruct
// // @securityDefinitions.apiKey token
// // @in header
// // @name Authorization
// // @Security JWT
func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when deleting appointment. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when deleting appointment. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	appointmentIdStr := c.Param("id")
	appointmentId, err := strconv.ParseInt(appointmentIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	err = h.service.DeleteAppointment(*authUserId, *authUserRole, appointmentId)
	if err != nil {
		log.Error("Happened error when deleting appointment. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when deleting appointment.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccessNoData())
}

// Appointment godoc
// @Summary      Get available slots of doctor
// @Description  Get available slots of doctor
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @Param		doctorUID	 	query		int     true  "Doctor ID"
// @Param		date	 	query		string     true  "Date (format: 2006-01-02)"
// @param Authorization header string true "Authorization"
// @Router       /api/appointments/availability [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) GetAvailableSlots(c *gin.Context) {
	defer pkg.PanicHandler(c)
	doctorUIDStr := c.Query("doctorUID")
	doctorUID, err := strconv.ParseInt(doctorUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Error("Happened error when converting date string to date format. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting date string to date format")
	}
	availableSlots, err := h.service.GetAvailableSlots(doctorUID, date)
	if err != nil {
		log.Error("Happened error when getting available slots. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrDoctorNotWorking):
			pkg.PanicExeption(constant.StatusForbidden, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting available slots.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, availableSlots))
}

// Appointment godoc
// @Summary      Get available slots of doctor
// @Description  Get available slots of doctor
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @Param		doctorUID	 	query		int     true  "Doctor ID"
// @Param		beginTime	 	query		string     true  "Begin Time (format: 2006-01-02 15:04)"
// @Param		finishTime	 	query		string     true  "Finish Time (format: 2006-01-02 15:04)"
// @param Authorization header string true "Authorization"
// @Router       /api/appointments/availability/check [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) CheckAvailableSlot(c *gin.Context) {
	defer pkg.PanicHandler(c)
	doctorUIDStr := c.Query("doctorUID")
	doctorUID, err := strconv.ParseInt(doctorUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	beginTimeStr := c.Query("beginTime")
	beginTime, err := time.Parse("2006-01-02 15:04", beginTimeStr)
	if err != nil {
		log.Error("Happened error when converting time string to time format. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting time string to time format")
	}
	finishTimeStr := c.Query("finishTime")
	finishTime, err := time.Parse("2006-01-02 15:04", finishTimeStr)
	if err != nil {
		log.Error("Happened error when converting time string to time format. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting time string to time format")
	}
	isAvailable, err := h.service.CheckAvailableSlot(doctorUID, beginTime, finishTime)
	if err != nil {
		log.Error("Happened error when getting available slots. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrDoctorNotWorking):
			pkg.PanicExeption(constant.StatusForbidden, err.Error())
		case errors.Is(err, service.ErrOutOfWorkingHours):
			pkg.PanicExeption(constant.InvalidRequest, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting available slots.")
		}
	}
	response := &dto.IsAvailableSlotResponse{
		IsAvailable: isAvailable,
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, response))
}

// Appointment godoc
// @Summary      Get all appointments of current user
// @Description  Get all appointments of current user
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/appointments [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting appointments. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting appointments. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	appointments, err := h.service.GetAllAppointments(*authUserId, *authUserRole)
	if err != nil {
		log.Error("Happened error when getting appointments. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting appointments.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, appointments))
}

// Appointment godoc
// @Summary      Get appointment by id
// @Description  Get appointment by id
// @Tags         Appointment
// @Accept 		json
// @Produce      json
// @Param		id	 	path		int		true	"Appointment id"
// @param Authorization header string true "Authorization"
// @Router       /api/appointments/{id} [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *AppointmentHandler) GetAppointmentById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting appointments. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting appointments. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	appointmentIdStr := c.Param("id")
	appointmentId, err := strconv.ParseInt(appointmentIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	appointment, err := h.service.GetAppointmentById(*authUserId, *authUserRole, appointmentId)
	if err != nil {
		log.Error("Happened error when getting appointments. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		case errors.Is(err, service.ErrAppointmentNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting appointments.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, appointment))
}
