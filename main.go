package main

import (
	"github.com/gin-gonic/gin"

	"github.com/eriicafes/fisrtginserver/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/", controllers.GetAllItems)
	r.POST("/", controllers.CreateItem)
	r.GET("/:id", controllers.GetItem)

	r.Run()
}
