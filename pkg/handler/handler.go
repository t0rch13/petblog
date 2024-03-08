package handler

import (
	"net/http"
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

	router.LoadHTMLGlob("front/pages/*")
	router.Static("/static", "./front")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})

	router.GET("/pets", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pets.html", nil)
	})

	router.GET("/pet/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pet.html", nil)
	})

	router.GET("/createpet", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createpet.html", nil)
	})

	router.GET("/createpost", func(c *gin.Context) {
		c.HTML(http.StatusOK, "createpost.html", nil)
	})

	router.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts.html", nil)
	})

	router.GET("/post/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", nil)
	})

	router.GET("/api/posts/all", h.getAllPosts)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		posts := api.Group("/posts")
		{
			posts.GET("/:id", h.getPostById)
			posts.POST("/create/:id", h.createPost)
			posts.PUT("/:id", h.updatePost)
			posts.DELETE("/:id", h.deletePost)
			posts.GET("/user", h.getUserPosts)
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
