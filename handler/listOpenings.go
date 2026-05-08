package handler

import (
	"net/http"

	"github.com/AndersonMeloo/Go.git/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(ctx *gin.Context) {

	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing opening")
	}

	sendSucces(ctx, "list-openings", openings)
}
