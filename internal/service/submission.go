package service

import (
	"github.com/gin-gonic/gin"
	"judge-engine/ent"
	"judge-engine/internal/repository"
	"net/http"
)

func ListSubmission(c *gin.Context) {
	var err error
	var submissions []*ent.Submission
	submissions, err = repository.ListSubmission(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": submissions,
	})
}
