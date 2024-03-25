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
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "validation error")
	}

	photo, err := h.service.Photo.GetPhoto(input.Id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"photo": photo,
	})
}

func (h *Handle) DeletePhoto(c *gin.Context) {
	var input models.Photo

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "validation error")
	}

	err := h.service.Photo.DeletePhoto(input.Id)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
