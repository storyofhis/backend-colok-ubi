package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/backend/colok-ubi/pkg/models"
)

func (h handler) GetHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Books
	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	ctx.JSON(http.StatusOK, &book)
}
