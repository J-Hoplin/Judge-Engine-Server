package controller

import (
	"github.com/gin-gonic/gin"
	"judge-engine/internal/service"
)

func SubmissionRoute(engine *gin.Engine) {
	router := engine.Group("/submissions")
	{
		router.GET("/", service.ListSubmission)
		router.POST("/")
	}

}
