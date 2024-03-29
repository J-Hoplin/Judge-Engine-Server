// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"judge-engine/ent/schema/enum"
	"judge-engine/ent/submission"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SubmissionCreate is the builder for creating a Submission entity.
type SubmissionCreate struct {
	config
	mutation *SubmissionMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (sc *SubmissionCreate) SetCode(s string) *SubmissionCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetCodeLength sets the "codeLength" field.
func (sc *SubmissionCreate) SetCodeLength(i int) *SubmissionCreate {
	sc.mutation.SetCodeLength(i)
	return sc
}

// SetMemory sets the "memory" field.
func (sc *SubmissionCreate) SetMemory(i int) *SubmissionCreate {
	sc.mutation.SetMemory(i)
	return sc
}

// SetResponse sets the "response" field.
func (sc *SubmissionCreate) SetResponse(et enum.ResponseType) *SubmissionCreate {
	sc.mutation.SetResponse(et)
	return sc
}

// SetCreatedAt sets the "createdAt" field.
func (sc *SubmissionCreate) SetCreatedAt(t time.Time) *SubmissionCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableCreatedAt(t *time.Time) *SubmissionCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updatedAt" field.
func (sc *SubmissionCreate) SetUpdatedAt(t time.Time) *SubmissionCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableUpdatedAt(t *time.Time) *SubmissionCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubmissionCreate) SetID(u uuid.UUID) *SubmissionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubmissionCreate) SetNillableID(u *uuid.UUID) *SubmissionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the SubmissionMutation object of the builder.
func (sc *SubmissionCreate) Mutation() *SubmissionMutation {
	return sc.mutation
}

// Save creates the Submission in the database.
func (sc *SubmissionCreate) Save(ctx context.Context) (*Submission, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubmissionCreate) SaveX(ctx context.Context) *Submission {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubmissionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubmissionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubmissionCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := submission.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := submission.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := submission.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubmissionCreate) check() error {
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Submission.code"`)}
	}
	if _, ok := sc.mutation.CodeLength(); !ok {
		return &ValidationError{Name: "codeLength", err: errors.New(`ent: missing required field "Submission.codeLength"`)}
	}
	if v, ok := sc.mutation.CodeLength(); ok {
		if err := submission.CodeLengthValidator(v); err != nil {
			return &ValidationError{Name: "codeLength", err: fmt.Errorf(`ent: validator failed for field "Submission.codeLength": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Memory(); !ok {
		return &ValidationError{Name: "memory", err: errors.New(`ent: missing required field "Submission.memory"`)}
	}
	if v, ok := sc.mutation.Memory(); ok {
		if err := submission.MemoryValidator(v); err != nil {
			return &ValidationError{Name: "memory", err: fmt.Errorf(`ent: validator failed for field "Submission.memory": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Response(); !ok {
		return &ValidationError{Name: "response", err: errors.New(`ent: missing required field "Submission.response"`)}
	}
	if v, ok := sc.mutation.Response(); ok {
		if err := submission.ResponseValidator(v); err != nil {
			return &ValidationError{Name: "response", err: fmt.Errorf(`ent: validator failed for field "Submission.response": %w`, err)}
		}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Submission.createdAt"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Submission.updatedAt"`)}
	}
	return nil
}

func (sc *SubmissionCreate) sqlSave(ctx context.Context) (*Submission, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubmissionCreate) createSpec() (*Submission, *sqlgraph.CreateSpec) {
	var (
		_node = &Submission{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(submission.Table, sqlgraph.NewFieldSpec(submission.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.SetField(submission.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := sc.mutation.CodeLength(); ok {
		_spec.SetField(submission.FieldCodeLength, field.TypeInt, value)
		_node.CodeLength = value
	}
	if value, ok := sc.mutation.Memory(); ok {
		_spec.SetField(submission.FieldMemory, field.TypeInt, value)
		_node.Memory = value
	}
	if value, ok := sc.mutation.Response(); ok {
		_spec.SetField(submission.FieldResponse, field.TypeEnum, value)
		_node.Response = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(submission.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(submission.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SubmissionCreateBulk is the builder for creating many Submission entities in bulk.
type SubmissionCreateBulk struct {
	config
	err      error
	builders []*SubmissionCreate
}

// Save creates the Submission entities in the database.
func (scb *SubmissionCreateBulk) Save(ctx context.Context) ([]*Submission, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Submission, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubmissionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubmissionCreateBulk) SaveX(ctx context.Context) []*Submission {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubmissionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubmissionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
