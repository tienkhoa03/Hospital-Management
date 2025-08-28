package api

import (
	staffManagement "BE_Hospital_Management/api/handler/staff-management"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerStaffManagementRoutes(api *gin.RouterGroup, h *staffManagement.StaffManagementHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.POST("/staff-management/staffs/:uid/task", middleware.RequireAnyRole([]string{constant.RoleManager}), h.AssignTaskToStaff)
}
