package router

import (
	"github.com/AndersonMeloo/Go.git/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// Version in API - To define routes within this group, use v1.GET(), v1.POST(), etc.
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
