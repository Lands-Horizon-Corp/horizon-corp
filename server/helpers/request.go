package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIDParam(ctx *gin.Context, paramName string) (uint, error) {
	idStr := ctx.Param(paramName)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return 0, err
	}

	if id < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID cannot be negative"})
		return 0, err
	}

	return uint(id), nil
}
