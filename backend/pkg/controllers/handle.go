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

	// corsConfig := cors.New(cors.Options{
	// AllowedOrigins: []string{"http://localhost:5173/"},
	// AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// AllowedHeaders: []string{"Content-Type", "Authorization"},
	//    })

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", h.signUp)
			auth.POST("/login", h.signIn)
		}
		album := api.Group("/album")
		{
			album.POST("/", h.createAlbum)
		}
	}

	return router
}

// 340 294
// 367 385
