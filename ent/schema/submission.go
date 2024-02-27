package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"judge-engine/ent/schema/enum"
	"time"
)

// Submission holds the schema definition for the Submission entity.
type Submission struct {
	ent.Schema
}

// Fields of the Submission.
func (Submission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Text("code").Comment("Source code of the submission"),
		field.Int("codeLength").Positive().Comment("Lenght of the source code"),
		field.Int("memory").Positive().Comment("Memory usage of the process. Unit is 'KB'"),
		field.Enum("response").GoType(enum.ResponseType("")).Comment("Response linux signal of process"),
		field.Time("createdAt").Default(time.Now).Comment("Submission created time"),
		field.Time("updatedAt").Default(time.Now).Comment("Submission updated time"),
	}
}

// Edges of the Submission.
func (Submission) Edges() []ent.Edge {
	return nil
}
