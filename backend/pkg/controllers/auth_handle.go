package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/takwot/tech-strelka.git/pkg/models"
)

func (h *Handle) signUp(c *gin.Context) {

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type SignInInput struct {
	Name     string `json:"name" binding: "required"`
	Password string `json:"password" binding: "required"`
}

func (h *Handle) uploadAvatar(c *gin.Context) {

	id, err := h.getUserId(c)

	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
	}

	file, _ := c.FormFile("file")

	c.SaveUploadedFile(file, fmt.Sprintf("./upload/%s/Avatar.png", id))

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
	})

}

func (h *Handle) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Auth.GenerateToken(input.Name, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error while creating token")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
