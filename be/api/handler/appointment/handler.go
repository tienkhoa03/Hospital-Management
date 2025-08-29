package appointment

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/appointment"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"

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
// @Success      200   {object}  dto.ApiResponseSuccessStruct
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
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, newAppointment))
}
