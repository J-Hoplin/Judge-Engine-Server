// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"judge-engine/ent/predicate"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeSubmission = "Submission"
)

// SubmissionMutation represents an operation that mutates the Submission nodes in the graph.
type SubmissionMutation struct {
	config
	op            Op
	typ           string
	id            *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Submission, error)
	predicates    []predicate.Submission
}

var _ ent.Mutation = (*SubmissionMutation)(nil)

// submissionOption allows management of the mutation configuration using functional options.
type submissionOption func(*SubmissionMutation)

// newSubmissionMutation creates new mutation for the Submission entity.
func newSubmissionMutation(c config, op Op, opts ...submissionOption) *SubmissionMutation {
	m := &SubmissionMutation{
		config:        c,
		op:            op,
		typ:           TypeSubmission,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withSubmissionID sets the ID field of the mutation.
func withSubmissionID(id int) submissionOption {
	return func(m *SubmissionMutation) {
		var (
			err   error
			once  sync.Once
			value *Submission
		)
		m.oldValue = func(ctx context.Context) (*Submission, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Submission.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withSubmission sets the old Submission of the mutation.
func withSubmission(node *Submission) submissionOption {
	return func(m *SubmissionMutation) {
		m.oldValue = func(context.Context) (*Submission, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m SubmissionMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m SubmissionMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *SubmissionMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *SubmissionMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Submission.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// Where appends a list predicates to the SubmissionMutation builder.
func (m *SubmissionMutation) Where(ps ...predicate.Submission) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the SubmissionMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *SubmissionMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Submission, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *SubmissionMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *SubmissionMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Submission).
func (m *SubmissionMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *SubmissionMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *SubmissionMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *SubmissionMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Submission field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SubmissionMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Submission field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *SubmissionMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *SubmissionMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *SubmissionMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Submission numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *SubmissionMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *SubmissionMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *SubmissionMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Submission nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *SubmissionMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Submission field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *SubmissionMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *SubmissionMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *SubmissionMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *SubmissionMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *SubmissionMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *SubmissionMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *SubmissionMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Submission unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *SubmissionMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Submission edge %s", name)
}
