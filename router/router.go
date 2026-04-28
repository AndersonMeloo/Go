package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
