package appointment

import (
	"BE_Hospital_Management/constant"
	"BE_Hospital_Management/internal/domain/dto"
	service "BE_Hospital_Management/internal/service/patient_management"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type PatientManagementHandler struct {
	service service.PatientManagementService
}

func NewAppointmentHandler(service service.PatientManagementService) *PatientManagementHandler {
	return &PatientManagementHandler{service: service}
}

// PatientManagement godoc
// @Summary      Create a treatment plan
// @Description  Create a treatment plan
// @Tags         PatientManagement
// @Accept 		json
// @Produce      json
// @Param		request	 	body		dto.TreatmentPlanRequest		true	"Treatment Plan"
// @param Authorization header string true "Authorization"
// @Router       /api/treatment [POST]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *PatientManagementHandler) CreateTreatmentPlan(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when creating treatment plan. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	var request dto.TreatmentPlanRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("Happened error when mapping request. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Invalid request format.")
	}
	newTreatmentPlan, err := h.service.CreateTreatmentPlan(*authUserId, request)
	if err != nil {
		log.Error("Happened error when creating treatment plan. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when creating treatment plan.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, newTreatmentPlan))
}

// PatientManagement godoc
// @Summary      Get all treatment plan of current user
// @Description  Get all treatment plan of current user
// @Tags         PatientManagement
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/treatment [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *PatientManagementHandler) GetAllTreatmentPlan(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting treatment plans. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting treatment plans. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	treatmentPlans, err := h.service.GetMedicalHistory(*authUserId, *authUserRole)
	if err != nil {
		log.Error("Happened error when getting treatment plans. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting treatment plans.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, treatmentPlans))
}

// PatientManagement godoc
// @Summary      Get treatment plan by Id
// @Description  Get treatment plan by Id
// @Tags         PatientManagement
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/treatment/{id} [GET]
// @Param        id   path      int  true  "Medical Record ID"
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *PatientManagementHandler) GetTreatmentPlanById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting treatment plan. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting treatment plan. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	medicalRecordIdStr := c.Param("id")
	medicalRecordId, err := strconv.ParseInt(medicalRecordIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	treatmentPlan, err := h.service.GetMedicalRecordById(*authUserId, *authUserRole, medicalRecordId)
	if err != nil {
		log.Error("Happened error when getting treatment plan. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting treatment plan.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, treatmentPlan))
}
