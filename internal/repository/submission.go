package repository

import (
	"context"
	"judge-engine/ent"
	"judge-engine/internal/database"
)

// List Submission Repository
func ListSubmission(c context.Context) (submissions []*ent.Submission, err error) {
	var client *ent.Client
	client = database.GetClient()
	submissions, err = client.Submission.Query().All(c)
	return
}
