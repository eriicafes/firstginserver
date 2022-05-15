package controllers

import (
	"net/http"

	"github.com/eriicafes/fisrtginserver/models"
	"github.com/eriicafes/fisrtginserver/schemas"
	"github.com/eriicafes/fisrtginserver/utils"
	"github.com/gin-gonic/gin"
)

var items = []*models.Item{}

func GetAllItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

func CreateItem(ctx *gin.Context) {
	// validate input
	input, ok := utils.Validate(ctx, schemas.CreateItem{})

	if !ok {
		return
	}

	// create item
	newItem := models.NewItem(input)

	// save item
	items = append(items, newItem)

	// return response
	ctx.JSON(http.StatusOK, gin.H{"data": newItem})
}

func GetItem(ctx *gin.Context) {
	if id, ok := ctx.Params.Get("id"); ok {

		for _, item := range items {
			if item.Id == id {
				ctx.JSON(http.StatusOK, gin.H{"data": item})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
	}
}

func UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")

	// validate input
	input, ok := utils.Validate(ctx, schemas.UpdateItem{})

	if !ok {
		return
	}

	for _, item := range items {
		if item.Id == id {
			item.Name = input.Name
			item.Description = input.Description
			item.Labels = input.Labels
			item.IsActive = input.IsActive

			ctx.JSON(http.StatusOK, gin.H{"data": item})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}

func RemoveItem(ctx *gin.Context) {
	// routes intended to match with a param do not need to check if the param exists
	id := ctx.Param("id")

	for index, item := range items {
		if item.Id == id {
			// take all elements up to the item's index (this is not inclusive of the item)
			// merge with all elements after the item's index
			items = append(items[:index], items[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"data": "item removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
}
