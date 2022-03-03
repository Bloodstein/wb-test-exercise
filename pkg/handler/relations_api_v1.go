package handler

import (
	"fmt"
	"net/http"

	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAll(ctx *gin.Context) {
	allItems, err := h.services.Relations.GetAll()

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
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

	result, err := h.services.Relations.GetOne(searchId)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": &result,
	})
}

func (h *Handler) Create(ctx *gin.Context) {
	var input domain.ModifyRequest

	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Request cannot be parsed")
		return
	}

	newObjectId, err := h.services.Relations.Create(&input)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result": success_result,
		"response": map[string]string{
			"get":    fmt.Sprintf("/api/v1/telegram-to-office-relations/items/%s", newObjectId),
			"update": fmt.Sprintf("/api/v1/telegram-to-office-relations/update/%s", newObjectId),
			"delete": fmt.Sprintf("/api/v1/telegram-to-office-relations/delete/%s", newObjectId),
			"id":     newObjectId,
		},
	})
}

func (h *Handler) Delete(ctx *gin.Context) {
	searchId, err := getRowId(ctx)

	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	check, _ := h.services.Relations.GetOne(searchId)

	if check == nil {
		errorResponse(ctx, http.StatusInternalServerError, "That document doesn't exist")
		return
	}

	result, err := h.services.Relations.Delete(searchId)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
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

	var input domain.ModifyRequest

	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "Request cannot be parsed")
		return
	}

	check, _ := h.services.Relations.GetOne(searchId)

	if check == nil {
		errorResponse(ctx, http.StatusInternalServerError, "That document doesn't exist")
		return
	}

	result, err := h.services.Relations.Update(searchId, &input)

	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"result":   success_result,
		"response": result,
	})
}
