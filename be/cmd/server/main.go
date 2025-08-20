package main

import (
	"BE_Hospital_Management/api/handler"
	api "BE_Hospital_Management/api/router"
	"BE_Hospital_Management/config"
	"BE_Hospital_Management/internal/repository"
	"BE_Hospital_Management/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	docs "BE_Hospital_Management/cmd/server/docs"
)

// @title           Hospital Management API
// @version         1.0
// @description     Hospital Management API
// @BasePath
// @schemes         http https

func main() {
	config.LoadEnv()
	db := config.ConnectToDB()
	repos := repository.NewRepository(db)

	services := service.NewService(repos)
	handlers := handler.NewHandlers(services)

	r := gin.Default()
	api.SetupRoutes(r, handlers, db)
	docs.SwaggerInfo.Host = config.BASE_URL_BACKEND_FOR_SWAGGER
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(config.Port); err != nil {
		log.Fatal("failed to run server:", err)
	}
}
