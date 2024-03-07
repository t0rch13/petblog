package handler

import (
	"petblog/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		posts := api.Group("/posts")
		{
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.POST("/create", h.createPost)
			posts.PUT("/:id", h.updatePost)
			posts.DELETE("/:id", h.deletePost)
		}

		pets := api.Group("/pets")
		{
			pets.GET("/", h.getAllPets)
			pets.GET("/:id", h.getPetById)
			pets.POST("/create", h.createPet)
			pets.PUT("/:id", h.updatePet)
			pets.DELETE("/:id", h.deletePet)
		}
	}

	return router
}
