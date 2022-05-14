package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	data   string
	status uint8
}

func (r response) String() string {
	return fmt.Sprintf("status: %d dat: %s", r.status, r.data)
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		fmt.Println("incoming request")
		ctx.JSON(http.StatusOK, gin.H{"data": response{"bob", 62}.String()})
	})

	r.Run()
}
