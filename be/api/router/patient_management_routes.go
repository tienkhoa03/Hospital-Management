package api

import (
	patientManagement "BE_Hospital_Management/api/handler/patient_management"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerPatientManagementRoutes(api *gin.RouterGroup, h *patientManagement.PatientManagementHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.POST("/treatment", middleware.RequireAnyRole([]string{constant.RoleDoctor}), h.CreateTreatmentPlan)
	api.GET("/treatment", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RolePatient}), h.GetAllTreatmentPlan)
	api.GET("/treatment/:id", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RolePatient}), h.GetTreatmentPlanById)
}
