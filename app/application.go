package app

import "github.com/gin-gonic/gin"

var router = gin.Default()

func StartApplication() {
	mapRoutes()
	router.Run(":8000")
	
}
