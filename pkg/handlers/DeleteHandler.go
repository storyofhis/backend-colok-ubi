package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/backend/colok-ubi/pkg/models"
)

func (h handler) DeletedHandler(c *gin.Context) {
	id := c.Param("id")

	var book models.Books

	if result := h.DB.First(&book, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.DB.Delete(&book)

	c.Status(http.StatusOK)
}
