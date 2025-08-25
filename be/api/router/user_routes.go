package api

import (
	user "BE_Hospital_Management/api/handler/user"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerUserRoutes(api *gin.RouterGroup, h *user.UserHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.GET("/users/me", h.GetCurrentUser)
	api.GET("users/me/patients", middleware.RequireAnyRole([]string{constant.RoleDoctor}), h.GetAllMyPatients)
	api.GET("users/me/patients/:id", middleware.RequireAnyRole([]string{constant.RoleDoctor}), h.GetMyPatientByUID)
	api.GET("users/me/staffs", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyStaffs)
	api.GET("users/me/doctors", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyDoctors)
	api.GET("users/me/nurses", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyNurses)
	api.GET("users/me/staffs/:id", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyStaffByUID)
}
