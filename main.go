package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eriicafes/fisrtginserver/models"
	"github.com/eriicafes/fisrtginserver/schemas"
)

func main() {
	items := []*models.Item{}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": items})
	})

	r.POST("/", func(ctx *gin.Context) {
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
	})

	r.Run()
}
