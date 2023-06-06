package app

import "medx/auth/controllers"

func mapRoutes() {
	router.POST("/api/v1/register", controllers.AuthController.Register)
	router.POST("/api/v1/login", controllers.AuthController.Login)
}
