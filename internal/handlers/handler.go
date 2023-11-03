package handlers

import (
	"github.com/gin-gonic/gin"
	"openai/internal/services"
)

type Handler struct {
	service services.Service
}

func NewHandler(service services.Service) *Handler {
	return &Handler{service: service}
}

// Init создает роутер с обработчиками
func (h *Handler) Init() *gin.Engine {
	r := gin.New()

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("query", h.sendSingleQuery)
		}
	}

	return r
}
