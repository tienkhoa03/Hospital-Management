package api

import (
	"BE_Hospital_Management/api/handler"
	"BE_Hospital_Management/api/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, handlers *handler.Handlers, db *gorm.DB) {
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.TimeoutMiddleware(5 * time.Second))
	api := r.Group("/api")
	authApi := r.Group("/api")
	registerUserRoutes(api, handlers.User, db)
	registerStaffManagementRoutes(api, handlers.StaffManagement, db)
	registerAuthRoutes(authApi, handlers.Auth, db)
	registerAppointmentRoutes(api, handlers.Appointment, db)
}
