package api

import (
	"BE_Hospital_Management/api/handler/appointment"
	"BE_Hospital_Management/api/middleware"
	"BE_Hospital_Management/constant"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerAppointmentRoutes(api *gin.RouterGroup, h *appointment.AppointmentHandler, db *gorm.DB) {
	api.Use(middleware.ValidateAccessToken())
	api.POST("/appointments", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RolePatient}), h.CreateAppointment)
	api.PATCH("/appointments/:id", middleware.RequireAnyRole([]string{constant.RolePatient}), h.UpdateAppointment)
	api.DELETE("/appointments/:id", middleware.RequireAnyRole([]string{constant.RoleDoctor, constant.RolePatient}), h.DeleteAppointment)
	api.GET("/appointments/availability", h.GetAvailableSlots)
	api.GET("appointments/availability/check", h.CheckAvailableSlot)
	api.GET("/appointments", h.GetAllAppointments)
	api.GET("/appointments/:id", h.GetAppointmentById)
}
