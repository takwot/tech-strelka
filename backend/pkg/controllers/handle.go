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

func (h *Handle) InitRoutes(router *gin.Engine) *gin.Engine {




	api := router.Group("/api")
	api.Use(CORS())
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.signUp)
			auth.POST("/login", h.signIn)
		}
		album := api.Group("/album")
		{
			album.POST("/", h.createAlbum)
			album.GET("/", h.getAllAlbum)
			album.GET("/:id", h.getAlbum)
			album.DELETE("/", h.deleteAlbum)
			album.PUT("/", h.updateAlbum)
			album.PUT("/rename", h.renameAlbum)
		}
		photo := api.Group("/photo", h.userIdentity)
		{
			photo.POST("/", h.CreatePhoto)
			photo.GET("/", h.GetPhoto)
			photo.DELETE("/", h.DeletePhoto)
		}
	}
	return router
}

// 340 294
// 367 385
