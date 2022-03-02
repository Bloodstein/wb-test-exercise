package handler

import (
	"net/http"

	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/gin-gonic/gin"
)

const (
	success_result = "ok"
	error_result   = "error"
)

func successResponse(ctx *gin.Context, response *struct{}) {
	ctx.JSON(http.StatusOK, &domain.SuccessResponse{
		Result:   success_result,
		Response: response,
	})
}

func errorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &domain.ErrorResponse{
		Result:  error_result,
		Message: message,
	})
}
