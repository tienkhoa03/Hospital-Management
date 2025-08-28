package staffmanagement

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/staff-management"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type StaffManagementHandler struct {
	service service.StaffManagementService
}

func NewStaffManagementHandler(service service.StaffManagementService) *StaffManagementHandler {
	return &StaffManagementHandler{service: service}
}

// StaffManagement godoc
// @Summary      Assign task to staff
// @Description  Assign task to staff
// @Tags         StaffManagement
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.TaskInfoRequest		true	"Task information"
// @param Authorization header string true "Authorization"
// @Router       /api/staff-management/staffs/{uid}/tasks [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *StaffManagementHandler) AssignTaskToStaff(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when assigning task to staff. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
	}
	staffUIDStr := c.Param("uid")
	staffUID, err := strconv.ParseInt(staffUIDStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	var request dto.TaskInfoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newTask, err := h.service.AssignTask(*authUserId, staffUID, &request)
	if err != nil {
		log.Error("Happened error when assigning task to staff. Error: ", err)
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
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when assigning task to staff.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, newTask))
}
