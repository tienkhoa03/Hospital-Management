package api

import (
	"BE_Hospital_Management/api/handler/billing"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerBillingRoutes(api *gin.RouterGroup, h *billing.BillingHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.PATCH("/bills/:id/status", middleware.RequireAnyRole([]string{constant.RoleCashingOfficer}), h.UpdateBillStatusPaid)
	api.GET("/bills", middleware.RequireAnyRole([]string{constant.RoleCashingOfficer, constant.RolePatient}), h.GetAllBills)
	api.GET("/bills/:id", middleware.RequireAnyRole([]string{constant.RoleCashingOfficer, constant.RoleNurse}), h.GetBillById)
}
