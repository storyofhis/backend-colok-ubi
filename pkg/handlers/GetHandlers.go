package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/backend/colok-ubi/pkg/models"
)

func (h handler) GetHandlers(ctx *gin.Context) {
	var books []models.Books

	if result := h.DB.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &books)
}
