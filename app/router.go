package app

import "medx/auth/controllers"

func mapRoutes() {

	router.Group("/api/v1/")
	router.POST("/api/v1/register", controllers.AuthController.Register)
	router.POST("/api/v1/login", controllers.AuthController.Login)
}
