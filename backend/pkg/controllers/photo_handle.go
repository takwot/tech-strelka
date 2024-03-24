package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handle) CreatePhoto(c *gin.Context) {
	id, err := h.getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}
	file, _ := c.FormFile("file")
	filename := file.Filename
	c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s/%s", id, filename))
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})
}
