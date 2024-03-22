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

		user := api.Group("/user", h.userIdentity)
		{
			user.POST("/avatar", h.uploadAvatar)
		}

		album := api.Group("/album", h.userIdentity)
		{
			album.POST("/", h.createAlbum)
			album.GET("/", h.getAllAlbum)
		}

		photo := api.Group("/photo", h.userIdentity)
		{
			photo.POST("/", h.addPhoto)
		}

	}

	return router
}


// 340 294
// 367 385