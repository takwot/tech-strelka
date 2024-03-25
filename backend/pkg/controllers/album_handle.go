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
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	id, err := h.service.Album.CreateAlbum(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while creating",
		})
		fmt.Println(err.Error())
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
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
	}

	album, err := h.service.Album.GetAllAlbum()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while getting album",
		})
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"albums": album,
	})
}

func (h *Handle) getAlbum(c *gin.Context) {
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	album, err := h.service.Album.GetAlbum(input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while getting album",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"album": album,
	})
}

func (h *Handle) deleteAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	err := h.service.Album.DeleteAlbum(input.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while deleting album",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}

func (h *Handle) updateAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	err := h.service.Album.UpdateAlbum(input.Id, input.Photos)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while updating album",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}

func (h *Handle) renameAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "validation error",
		})
		return
	}

	err := h.service.Album.RenameAlbum(input.Id, input.Name)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "error while renaming album",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
