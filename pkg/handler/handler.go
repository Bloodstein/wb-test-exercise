package handler

import (
	"errors"
	"fmt"
	"strconv"

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

	router.Group("/api")
	{
		router.Group("/v1")
		{
			rel := router.Group("/telegram-to-office-relations")
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

func getRowId(ctx *gin.Context) (int, error) {
	param := ctx.Param("id")

	if len(param) == 0 {
		return 0, errors.New(fmt.Sprintf("You've missed the required param \"ID\": %s", param))
	}

	rowId, err := strconv.Atoi(param)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("Fail to parse \"ID\": %s", param))
	}

	return rowId, nil
}
