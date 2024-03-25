package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

func (h *Handle) CreatePhoto(c *gin.Context) {
	var input models.Photo
	id, err := h.service.Photo.CreatePhoto(input)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"id":     id,
	})
}

func (h *Handle) GetPhoto(c *gin.Context) {
	var input models.Photo

	photo, _ := h.service.Photo.GetPhoto(input.Id)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"photo": photo,
	})
}

func (h *Handle) DeletePhoto(c *gin.Context) {
	var input models.Photo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	err := h.service.Photo.DeletePhoto(input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while deleting photo",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
