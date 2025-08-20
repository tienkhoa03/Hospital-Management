package api

import (
	"BE_Hospital_Management/api/handler/auth"
	"BE_Hospital_Management/api/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func registerAuthRoutes(api *gin.RouterGroup, h *auth.AuthHandler, db *gorm.DB) {
	api.POST("/auth/register", middleware.ParseAccessToken(), h.RegisterUser)
	api.POST("/auth/login", h.Login)
	api.POST("/auth/refresh", h.RefreshAccessToken)
	api.POST("auth/logout", h.Logout)
}
