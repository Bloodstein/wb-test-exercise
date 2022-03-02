package handler

import (
	"github.com/Bloodstein/wb-test-exercise/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) Routes() *gin.Engine {

	router := gin.New()

	router.Group("/api")
	{
		ver1 := router.Group("/v1")
		{
			ver1.GET("/items", h.GetAll)
			ver1.GET("/items/:id", h.GetOne)
			ver1.POST("/create", h.Create)
			ver1.POST("/delete", h.Delete)
			ver1.POST("/update", h.Update)
		}
	}

	return router
}
