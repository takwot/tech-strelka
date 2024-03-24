package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

func (h *Handle) createAlbum(c *gin.Context) {

	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "validation error")
		return
	}

	id, err := h.service.Album.CreateAlbum(input)

	if err != nil {
		fmt.Print(err.Error())
		newErrorResponse(c, http.StatusBadRequest, "error while creating")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     id,
		"status": true,
	})
}

func (h *Handle) getAllAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "validation error")
	}

	album, err := h.service.Album.GetAllAlbum()

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"albums": album,
	})
}
