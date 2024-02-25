package internal

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func ApplicationBootstrap(port int) error {
	// Set gin colorize log
	gin.ForceConsoleColor()

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r.Run(":" + strconv.Itoa(port))
}
