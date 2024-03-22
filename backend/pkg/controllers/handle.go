package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/takwot/tech-strelka.git/pkg/service"
)

type Handle struct {
	service *service.Service
}

func NewHandle(service *service.Service) *Handle {
	return &Handle{service: service}
}

func (h *Handle) InitRoutes() *gin.Engine {

	router := gin.New()

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.signUp)
			auth.POST("/login", h.signIn)
			auth.POST("/avatar", h.uploadAvatar)
		}
		album := api.Group("/album")
		{
			album.POST("/", h.uploadFiles)
		}
	}

	return router
}
