package service

import (
	"github.com/gin-gonic/gin"
	"judge-engine/ent"
	"judge-engine/internal/common"
	"judge-engine/internal/repository"
	"net/http"
)

func ListSubmission(c *gin.Context) {
	type ListSubmissionQueryString struct {
		Pagination common.PaginationQuery
	}

	var err error
	var queryString = new(ListSubmissionQueryString)
	var submissions []*ent.Submission

	// Parse QueryString
	if err = c.ShouldBind(&queryString); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// Calculate Pagination for query
	queryString.Pagination.CalculatePaginationOption()

	submissions, err = repository.ListSubmission(c, queryString.Pagination)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": submissions,
	})
}
