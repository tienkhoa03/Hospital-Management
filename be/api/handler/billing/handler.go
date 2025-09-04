package billing

import (
	"BE_Hospital_Management/constant"
	service "BE_Hospital_Management/internal/service/billing"
	"BE_Hospital_Management/pkg"
	"BE_Hospital_Management/pkg/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BillingHandler struct {
	service service.BillingService
}

func NewBillingHandler(service service.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

// Billing godoc
// @Summary      Update bill status to paid
// @Description  Update the status of a bill to "paid" by a cashing officer
// @Tags         Billing
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/bills/{id}/status [PATCH]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *BillingHandler) UpdateBillStatusPaid(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when updating bill. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	billIdStr := c.Param("id")
	billId, err := strconv.ParseInt(billIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	updatedBill, err := h.service.UpdateBillStatusPaid(*authUserId, billId)
	if err != nil {
		log.Error("Happened error when updating bill. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when updating bill.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, updatedBill))
}

// Billing godoc
// @Summary      Get all bills of current user
// @Description  Get all bills of current user (patient or cashing officer)
// @Tags         Billing
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/bills [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *BillingHandler) GetAllBills(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting bills. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting getting bills. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	bills, err := h.service.GetAllBills(*authUserId, *authUserRole)
	if err != nil {
		log.Error("Happened error when getting bills. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting bills.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, bills))
}

// Billing godoc
// @Summary      Get bill by id
// @Description  Get bill by id
// @Tags         Billing
// @Accept 		json
// @Produce      json
// @param Authorization header string true "Authorization"
// @Router       /api/bills/{id} [GET]
// @Success      200   {object}  dto.ApiResponseSuccessStruct
// @securityDefinitions.apiKey token
// @in header
// @name Authorization
// @Security JWT
func (h *BillingHandler) GetBillById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	authUserId := utils.GetAuthUserId(c)
	if authUserId == nil {
		log.Error("Happened error when getting bills. Error: ", "Missing user ID in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user ID in context")
		return
	}
	authUserRole := utils.GetAuthUserRole(c)
	if authUserRole == nil {
		log.Error("Happened error when getting getting bills. Error: ", "Missing user role in context")
		pkg.PanicExeption(constant.Unauthorized, "Missing user role in context")
		return
	}
	billIdStr := c.Param("id")
	billId, err := strconv.ParseInt(billIdStr, 10, 64)
	if err != nil {
		log.Error("Happened error when converting Id to int64. Error: ", err)
		pkg.PanicExeption(constant.InvalidRequest, "Happened error when converting Id to int64")
	}
	bill, err := h.service.GetBillById(*authUserId, *authUserRole, billId)
	if err != nil {
		log.Error("Happened error when getting bills. Error: ", err)
		switch {
		case errors.Is(err, service.ErrNotPermitted):
			pkg.PanicExeption(constant.Unauthorized, err.Error())
		case errors.Is(err, service.ErrUserNotFound):
			pkg.PanicExeption(constant.DataNotFound, err.Error())
		default:
			pkg.PanicExeption(constant.UnknownError, "Happened error when getting bills.")
		}
	}
	c.JSON(http.StatusOK, pkg.BuildResponseSuccess(constant.Success, bill))
}
