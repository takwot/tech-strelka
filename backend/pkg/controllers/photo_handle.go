package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

func (h *Handle) CreatePhoto(c *gin.Context) {

	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is of invalid type")
		return
	}

	idInt, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is of invalid type")
		return 
	}

	fmt.Print(idInt)

	tags := c.Query("tags")

	fmt.Print(tags)

	file, err := c.FormFile("file")

	if err!= nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s/%s", idInt, file.Filename))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"tags": tags,
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
