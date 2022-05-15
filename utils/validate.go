package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate[K interface{}](ctx *gin.Context, input K) (validated K, ok bool) {
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return input, false
	}
	return input, true
}
