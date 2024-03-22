package controllers

import (
	"net/http"
	"path/filepath"
	"strings"

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

		filepath := strings.Join([]string{file.Filename, filepath.Ext(file.Filename)}, "")
		err = c.SaveUploadedFile(file, filepath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Файлы успешно загружены"})
}
