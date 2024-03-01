package repository

import (
	"context"
	"judge-engine/ent"
	"judge-engine/internal/common"
	"judge-engine/internal/database"
)

// List Submission Repository
func ListSubmission(c context.Context, pagination common.PaginationQuery) (submissions []*ent.Submission, err error) {
	var client *ent.Client
	client = database.GetClient()
	submissions, err = client.Submission.Query().Limit(pagination.QueryLimit).Offset(pagination.QueryOffset).All(c)
	return
}
