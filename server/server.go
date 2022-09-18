package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "I'm good don't worry",
	})
}

func StartServer() {
	r := gin.Default()
	r.GET("/", healthCheck)
	r.Run(":3210")
}
