package main

import (
	"BE_Friends_Management/config"
)

// @title           Hospital Management API
// @version         1.0
// @description     Hospital Management API
// @BasePath
// @schemes         http https

func main() {
	config.LoadEnv()
	config.ConnectToDB()
	//repos := repository.NewRepository(db)
	//
	//services := service.NewService(repos)
	//handlers := handler.NewHandlers(services)
	//
	//r := gin.Default()
	//api.SetupRoutes(r, handlers, db)
	//docs.SwaggerInfo.Host = config.BASE_URL_BACKEND_FOR_SWAGGER
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//
	//if err := r.Run(config.Port); err != nil {
	//	log.Fatal("failed to run server:", err)
	//}
}
