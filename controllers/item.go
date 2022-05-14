package controllers

import (
	"net/http"

	"github.com/eriicafes/fisrtginserver/models"
	"github.com/eriicafes/fisrtginserver/schemas"
	"github.com/gin-gonic/gin"
)

var items = []*models.Item{}

func GetAllItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func CreateItem(ctx *gin.Context) {
	// collect input
	var input schemas.CreateItem

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create item
	newItem := models.NewItem(input)

	// save item
	items = append(items, newItem)

	// return response
	ctx.JSON(http.StatusOK, gin.H{"data": newItem})
}