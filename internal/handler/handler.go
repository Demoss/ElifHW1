package handler

import (
	"ElifHW1/internal/service"
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

	api := router.Group("/api")
	{
		sender := api.Group("/")
		{
			sender.POST("/:sender/:receiver", h.Send)
			sender.PUT("/:sender/:receiver/:message", h.UpdateMessage)
			sender.DELETE("/:sender/:receiver", h.DeleteAllMessages)
		}
		receiver := api.Group("/")
		{
			receiver.GET("/:receiver/:sender", h.ReceiveMessages)
		}
	}

	return router
}
