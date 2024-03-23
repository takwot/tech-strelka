package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handle) addPhoto(c *gin.Context) {

	form, err := c.MultipartForm()

	files := form.File["photo[]"]

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, file := range files {

		c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s/%s", 3, file.Filename))
	}

	c.JSON(http.StatusOK, gin.H{"status": true})
}
