package handler

import (
	"ElifHW1/internal/repository"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repos *repository.Repository
}

func NewHandler(repos *repository.Repository) *Handler {
	return &Handler{repos: repos}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		sender := api.Group("/")
		{
			sender.POST("/:sender/:receiver", h.Send)
		}
		receiver := api.Group("/")
		{
			receiver.GET("/:receiver/:sender", h.ReceiveMessages)
		}
	}

	return router
}
