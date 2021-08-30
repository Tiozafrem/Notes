package handler

import (
	"notes/pkg/usecases"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecases *usecases.Usecases
}

func NewHandler(usecases *usecases.Usecases) *Handler {
	return &Handler{usecases: usecases}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.signUp)
		auth.POST("/sign-in", handler.signIn)
	}

	api := router.Group("/api", handler.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", handler.getAllLists)
			lists.POST("/", handler.createList)
			lists.GET("/:id", handler.getListById)
			lists.PUT("/:id", handler.updateList)
			lists.DELETE("/:id", handler.deleteList)

		}

	}

	return router
}
