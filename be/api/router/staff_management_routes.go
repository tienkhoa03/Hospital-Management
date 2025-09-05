package api

import (
	staffManagement "BE_Hospital_Management/api/handler/staff_management"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerStaffManagementRoutes(api *gin.RouterGroup, h *staffManagement.StaffManagementHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.POST("/staff-management/staffs/:uid/tasks", middleware.RequireAnyRole([]string{constant.RoleManager}), h.AssignTaskToStaff)
	api.GET("/staff-management/me/tasks", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RoleCashingOfficer, constant.RoleNurse}), h.GetMyTasks)
	api.GET("/staff-management/me/assigned-tasks", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyAssignedTasks)
	api.GET("/staff-management/staffs/:uid/tasks", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyAssignedTasksToAStaff)
	api.GET("/staff-management/tasks/:id", middleware.RequireAnyRole([]string{constant.RoleManager, constant.RoleDoctor, constant.RoleCashingOfficer, constant.RoleNurse}), h.GetTaskById)
	api.DELETE("/staff-management/tasks/:id", middleware.RequireAnyRole([]string{constant.RoleManager}), h.DeleteTaskById)
	api.PATCH("/staff-management/tasks/:id", middleware.RequireAnyRole([]string{constant.RoleManager}), h.UpdateTaskById)
	api.GET("/staff-management/me/tasks/filter", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RoleCashingOfficer, constant.RoleNurse}), h.GetMyTasksWithFilter)
	api.GET("/staff-management/me/assigned-tasks/filter", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyAssignedTasksWithFilter)
	api.GET("/staff-management/staffs/:uid/tasks/filter", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyAssignedTasksToAStaffWithFilter)
}
