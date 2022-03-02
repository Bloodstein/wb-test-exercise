package handler

import (
	"errors"
	"fmt"

	"github.com/Bloodstein/wb-test-exercise/pkg/service"
	"github.com/gin-gonic/gin"
)

const (
	success_result = "ok"
	error_result   = "error"
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

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			rel := v1.Group("/telegram-to-office-relations")
			{
				rel.GET("/items", h.GetAll)
				rel.GET("/items/:id", h.GetOne)
				rel.POST("/create/", h.Create)
				rel.POST("/delete/:id", h.Delete)
				rel.POST("/update/:id", h.Update)
			}
		}
	}

	return router
}

func getRowId(ctx *gin.Context) (string, error) {
	param := ctx.Param("id")

	if len(param) == 0 {
		return "", errors.New(fmt.Sprintf("You've missed the required param \"ID\": %s", param))
	}

	return param, nil
}
