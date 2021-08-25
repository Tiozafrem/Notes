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

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	return router
}
