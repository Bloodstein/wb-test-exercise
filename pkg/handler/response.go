package handler

import (
	"github.com/Bloodstein/wb-test-exercise/domain"
	"github.com/gin-gonic/gin"
)

func errorResponse(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(code, &domain.ErrorResponse{
		Result:  error_result,
		Message: message,
	})
}
