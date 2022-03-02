package handler

import (
	"fmt"
	"net/http"

	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(ctx *gin.Context) {
	allItems, err := h.services.GetAll()

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": allItems,
	})
}

func (h *Handler) GetOne(ctx *gin.Context) {
	searchId, err := getRowId(ctx)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.GetOne(searchId)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": &result,
	})
}

func (h *Handler) Create(ctx *gin.Context) {
	var input domain.TelegramToOfficeRelation

	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Request cannot be parsed")
		return
	}

	newId, err := h.services.Relations.Create(&input)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": success_result,
		"response": map[string]interface{}{
			"id":     newId,
			"get":    fmt.Sprintf("/api/v1/telegram-to-office-relations/items/%d", newId),
			"update": fmt.Sprintf("/api/v1/telegram-to-office-relations/update/%d", newId),
			"delete": fmt.Sprintf("/api/v1/telegram-to-office-relations/delete/%d", newId),
		},
	})
}

func (h *Handler) Delete(ctx *gin.Context) {
	searchId, err := getRowId(ctx)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	result, err := h.services.Delete(searchId)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": result,
	})
}

func (h *Handler) Update(ctx *gin.Context) {
	searchId, err := getRowId(ctx)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var input domain.TelegramToOfficeRelation

	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Request cannot be parsed")
		return
	}

	result, err := h.services.Update(searchId, &input)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": result,
	})
}
