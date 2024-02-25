package internal

import "github.com/gin-gonic/gin"

func ApplicationBootstrap() {
	// Set gin colorize log
	gin.ForceConsoleColor()

	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
