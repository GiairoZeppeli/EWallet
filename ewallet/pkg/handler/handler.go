package handler

import (
	"github.com/GiairoZeppeli/EWallet/pkg/service"
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
		api.POST("/send", h.Send)
		api.GET("/wallet/:address/balance", h.GetBalance)
		api.GET("/transactions", h.GetLast)
	}
	return router
}
