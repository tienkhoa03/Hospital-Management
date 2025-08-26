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
	api.GET("users/me/patients/:uid", middleware.RequireAnyRole([]string{constant.RoleDoctor}), h.GetMyPatientByUID)
	api.GET("users/me/staffs", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyStaffs)
	api.GET("users/me/doctors", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyDoctors)
	api.GET("users/me/nurses", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyNurses)
	api.GET("users/me/cashing-officers", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetAllMyCashingOfficers)
	api.GET("users/me/staffs/:uid", middleware.RequireAnyRole([]string{constant.RoleManager}), h.GetMyStaffByUID)
	api.DELETE("users/managers/:uid", middleware.RequireAnyRole([]string{constant.RoleAdmin}), h.DeleteManagerByUID)
	api.DELETE("users/me/staffs/:uid", middleware.RequireAnyRole([]string{constant.RoleManager}), h.DeleteStaffByUID)
	api.PATCH("users/me", h.UpdateMyProfile)
	api.PATCH("managers/:uid", middleware.RequireAnyRole([]string{constant.RoleAdmin}), h.UpdateManagerProfile)
	api.PATCH("doctors/:uid", middleware.RequireAnyRole([]string{constant.RoleManager}), h.UpdateDoctorProfile)
	api.PATCH("nurses/:uid", middleware.RequireAnyRole([]string{constant.RoleManager}), h.UpdateNurseProfile)
	api.PATCH("cashing-officers/:uid", middleware.RequireAnyRole([]string{constant.RoleManager}), h.UpdateCashingOfficerProfile)
	api.PATCH("patients/:uid", middleware.RequireAnyRole([]string{constant.RoleDoctor}), h.UpdatePatientProfile)
}
