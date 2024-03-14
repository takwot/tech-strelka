package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handle) signUp(c *gin.Context){
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": false,
	})
}